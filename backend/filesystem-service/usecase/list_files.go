package usecase

import (
	"github.com/augment-local-manus-clone/backend/filesystem-service/domain"
)

// ListFilesUseCase handles listing files
type ListFilesUseCase struct {
	filesystemClient domain.FilesystemClient
}

// NewListFilesUseCase creates a new instance of ListFilesUseCase
func NewListFilesUseCase(filesystemClient domain.FilesystemClient) *ListFilesUseCase {
	return &ListFilesUseCase{
		filesystemClient: filesystemClient,
	}
}

// Execute lists files in a directory
func (uc *ListFilesUseCase) Execute(path string) ([]*domain.FileInfo, error) {
	// Create a new file operation
	operation := domain.NewFileOperation(path, "list", "")

	// List files
	files, err := uc.filesystemClient.ListFiles(path)
	operation.SetResult(err == nil, err)

	return files, err
}
