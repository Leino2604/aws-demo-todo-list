# README-local.md

# To-Do List Application - Local Development Setup

This document provides instructions for running the To-Do List application locally using Docker and Angular.

## Prerequisites

- Docker and Docker Compose installed on your machine.
- Node.js and npm installed (for Angular frontend).
- Go installed (for backend development).

## Running the Application Locally

1. **Clone the Repository**

   Clone the repository to your local machine:

   ```
   git clone <repository-url>
   cd todo-app
   ```

2. **Set Up Environment Variables**

   Create a `.env` file in the root directory and add the following variables:

   ```
   PASSWORD=your_password_here
   ```

   Replace `your_password_here` with the desired password for authentication.

3. **Build the Docker Images**

   Run the following command to build the Docker images for both frontend and backend:

   ```
   ./infrastructure/scripts/build.sh
   ```

4. **Run the Application**

   Use Docker Compose to start the application:

   ```
   docker-compose up
   ```

   This command will start both the frontend and backend services.

5. **Access the Application**

   Open your web browser and navigate to `http://localhost:4200` to access the To-Do List application.

## Local Development Features

- **In-Memory Database**: The application uses an in-memory database for storing to-do items. Data will be lost when the application is terminated.
- **Authentication**: Users must enter a password to access the application. The password is validated against the one stored in the `.env` file.
- **Task Management**: Users can add, remove, and start tasks, with estimated times displayed and countdown functionality.

## Auto-Redirect

The application will automatically redirect to the password entry page every 30 minutes to ensure security.

## Stopping the Application

To stop the application, press `CTRL+C` in the terminal where Docker Compose is running. This will terminate all services.

## Troubleshooting

- Ensure that Docker is running before executing the commands.
- Check the console for any error messages if the application does not start as expected.

## Conclusion

You are now set up to run the To-Do List application locally. For deployment instructions to AWS, please refer to `README-aws.md`.