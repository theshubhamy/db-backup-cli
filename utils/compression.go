package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CompressFile compresses the specified file into a .zip file.
func CompressFile(source, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	fileToZip, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("failed to open file for compression: %v", err)
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return fmt.Errorf("could not obtain file information: %v", err)
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("could not create zip header: %v", err)
	}

	header.Method = zip.Deflate
	writer, err := archive.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("could not create writer in archive: %v", err)
	}

	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return fmt.Errorf("could not write file to zip archive: %v", err)
	}

	fmt.Println("File successfully compressed:", target)
	return nil
}

// CompressDirectory compresses all the files in a directory into a .zip file.
func CompressDirectory(sourceDir, targetZip string) error {
	zipFile, err := os.Create(targetZip)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		fileToZip, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %v", err)
		}
		defer fileToZip.Close()

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("failed to create zip header: %v", err)
		}

		header.Name = path
		header.Method = zip.Deflate
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("failed to create zip writer: %v", err)
		}

		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return fmt.Errorf("failed to write file to archive: %v", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking through directory: %v", err)
	}

	fmt.Println("Directory successfully compressed:", targetZip)
	return nil
}
