package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/augment-local-manus-clone/backend/filesystem-service/domain"
)

// FilesystemClientImpl implements the FilesystemClient interface
type FilesystemClientImpl struct {
	workspaceDir string
}

// NewFilesystemClient creates a new FilesystemClientImpl
func NewFilesystemClient(workspaceDir string) *FilesystemClientImpl {
	return &FilesystemClientImpl{
		workspaceDir: workspaceDir,
	}
}

// ReadFile reads a file
func (fc *FilesystemClientImpl) ReadFile(path string) (*domain.FileContent, error) {
	// Check if path is safe
	safePath, err := fc.IsPathSafe(path)
	if err != nil {
		return nil, err
	}

	// Check if path exists
	fileInfo, err := os.Stat(safePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.ErrPathNotFound
		}
		return nil, err
	}

	// Check if path is a file
	if fileInfo.IsDir() {
		return nil, domain.ErrNotAFile
	}

	// Read file
	content, err := ioutil.ReadFile(safePath)
	if err != nil {
		return nil, err
	}

	return &domain.FileContent{
		Path:    path,
		Content: string(content),
	}, nil
}

// WriteFile writes to a file
func (fc *FilesystemClientImpl) WriteFile(path, content string) error {
	// Check if path is safe
	safePath, err := fc.IsPathSafe(path)
	if err != nil {
		return err
	}

	// Create parent directories if they don't exist
	dir := filepath.Dir(safePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write file
	return ioutil.WriteFile(safePath, []byte(content), 0644)
}

// ListFiles lists files in a directory
func (fc *FilesystemClientImpl) ListFiles(path string) ([]*domain.FileInfo, error) {
	// Check if path is safe
	safePath, err := fc.IsPathSafe(path)
	if err != nil {
		return nil, err
	}

	// Check if path exists
	fileInfo, err := os.Stat(safePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.ErrPathNotFound
		}
		return nil, err
	}

	// Check if path is a directory
	if !fileInfo.IsDir() {
		return nil, domain.ErrNotADirectory
	}

	// Read directory
	entries, err := ioutil.ReadDir(safePath)
	if err != nil {
		return nil, err
	}

	// Convert to FileInfo
	var files []*domain.FileInfo
	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		files = append(files, domain.FileInfoFromOsFileInfo(entry, entryPath))
	}

	return files, nil
}

// DeleteFile deletes a file
func (fc *FilesystemClientImpl) DeleteFile(path string) error {
	// Check if path is safe
	safePath, err := fc.IsPathSafe(path)
	if err != nil {
		return err
	}

	// Check if path exists
	_, err = os.Stat(safePath)
	if err != nil {
		if os.IsNotExist(err) {
			return domain.ErrPathNotFound
		}
		return err
	}

	// Delete file
	return os.RemoveAll(safePath)
}

// MakeDirectory creates a directory
func (fc *FilesystemClientImpl) MakeDirectory(path string) error {
	// Check if path is safe
	safePath, err := fc.IsPathSafe(path)
	if err != nil {
		return err
	}

	// Create directory
	return os.MkdirAll(safePath, 0755)
}

// IsPathSafe checks if a path is safe (within the workspace)
func (fc *FilesystemClientImpl) IsPathSafe(path string) (string, error) {
	// Clean the path
	cleanPath := filepath.Clean(path)

	// Check for path traversal attempts
	if strings.Contains(cleanPath, "..") {
		return "", domain.ErrPathTraversal
	}

	// Remove leading slash if present
	if strings.HasPrefix(cleanPath, "/") {
		cleanPath = cleanPath[1:]
	}

	// Join with workspace directory
	safePath := filepath.Join(fc.workspaceDir, cleanPath)

	// Check if path is within workspace
	absWorkspace, err := filepath.Abs(fc.workspaceDir)
	if err != nil {
		return "", err
	}

	absSafePath, err := filepath.Abs(safePath)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(absSafePath, absWorkspace) {
		return "", domain.ErrPathTraversal
	}

	return safePath, nil
}

// GetFileInfo gets information about a file
func (fc *FilesystemClientImpl) GetFileInfo(path string) (*domain.FileInfo, error) {
	// Check if path is safe
	safePath, err := fc.IsPathSafe(path)
	if err != nil {
		return nil, err
	}

	// Get file info
	fileInfo, err := os.Stat(safePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.ErrPathNotFound
		}
		return nil, err
	}

	return domain.FileInfoFromOsFileInfo(fileInfo, path), nil
}
