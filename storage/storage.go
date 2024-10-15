package storage

import (
	"fmt"
)

type Storage struct {
	storageType string
}

func NewStorage(storageType string) *Storage {
	return &Storage{storageType: storageType}
}

func (s *Storage) StoreBackup(backupType, dbName string) error {
	switch s.storageType {
	case "local":
		fmt.Println("Storing backup locally")
		// Store locally
	case "aws":
		fmt.Println("Uploading backup to AWS S3")
		// AWS S3 upload logic
	case "gcp":
		fmt.Println("Uploading backup to Google Cloud")
		// GCP upload logic
	case "azure":
		fmt.Println("Uploading backup to Azure Blob Storage")
		// Azure Blob upload logic
	default:
		return fmt.Errorf("unsupported storage type: %s", s.storageType)
	}
	return nil
}
