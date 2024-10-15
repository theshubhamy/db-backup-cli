package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theshubhamy/db-backup-cli/db"
	"github.com/theshubhamy/db-backup-cli/storage"
	"github.com/theshubhamy/db-backup-cli/utils"
)

var (
	dbType      string
	host        string
	port        string
	user        string
	password    string
	dbName      string
	backupType  string
	storageType string
	compress    bool
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Performs a database backup",
	Run: func(cmd *cobra.Command, args []string) {
		// Database connection
		connector, err := db.NewConnector(dbType, host, port, user, password, dbName)
		if err != nil {
			fmt.Println("Error connecting to database:", err)
			return
		}

		// Perform backup
		backupFile := fmt.Sprintf("%s_backup.sql", dbName)
		err = connector.PerformBackup(backupType)
		if err != nil {
			fmt.Println("Error during backup:", err)
			return
		}

		// Compress backup if requested
		if compress {
			compressedFile := fmt.Sprintf("%s_backup.zip", dbName)
			err = utils.CompressFile(backupFile, compressedFile)
			if err != nil {
				fmt.Println("Error compressing backup:", err)
				return
			}
			backupFile = compressedFile // Update to store compressed file
		}

		// Handle storage
		store := storage.NewStorage(storageType)
		err = store.StoreBackup(backupFile, backupType)
		if err != nil {
			fmt.Println("Error storing backup:", err)
			return
		}
		fmt.Println("Backup successful!")
	},
}

func init() {
	backupCmd.Flags().StringVar(&dbType, "db-type", "", "Database type (mysql, postgres, mongodb)")
	backupCmd.Flags().StringVar(&host, "host", "", "Database host")
	backupCmd.Flags().StringVar(&port, "port", "", "Database port")
	backupCmd.Flags().StringVar(&user, "user", "", "Database user")
	backupCmd.Flags().StringVar(&password, "password", "", "Database password")
	backupCmd.Flags().StringVar(&dbName, "db-name", "", "Database name")
	backupCmd.Flags().StringVar(&backupType, "backup-type", "full", "Backup type (full, incremental, differential)")
	backupCmd.Flags().StringVar(&storageType, "storage", "local", "Storage type (local, aws, gcp)")
	backupCmd.Flags().BoolVar(&compress, "compress", false, "Compress backup file")
	rootCmd.AddCommand(backupCmd)
}
