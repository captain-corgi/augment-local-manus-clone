package usecase

import (
	"github.com/augment-local-manus-clone/backend/filesystem-service/domain"
)

// DeleteFileUseCase handles deleting files
type DeleteFileUseCase struct {
	filesystemClient domain.FilesystemClient
}

// NewDeleteFileUseCase creates a new instance of DeleteFileUseCase
func NewDeleteFileUseCase(filesystemClient domain.FilesystemClient) *DeleteFileUseCase {
	return &DeleteFileUseCase{
		filesystemClient: filesystemClient,
	}
}

// Execute deletes a file
func (uc *DeleteFileUseCase) Execute(path string) (*domain.FileOperation, error) {
	// Create a new file operation
	operation := domain.NewFileOperation(path, "delete", "")

	// Delete the file
	err := uc.filesystemClient.DeleteFile(path)
	operation.SetResult(err == nil, err)

	return operation, err
}
