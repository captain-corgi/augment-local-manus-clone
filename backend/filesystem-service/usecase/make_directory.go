package usecase

import (
	"github.com/augment-local-manus-clone/backend/filesystem-service/domain"
)

// MakeDirectoryUseCase handles creating directories
type MakeDirectoryUseCase struct {
	filesystemClient domain.FilesystemClient
}

// NewMakeDirectoryUseCase creates a new instance of MakeDirectoryUseCase
func NewMakeDirectoryUseCase(filesystemClient domain.FilesystemClient) *MakeDirectoryUseCase {
	return &MakeDirectoryUseCase{
		filesystemClient: filesystemClient,
	}
}

// Execute creates a directory
func (uc *MakeDirectoryUseCase) Execute(path string) (*domain.FileOperation, error) {
	// Create a new file operation
	operation := domain.NewFileOperation(path, "mkdir", "")

	// Create the directory
	err := uc.filesystemClient.MakeDirectory(path)
	operation.SetResult(err == nil, err)

	return operation, err
}
