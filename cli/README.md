# CLI README

# To-Do List Application CLI

This document provides instructions for using the command-line interface (CLI) for the To-Do List application. The CLI allows you to interact with the application, manage tasks, and handle user authentication.

## Installation

To use the CLI, ensure you have Go installed on your machine. You can download it from the official Go website.

1. Clone the repository:
   ```
   git clone <repository-url>
   cd todo-app/cli
   ```

2. Build the CLI application:
   ```
   go build -o todo-cli main.go
   ```

3. Run the CLI application:
   ```
   ./todo-cli
   ```

## Usage

The CLI provides several commands to interact with the To-Do List application:

### Authentication

- **Sign Up**
  ```
  ./todo-cli signup --username <username> --password <password>
  ```
  This command creates a new user account.

- **Sign In**
  ```
  ./todo-cli signin --username <username> --password <password>
  ```
  This command allows an existing user to sign in.

### Task Management

- **Add Task**
  ```
  ./todo-cli add-task --name <task-name> --estimated-time <time-in-minutes>
  ```
  This command adds a new task to the to-do list with an estimated time.

- **List Tasks**
  ```
  ./todo-cli list-tasks
  ```
  This command lists all tasks in the to-do list.

- **Start Task**
  ```
  ./todo-cli start-task --task-id <task-id>
  ```
  This command starts the specified task and begins the countdown of the estimated time.

- **Complete Task**
  ```
  ./todo-cli complete-task --task-id <task-id>
  ```
  This command marks the specified task as completed.

### Pomodoro Timer

- **Start Pomodoro**
  ```
  ./todo-cli start-pomodoro --work-time <time-in-minutes> --rest-time <time-in-minutes> --cycles <number-of-cycles>
  ```
  This command starts the Pomodoro timer with specified work and rest times.

## Local Development

For local development, the CLI interacts with an in-memory database. All data will be lost when the application is terminated.

## AWS Deployment

For AWS deployment, ensure that the backend is running and connected to the DynamoDB instance. Refer to the deployment instructions in the `deploy/README.md` for more details.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.