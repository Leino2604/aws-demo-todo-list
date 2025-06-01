# To-Do List Application

This is a medium-level to-do list application built with Angular for the frontend and Go for the backend. The application allows users to manage their tasks with estimated times and includes authentication features.

## Features

- **To-Do List**: Users can add, remove, and manage tasks.
- **Estimated Times**: Each task can have an estimated time attached.
- **Task Countdown**: Users can start a task, and the estimated time will count down.
- **Authentication**: A password-protected entry point for the application.
- **Auto-Redirect**: Users will be redirected to the password entry page every 30 minutes.

## Tech Stack

- **Frontend**: Angular
- **Backend**: Go
- **Database**: AWS DynamoDB

## Local Development

### Prerequisites

- Node.js and npm
- Docker and Docker Compose
- Go

### Running Locally

1. Clone the repository:
   ```
   git clone <repository-url>
   cd todo-app
   ```

2. Set up environment variables:
   - Copy `.env.example` to `.env` and set your password and other secrets.

3. Build and run the application using Docker Compose:
   ```
   docker-compose up --build
   ```

4. Access the application at `http://localhost:4200`.

### Local Database

- The application uses an in-memory database for local development. Data will be lost when the application is terminated.

## AWS Deployment

### Prerequisites

- AWS Account
- AWS CLI configured with appropriate permissions

### Deployment Steps

1. Navigate to the infrastructure directory:
   ```
   cd infrastructure
   ```

2. Deploy the CloudFormation stack:
   ```
   ./scripts/deploy.sh
   ```

3. Follow the prompts to set up your AWS resources.

### Cost Management

- This application is designed to be cost-effective. Ensure that you monitor your AWS usage to keep costs under $0.5 for 2 days of running.

## Additional Documentation

- For local setup instructions, refer to `README-local.md`.
- For AWS deployment instructions, refer to `README-aws.md`.

## License

This project is licensed under the MIT License.