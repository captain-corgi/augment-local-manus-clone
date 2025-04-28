package usecase

import (
	"github.com/augment-local-manus-clone/backend/filesystem-service/domain"
)

// ReadFileUseCase handles reading files
type ReadFileUseCase struct {
	filesystemClient domain.FilesystemClient
}

// NewReadFileUseCase creates a new instance of ReadFileUseCase
func NewReadFileUseCase(filesystemClient domain.FilesystemClient) *ReadFileUseCase {
	return &ReadFileUseCase{
		filesystemClient: filesystemClient,
	}
}

// Execute reads a file
func (uc *ReadFileUseCase) Execute(path string) (*domain.FileContent, error) {
	// Create a new file operation
	operation := domain.NewFileOperation(path, "read", "")

	// Read the file
	content, err := uc.filesystemClient.ReadFile(path)
	operation.SetResult(err == nil, err)

	return content, err
}
