# To-Do List Application

This is a simple to-do list application built with Angular for the frontend and Go for the backend. The application allows users to manage tasks with estimated times, supports a Pomodoro timer, and includes user authentication features.

## Features

- **To-Do List**: Create, update, and delete tasks.
- **Estimated Times**: Attach estimated times to tasks and track them.
- **Pomodoro Mode**: Start a Pomodoro timer with customizable work and rest durations.
- **User Authentication**: Simple sign-in and sign-up functionality.

## Tech Stack

- **Frontend**: Angular
- **Backend**: Go
- **Database**: AWS DynamoDB

## Local Development

### Prerequisites

- Go (1.16 or later)
- Node.js and npm
- Angular CLI

### Setup

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd todo-app
   ```

2. **Backend Setup**:
   - Navigate to the backend directory:
     ```
     cd backend
     ```
   - Install dependencies:
     ```
     go mod tidy
     ```
   - Run the backend server:
     ```
     go run main.go
     ```

3. **Frontend Setup**:
   - Navigate to the frontend directory:
     ```
     cd ../frontend
     ```
   - Install dependencies:
     ```
     npm install
     ```
   - Run the Angular application:
     ```
     ng serve
     ```

4. **Access the Application**:
   Open your browser and go to `http://localhost:4200`.

### Local Database

For local development, the application uses an in-memory database. All data will be lost when the application is terminated.

## AWS Deployment

### Prerequisites

- AWS Account
- AWS CLI configured with appropriate permissions

### Deployment Steps

1. **Navigate to the deploy directory**:
   ```
   cd deploy
   ```

2. **Deploy using CloudFormation**:
   - Create a CloudFormation stack:
     ```
     aws cloudformation create-stack --stack-name todo-app-stack --template-body file://cloudformation.yaml --capabilities CAPABILITY_IAM
     ```

3. **Access the Application**:
   After the stack is created, you will receive an output with the URL to access your application.

### Cost Management

To keep costs under $0.5 for 2 days:
- Use the free tier for DynamoDB.
- Ensure that you terminate the stack after testing:
  ```
  aws cloudformation delete-stack --stack-name todo-app-stack
  ```

## Additional Information

For more detailed instructions on the backend, frontend, and CLI, please refer to their respective README files located in the `backend`, `frontend`, and `cli` directories.