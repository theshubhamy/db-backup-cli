# Database Backup CLI Utility (Golang)

This is a command-line interface (CLI) utility built with Golang that automates database backups. It supports multiple DBMS types, including MySQL, PostgreSQL, and MongoDB, with features such as automatic scheduling, file compression, and cloud/local storage options.

## Features

- **Multiple DBMS Support**: MySQL, PostgreSQL, MongoDB
- **Backup Types**: Full, incremental, differential (customize based on DBMS)
- **Scheduling**: Cron-based automatic backup scheduling
- **Compression**: Compress backup files for storage efficiency
- **Storage Options**: Local storage, AWS S3, Google Cloud Storage (coming soon)
- **Logging**: Track backup and restore operations

## Prerequisites

Before you start, ensure you have the following installed:

- **Golang**: Download from [https://golang.org/dl/](https://golang.org/dl/)
- **MySQL**: Download from [https://dev.mysql.com/downloads/](https://dev.mysql.com/downloads/)
- **PostgreSQL**: Download from [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
- **MongoDB**: Download from [https://www.mongodb.com/try/download/community](https://www.mongodb.com/try/download/community)

Ensure that the following command-line tools are installed for database backups:

- **MySQL**: `mysqldump`
- **PostgreSQL**: `pg_dump`
- **MongoDB**: `mongodump`

## Installation

1. **Clone the Repository:**

   Clone the repository from GitHub (or download the project):

   ```bash
   git clone https://github.com/your-username/db-backup-cli.git
   ```

2. **Navigate to the Project Directory:**

   ```bash
   cd db-backup-cli
   ```

3. **Install Dependencies:**

   Install the necessary Go dependencies:

   ```bash
   go mod tidy
   ```

4. **Build the Project:**

   Compile the project to create an executable:

   ```bash
   go build -o db-backup-cli
   ```

5. **Run the CLI Tool:**

   After building, you can run the CLI with:

   ```bash
   ./db-backup-cli
   ```

## Configuration

### Setting Environment Variables

For MySQL and PostgreSQL backups, you'll need to set the following environment variables for authentication:

- MySQL:

  ```bash
  export MYSQL_USER="your_mysql_user"
  export MYSQL_PASSWORD="your_mysql_password"
  ```

- PostgreSQL:
  ```bash
  export PG_USER="your_postgres_user"
  export PG_HOST="your_postgres_host"
  ```

These credentials will be used by the CLI to perform backups.

## Usage

### 1. **Perform a Backup**

To perform a manual backup, use the `backup` command and specify the database type (`mysql`, `postgres`, `mongodb`), along with other required flags.

#### Example for MySQL:

```bash
./db-backup-cli backup --db-type=mysql --host=localhost --port=3306 --user=root --password=secret --db-name=mydb --backup-type=full --storage=local --compress
```

#### Example for PostgreSQL:

```bash
./db-backup-cli backup --db-type=postgres --host=localhost --port=5432 --user=postgres --password=secret --db-name=mydb --backup-type=full --storage=local --compress
```

#### Example for MongoDB:

```bash
./db-backup-cli backup --db-type=mongodb --host=localhost --port=27017 --user=mongoUser --password=mongoPass --db-name=mydb --backup-type=full --storage=local --compress
```

### 2. **Schedule a Backup**

You can schedule automatic backups using the `schedule` command and providing a cron expression.

#### Example: Schedule a Daily Backup at Midnight for MySQL

```bash
./db-backup-cli schedule --db-type=mysql --host=localhost --port=3306 --user=root --password=secret --db-name=mydb --backup-type=full --storage=local --compress --cron="0 0 * * *"
```

### 3. **Restore a Backup**

To restore a database from a backup file, use the `restore` command.

#### Example for MySQL:

```bash
./db-backup-cli restore --db-type=mysql --host=localhost --port=3306 --user=root --password=secret --db-name=mydb --backup-file=mydb_backup.sql
```

### Available Flags:

- `--db-type`: Type of database (`mysql`, `postgres`, `mongodb`).
- `--host`: Database host (e.g., `localhost`).
- `--port`: Database port (e.g., `3306` for MySQL).
- `--user`: Database username.
- `--password`: Database password.
- `--db-name`: Name of the database.
- `--backup-type`: Type of backup (`full`, `incremental`, `differential`).
- `--storage`: Where to store the backup (`local`, `aws`, `gcp`).
- `--compress`: Whether to compress the backup file.
- `--cron`: Cron expression for scheduling.
- `--backup-file`: Path to the backup file for restoration.

## Logging

Backup and restore operations are logged in `logs/backup.log`. You can review this file for the status of operations, including success and error messages.

## Contributing

Feel free to fork this repository and submit pull requests for new features or bug fixes. Contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Additional Notes:

- Ensure that `mysqldump`, `pg_dump`, and `mongodump` are installed and accessible in your system's `PATH`.
- For cloud storage options (like AWS S3 or Google Cloud Storage), integration is currently in progress. You can contribute or extend the functionality for cloud backups.

---

Let me know if you need further customizations or clarifications for the README file!
