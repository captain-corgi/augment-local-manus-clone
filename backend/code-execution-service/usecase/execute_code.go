package usecase

import (
	"fmt"
	"time"

	"github.com/augment-local-manus-clone/backend/code-execution-service/domain"
	"github.com/google/uuid"
)

// ExecuteCodeInput represents the input for executing code
type ExecuteCodeInput struct {
	Code     string         `json:"code"`
	Language domain.Language `json:"language"`
	Input    string         `json:"input,omitempty"`
}

// ExecuteCodeUseCase handles code execution
type ExecuteCodeUseCase struct {
	dockerClient domain.DockerClient
}

// NewExecuteCodeUseCase creates a new instance of ExecuteCodeUseCase
func NewExecuteCodeUseCase(dockerClient domain.DockerClient) *ExecuteCodeUseCase {
	return &ExecuteCodeUseCase{
		dockerClient: dockerClient,
	}
}

// Execute executes code in a Docker container
func (uc *ExecuteCodeUseCase) Execute(input ExecuteCodeInput) (*domain.CodeExecution, error) {
	// Create a new code execution
	execution, err := domain.NewCodeExecution(input.Code, input.Language, input.Input)
	if err != nil {
		return nil, err
	}

	// Generate a unique ID
	execution.ID = fmt.Sprintf("exec_%s", uuid.New().String())

	// Record start time
	startTime := time.Now()

	// Execute the code
	err = uc.dockerClient.ExecuteCode(execution)
	if err != nil {
		return nil, err
	}

	// Calculate duration
	execution.Duration = time.Since(startTime).Seconds()

	return execution, nil
}
