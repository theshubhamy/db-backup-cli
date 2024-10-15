package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connector struct {
	dbType string
	dbConn interface{}
	dbName string
}

func NewConnector(dbType, host, port, user, password, dbName string) (*Connector, error) {
	switch dbType {
	case "mysql":
		return connectMySQL(host, port, user, password, dbName)
	case "postgres":
		return connectPostgres(host, port, user, password, dbName)
	case "mongodb":
		return connectMongoDB(host, port, user, password, dbName)
	default:
		return nil, errors.New("unsupported database type")
	}
}

// connectMySQL connects to a MySQL database.
func connectMySQL(host, port, user, password, dbName string) (*Connector, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}
	return &Connector{dbType: "mysql", dbConn: db, dbName: dbName}, nil
}

// connectPostgres connects to a PostgreSQL database.
func connectPostgres(host, port, user, password, dbName string) (*Connector, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}
	return &Connector{dbType: "postgres", dbConn: db, dbName: dbName}, nil
}

// connectMongoDB connects to a MongoDB database.
func connectMongoDB(host, port, user, password, dbName string) (*Connector, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("MongoDB connection failed: %v", err)
	}

	return &Connector{dbType: "mongodb", dbConn: client, dbName: dbName}, nil
}

// PerformBackup performs a backup based on the type of the database and backup type.
func (c *Connector) PerformBackup(backupType string) error {
	switch c.dbType {
	case "mysql":
		return c.performMySQLBackup(backupType)
	case "postgres":
		return c.performPostgresBackup(backupType)
	case "mongodb":
		return c.performMongoBackup(backupType)
	default:
		return errors.New("unsupported database type for backup")
	}
}

// performMySQLBackup runs the mysqldump command to create a backup for MySQL.
func (c *Connector) performMySQLBackup(backupType string) error {
	fileName := fmt.Sprintf("%s_backup.sql", c.dbName)
	cmd := exec.Command("mysqldump", "-u", os.Getenv("MYSQL_USER"), "-p"+os.Getenv("MYSQL_PASSWORD"), c.dbName, ">", fileName)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to backup MySQL database: %v", err)
	}
	fmt.Printf("MySQL backup successful: %s\n", fileName)
	return nil
}

// performPostgresBackup runs the pg_dump command to create a backup for PostgreSQL.
func (c *Connector) performPostgresBackup(backupType string) error {
	fileName := fmt.Sprintf("%s_backup.sql", c.dbName)
	cmd := exec.Command("pg_dump", "-U", os.Getenv("PG_USER"), "-h", os.Getenv("PG_HOST"), "-d", c.dbName, "-f", fileName)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to backup PostgreSQL database: %v", err)
	}
	fmt.Printf("PostgreSQL backup successful: %s\n", fileName)
	return nil
}

// performMongoBackup runs the mongodump command to create a backup for MongoDB.
func (c *Connector) performMongoBackup(backupType string) error {
	dirName := fmt.Sprintf("%s_backup", c.dbName)
	cmd := exec.Command("mongodump", "--db", c.dbName, "--out", dirName)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to backup MongoDB database: %v", err)
	}
	fmt.Printf("MongoDB backup successful: %s\n", dirName)
	return nil
}
