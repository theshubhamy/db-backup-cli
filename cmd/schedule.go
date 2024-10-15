package cmd

import (
	"fmt"
	"time"

	"github.com/theshubhamy/db-backup-cli/db"
	"github.com/theshubhamy/db-backup-cli/storage"
	"github.com/theshubhamy/db-backup-cli/utils"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var (
	scheduleType       string // Daily, weekly, monthly, or custom
	scheduleCron       string // Cron expression for custom schedules
	compressOnSchedule bool   // Flag to enable compression during scheduled backups
)

// scheduleCmd defines the command to schedule automatic backups
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedules automatic backups using predefined schedules or custom cron expressions",
	Run: func(cmd *cobra.Command, args []string) {
		// Determine the cron expression based on the scheduleType
		switch scheduleType {
		case "daily":
			scheduleCron = "0 0 * * *" // Every day at midnight
		case "weekly":
			scheduleCron = "0 0 * * 0" // Every Sunday at midnight
		case "monthly":
			scheduleCron = "0 0 1 * *" // First day of every month at midnight
		case "custom":
			if scheduleCron == "" {
				fmt.Println("For custom scheduling, please provide a valid cron expression using --cron")
				return
			}
		default:
			fmt.Println("Invalid schedule type. Please use one of: daily, weekly, monthly, custom.")
			return
		}

		// Validate cron expression
		c := cron.New()
		_, err := c.AddFunc(scheduleCron, func() {
			// Perform the scheduled backup
			err := runScheduledBackup()
			if err != nil {
				fmt.Println("Error during scheduled backup:", err)
			} else {
				fmt.Println("Scheduled backup completed successfully at", time.Now())
			}
		})
		if err != nil {
			fmt.Println("Invalid cron expression:", err)
			return
		}

		// Start the cron scheduler
		c.Start()
		fmt.Printf("Scheduled backup with cron expression [%s]. Waiting for next run...\n", scheduleCron)

		// Keep the program running to allow scheduled jobs
		select {}
	},
}

// runScheduledBackup performs the actual backup operation for the scheduled task
func runScheduledBackup() error {
	// Database connection setup
	connector, err := db.NewConnector(dbType, host, port, user, password, dbName)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	// Perform the database backup
	backupFile := fmt.Sprintf("%s_backup.sql", dbName)
	err = connector.PerformBackup(backupType)
	if err != nil {
		return fmt.Errorf("error performing backup: %v", err)
	}

	// Compress the backup if the compress flag is set
	if compressOnSchedule {
		compressedFile := fmt.Sprintf("%s_backup.zip", dbName)
		err = utils.CompressFile(backupFile, compressedFile)
		if err != nil {
			return fmt.Errorf("error compressing backup: %v", err)
		}
		backupFile = compressedFile // Update to the compressed file for storage
	}

	// Store the backup (local or cloud)
	store := storage.NewStorage(storageType)
	err = store.StoreBackup(backupFile, backupType) // Store both the file and the backup type
	if err != nil {
		return fmt.Errorf("error storing backup: %v", err)
	}

	return nil
}

func init() {
	// Define flags for scheduling the backup
	scheduleCmd.Flags().StringVar(&scheduleType, "schedule-type", "", "Type of schedule (daily, weekly, monthly, custom)")
	scheduleCmd.Flags().StringVar(&scheduleCron, "cron", "", "Cron expression for custom scheduling (required if schedule-type is 'custom')")
	scheduleCmd.Flags().StringVar(&dbType, "db-type", "", "Database type (mysql, postgres, mongodb)")
	scheduleCmd.Flags().StringVar(&host, "host", "", "Database host")
	scheduleCmd.Flags().StringVar(&port, "port", "", "Database port")
	scheduleCmd.Flags().StringVar(&user, "user", "", "Database user")
	scheduleCmd.Flags().StringVar(&password, "password", "", "Database password")
	scheduleCmd.Flags().StringVar(&dbName, "db-name", "", "Database name")
	scheduleCmd.Flags().StringVar(&backupType, "backup-type", "full", "Backup type (full, incremental, differential)")
	scheduleCmd.Flags().StringVar(&storageType, "storage", "local", "Storage type (local, aws, gcp)")
	scheduleCmd.Flags().BoolVar(&compressOnSchedule, "compress", false, "Compress backup file during scheduled backups")
	rootCmd.AddCommand(scheduleCmd)
}
