package domain

import (
	"errors"
	"os"
	"time"
)

// FileType represents the type of a file
type FileType string

const (
	FileTypeFile      FileType = "file"
	FileTypeDirectory FileType = "directory"
)

// FileInfo represents information about a file
type FileInfo struct {
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Type         FileType  `json:"type"`
	Size         int64     `json:"size"`
	ModifiedTime time.Time `json:"modified_time"`
	IsHidden     bool      `json:"is_hidden"`
}

// FileContent represents the content of a file
type FileContent struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// FileOperation represents a file operation
type FileOperation struct {
	Path      string    `json:"path"`
	Operation string    `json:"operation"`
	Content   string    `json:"content,omitempty"`
	Success   bool      `json:"success"`
	Error     string    `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// NewFileOperation creates a new file operation
func NewFileOperation(path, operation string, content string) *FileOperation {
	return &FileOperation{
		Path:      path,
		Operation: operation,
		Content:   content,
		Timestamp: time.Now(),
	}
}

// SetResult sets the result of the file operation
func (fo *FileOperation) SetResult(success bool, err error) {
	fo.Success = success
	if err != nil {
		fo.Error = err.Error()
	}
}

// FilesystemClient defines the interface for filesystem operations
type FilesystemClient interface {
	// ReadFile reads a file
	ReadFile(path string) (*FileContent, error)

	// WriteFile writes to a file
	WriteFile(path, content string) error

	// ListFiles lists files in a directory
	ListFiles(path string) ([]*FileInfo, error)

	// DeleteFile deletes a file
	DeleteFile(path string) error

	// MakeDirectory creates a directory
	MakeDirectory(path string) error

	// IsPathSafe checks if a path is safe (within the workspace)
	IsPathSafe(path string) (string, error)

	// GetFileInfo gets information about a file
	GetFileInfo(path string) (*FileInfo, error)
}

// ErrPathTraversal is returned when a path traversal attempt is detected
var ErrPathTraversal = errors.New("path traversal attempt detected")

// ErrPathNotFound is returned when a path is not found
var ErrPathNotFound = errors.New("path not found")

// ErrNotAFile is returned when a path is not a file
var ErrNotAFile = errors.New("not a file")

// ErrNotADirectory is returned when a path is not a directory
var ErrNotADirectory = errors.New("not a directory")

// ErrFileAlreadyExists is returned when a file already exists
var ErrFileAlreadyExists = errors.New("file already exists")

// ErrPermissionDenied is returned when permission is denied
var ErrPermissionDenied = errors.New("permission denied")

// FileInfoFromOsFileInfo converts os.FileInfo to FileInfo
func FileInfoFromOsFileInfo(info os.FileInfo, path string) *FileInfo {
	fileType := FileTypeFile
	if info.IsDir() {
		fileType = FileTypeDirectory
	}

	return &FileInfo{
		Name:         info.Name(),
		Path:         path,
		Type:         fileType,
		Size:         info.Size(),
		ModifiedTime: info.ModTime(),
		IsHidden:     info.Name()[0] == '.',
	}
}
