package docker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/augment-local-manus-clone/backend/code-execution-service/domain"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// DockerClientImpl implements the DockerClient interface
type DockerClientImpl struct {
	client *client.Client
}

// NewDockerClient creates a new DockerClientImpl
func NewDockerClient() (*DockerClientImpl, error) {
	// Create Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}

	return &DockerClientImpl{
		client: cli,
	}, nil
}

// ExecuteCode executes code in a Docker container
func (dc *DockerClientImpl) ExecuteCode(execution *domain.CodeExecution) error {
	ctx := context.Background()

	// Create a temporary directory for code files
	tempDir, err := os.MkdirTemp("", "code-execution-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Write code to file
	codeFile, err := dc.writeCodeFile(tempDir, execution.Language, execution.Code)
	if err != nil {
		return fmt.Errorf("failed to write code file: %w", err)
	}

	// Write input to file if provided
	var inputFile string
	if execution.Input != "" {
		inputFile, err = dc.writeInputFile(tempDir, execution.Input)
		if err != nil {
			return fmt.Errorf("failed to write input file: %w", err)
		}
	}

	// Get Docker image and command for the language
	image, cmd := dc.getImageAndCommand(execution.Language, codeFile, inputFile)

	// Create container
	resp, err := dc.client.ContainerCreate(
		ctx,
		&container.Config{
			Image: image,
			Cmd:   cmd,
			Tty:   false,
		},
		&container.HostConfig{
			Binds: []string{
				fmt.Sprintf("%s:/code", tempDir),
			},
			Resources: container.Resources{
				Memory:    256 * 1024 * 1024, // 256 MB
				CPUPeriod: 100000,
				CPUQuota:  50000, // 0.5 CPU
			},
			NetworkMode: "none", // Disable network
		},
		nil,
		nil,
		"",
	)
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}
	defer dc.client.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})

	// Start container
	if err := dc.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}

	// Set timeout for execution
	timeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Wait for container to finish
	statusCh, errCh := dc.client.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	var statusCode int64
	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("error waiting for container: %w", err)
		}
	case status := <-statusCh:
		statusCode = status.StatusCode
	}

	// Get container logs
	out, err := dc.client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return fmt.Errorf("failed to get container logs: %w", err)
	}
	defer out.Close()

	// Read stdout and stderr
	var stdout, stderr bytes.Buffer
	_, err = stdcopy.StdCopy(&stdout, &stderr, out)
	if err != nil {
		return fmt.Errorf("failed to read container output: %w", err)
	}

	// Update execution with results
	execution.Output = stdout.String()
	execution.Error = stderr.String()
	execution.ExitCode = int(statusCode)

	return nil
}

// writeCodeFile writes code to a file
func (dc *DockerClientImpl) writeCodeFile(dir string, language domain.Language, code string) (string, error) {
	var filename string
	switch language {
	case domain.LanguagePython:
		filename = "main.py"
	case domain.LanguageJavaScript:
		filename = "main.js"
	case domain.LanguageGo:
		filename = "main.go"
	case domain.LanguageRuby:
		filename = "main.rb"
	case domain.LanguageJava:
		filename = "Main.java"
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	path := filepath.Join(dir, filename)
	err := os.WriteFile(path, []byte(code), 0644)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// writeInputFile writes input to a file
func (dc *DockerClientImpl) writeInputFile(dir string, input string) (string, error) {
	filename := "input.txt"
	path := filepath.Join(dir, filename)
	err := os.WriteFile(path, []byte(input), 0644)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// getImageAndCommand returns the Docker image and command for the language
func (dc *DockerClientImpl) getImageAndCommand(language domain.Language, codeFile, inputFile string) (string, []string) {
	var image string
	var cmd []string

	switch language {
	case domain.LanguagePython:
		image = "python:3.9-slim"
		cmd = []string{"python", fmt.Sprintf("/code/%s", codeFile)}
	case domain.LanguageJavaScript:
		image = "node:16-slim"
		cmd = []string{"node", fmt.Sprintf("/code/%s", codeFile)}
	case domain.LanguageGo:
		image = "golang:1.21-alpine"
		cmd = []string{"go", "run", fmt.Sprintf("/code/%s", codeFile)}
	case domain.LanguageRuby:
		image = "ruby:3.0-slim"
		cmd = []string{"ruby", fmt.Sprintf("/code/%s", codeFile)}
	case domain.LanguageJava:
		image = "openjdk:17-slim"
		cmd = []string{
			"/bin/sh",
			"-c",
			fmt.Sprintf("cd /code && javac %s && java Main", codeFile),
		}
	}

	// Add input redirection if input file is provided
	if inputFile != "" {
		cmdStr := strings.Join(cmd, " ")
		cmdStr = fmt.Sprintf("%s < /code/%s", cmdStr, inputFile)
		cmd = []string{"/bin/sh", "-c", cmdStr}
	}

	return image, cmd
}

// pullImageIfNeeded pulls the Docker image if it doesn't exist locally
func (dc *DockerClientImpl) pullImageIfNeeded(ctx context.Context, image string) error {
	// Check if image exists locally
	_, _, err := dc.client.ImageInspectWithRaw(ctx, image)
	if err == nil {
		// Image exists locally
		return nil
	}

	// Pull image
	reader, err := dc.client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()

	// Wait for pull to complete
	_, err = io.Copy(io.Discard, reader)
	return err
}
