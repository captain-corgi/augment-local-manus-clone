package usecase

import (
	"github.com/augment-local-manus-clone/backend/filesystem-service/domain"
)

// WriteFileUseCase handles writing to files
type WriteFileUseCase struct {
	filesystemClient domain.FilesystemClient
}

// NewWriteFileUseCase creates a new instance of WriteFileUseCase
func NewWriteFileUseCase(filesystemClient domain.FilesystemClient) *WriteFileUseCase {
	return &WriteFileUseCase{
		filesystemClient: filesystemClient,
	}
}

// Execute writes to a file
func (uc *WriteFileUseCase) Execute(path, content string) (*domain.FileOperation, error) {
	// Create a new file operation
	operation := domain.NewFileOperation(path, "write", content)

	// Write to the file
	err := uc.filesystemClient.WriteFile(path, content)
	operation.SetResult(err == nil, err)

	return operation, err
}
