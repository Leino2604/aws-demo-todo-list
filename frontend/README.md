# Frontend README.md

# To-Do List Application - Frontend

This is the frontend part of the To-Do List application built with Angular. It provides a user interface for managing tasks, user authentication, and Pomodoro timer functionalities.

## Features

- **To-Do List**: Users can create, update, and delete tasks with estimated times.
- **Pomodoro Mode**: Users can start a Pomodoro timer with customizable work and rest durations.
- **User Authentication**: Simple sign-in and sign-up functionalities.

## Prerequisites

- Node.js (version 14 or higher)
- Angular CLI (version 12 or higher)

## Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd todo-app/frontend
   ```

2. **Install Dependencies**:
   ```bash
   npm install
   ```

3. **Run the Application Locally**:
   ```bash
   ng serve
   ```
   Navigate to `http://localhost:4200/` in your browser to access the application.

## Environment Configuration

- The application uses environment files located in `src/environments/`.
- Modify `environment.ts` for local development settings.
- Modify `environment.prod.ts` for production settings.

## Components Overview

- **AppComponent**: The root component of the application.
- **TodoComponent**: Manages the to-do list functionalities.
- **AuthComponent**: Handles user authentication.
- **PomodoroComponent**: Manages the Pomodoro timer functionalities.

## Build for Production

To build the application for production, run:
```bash
ng build --prod
```
The output will be in the `dist/` directory.

## Additional Notes

- Ensure that the backend server is running and accessible for the frontend to function correctly.
- For any issues or contributions, please refer to the main project repository.