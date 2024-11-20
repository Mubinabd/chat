package usecase

import (
    "github.com/Mubinabd/chat/internal/storage/repository"
    "github.com/Mubinabd/chat/internal/entity"
)

type FileUseCase struct {
    storage repository.Storage
}

func NewFileUseCase(fileRepo *repository.Storage) *FileUseCase {
    return &FileUseCase{storage: *fileRepo}
}

func (uc *FileUseCase) SaveFile(req *entity.SaveFile) (*entity.File, error) {
    return uc.storage.File().SaveFile(req)
}
