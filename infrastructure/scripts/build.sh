#!/bin/bash

# Build the frontend application
cd frontend
npm install
ng build --prod

# Build the backend application
cd ../backend
go mod tidy
go build -o todo-app ./cmd/main.go

# Build Docker images
docker build -t todo-app-frontend -f ../Dockerfile.frontend ../frontend
docker build -t todo-app-backend -f ../Dockerfile.backend .

# Clean up
cd ..
echo "Build completed successfully."