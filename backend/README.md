# Backend README.md

# Todo App Backend

This document provides instructions for setting up and running the backend of the Todo App, which is built using Go. The backend includes functionalities for user authentication, task management, and Pomodoro timer features.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Local Development](#local-development)
- [AWS Deployment](#aws-deployment)
- [API Endpoints](#api-endpoints)

## Prerequisites

- Go 1.16 or later
- AWS account (for deployment)
- AWS CLI configured with your credentials (for deployment)

## Local Development

To run the backend locally, follow these steps:

1. Clone the repository:
   ```
   git clone <repository-url>
   cd todo-app/backend
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run main.go
   ```

4. The server will start on `localhost:8080`. You can use tools like Postman or curl to interact with the API.

### In-Memory Database

For local development, the application uses an in-memory database. This means that all data will be lost when the application is terminated.

## AWS Deployment

To deploy the backend to AWS, you can use the provided CloudFormation template. Follow these steps:

1. Navigate to the `deploy` directory:
   ```
   cd ../deploy
   ```

2. Deploy the CloudFormation stack:
   ```
   aws cloudformation create-stack --stack-name todo-app-backend --template-body file://cloudformation.yaml --capabilities CAPABILITY_IAM
   ```

3. Once the stack is created, note the API Gateway endpoint provided in the CloudFormation outputs.

4. Update the application configuration to point to the DynamoDB instance created by the CloudFormation stack.

## API Endpoints

### Authentication

- **POST /auth/signup**: Create a new user.
- **POST /auth/signin**: Sign in an existing user.

### Tasks

- **GET /tasks**: Retrieve all tasks.
- **POST /tasks**: Create a new task.
- **PUT /tasks/{id}**: Update an existing task.
- **DELETE /tasks/{id}**: Delete a task.

### Pomodoro

- **POST /pomodoro/start**: Start the Pomodoro timer.
- **POST /pomodoro/stop**: Stop the Pomodoro timer.

## License

This project is licensed under the MIT License. See the LICENSE file for details.