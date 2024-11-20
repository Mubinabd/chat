package repository

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Mubinabd/chat/internal/entity"
)

type FileRepo struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) *FileRepo {
	return &FileRepo{
		db: db,
	}
}
func (r *FileRepo) SaveFile(req *entity.SaveFile) (*entity.File, error) {
	destFile, err := os.Create(req.FilePath)
	if err != nil {
		log.Printf("Failed to create file at path %s: %v", req.FilePath, err)
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer destFile.Close()

	sourceFile, err := os.Open(req.FileName)
	if err != nil {
		log.Printf("Failed to open source file %s: %v", req.FileName, err)
		return nil, fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	bytesWritten, err := io.Copy(destFile, sourceFile)
	if err != nil {
		log.Printf("Failed to copy file contents to %s: %v", req.FilePath, err)
		return nil, fmt.Errorf("failed to copy file contents: %w", err)
	}

	log.Printf("File %s successfully saved at %s. Total bytes written: %d", req.FileName, req.FilePath, bytesWritten)

	return &entity.File{
		FileName: req.FileName,
		URL:      req.FilePath,
	}, nil
}
