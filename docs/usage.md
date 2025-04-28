# Usage Guide

This guide provides instructions for using the Local AI Agent System.

## Getting Started

After installing the system, you can access the web interface at:

```
http://localhost:5000
```

The interface consists of two main sections:

1. **Tasks**: For submitting and managing tasks
2. **Results**: For viewing task results

## Task Management

### Creating a Task

1. Navigate to the Tasks page
2. Fill in the task form:
   - **Title**: A descriptive name for the task
   - **Description** (optional): Additional details about the task
   - **Input**: The task input, which can be:
     - Code to execute
     - Natural language prompt for the AI
     - URL to browse
     - File operation instructions
3. Click "Create Task" to submit the task

### Task Types

The system supports several types of tasks:

#### Code Execution

To execute code, provide the code in the input field and specify the language:

```
language: python
print("Hello, world!")
```

Supported languages:
- Python
- JavaScript (Node.js)
- Go
- Ruby
- Java

#### AI Reasoning

To use the AI for reasoning or planning, provide a natural language prompt:

```
language: ai
Explain how to implement a binary search algorithm in Python.
```

#### Web Browsing

To browse a web page, provide a URL and optional instructions:

```
language: web
url: https://example.com
extract: h1, p
```

#### Filesystem Operations

To perform filesystem operations, provide instructions:

```
language: fs
operation: write
path: example.txt
content: Hello, world!
```

Supported operations:
- read
- write
- list
- delete
- mkdir

### Monitoring Tasks

The Tasks page displays a list of all tasks with their status:

- **Pending**: Task is waiting to be processed
- **Running**: Task is currently being processed
- **Completed**: Task has been successfully completed
- **Failed**: Task encountered an error

The list automatically refreshes every 5 seconds.

## Viewing Results

### Result List

The Results page displays a list of completed tasks. Click on a task to view its details.

### Result Details

The result details view shows:

- Task information (title, description, status)
- Result output with syntax highlighting
- Input used for the task
- Timestamps for creation and completion

For code execution tasks, the result includes:
- Output of the code
- Execution time
- Exit code

For AI reasoning tasks, the result includes:
- AI response
- Tokens used
- Model used

For web browsing tasks, the result includes:
- Page content
- Extracted data
- Screenshot (if requested)

For filesystem operations, the result includes:
- Operation status
- File content (for read operations)
- File list (for list operations)

## Examples

### Example 1: Python Code Execution

Input:
```
language: python
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

for i in range(10):
    print(fibonacci(i))
```

Result:
```
0
1
1
2
3
5
8
13
21
34
```

### Example 2: AI Reasoning

Input:
```
language: ai
Explain the concept of recursion in programming.
```

Result:
```
Recursion in programming is a technique where a function calls itself to solve a problem. It's based on the principle of solving a complex problem by breaking it down into simpler instances of the same problem.

Key components of recursion:
1. Base case: A condition that stops the recursion
2. Recursive case: The function calling itself with modified parameters

Example in Python:
```python
def factorial(n):
    # Base case
    if n == 0 or n == 1:
        return 1
    # Recursive case
    else:
        return n * factorial(n-1)
```

Recursion is powerful for problems like tree traversal, graph algorithms, and divide-and-conquer strategies. However, it can lead to stack overflow errors if not implemented carefully, as each recursive call adds a new frame to the call stack.

Iterative solutions often exist as alternatives to recursion and may be more efficient in terms of memory usage.
```

### Example 3: Web Browsing

Input:
```
language: web
url: https://example.com
extract: h1, p
```

Result:
```
Extracted data:
h1: Example Domain
p: This domain is for use in illustrative examples in documents. You may use this domain in literature without prior coordination or asking for permission.
```

### Example 4: Filesystem Operation

Input:
```
language: fs
operation: write
path: example.txt
content: Hello, world!
```

Result:
```
Operation: write
Path: example.txt
Success: true
```
