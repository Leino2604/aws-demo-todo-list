# README for AWS Deployment

## Overview
This document provides instructions for deploying the To-Do List application to AWS using CloudFormation. The application consists of a Go backend, an Angular frontend, and uses AWS DynamoDB for data storage.

## Prerequisites
- AWS Account
- AWS CLI installed and configured
- Node.js and Angular CLI installed for frontend setup
- Go installed for backend setup

## Deployment Steps

### 1. Clone the Repository
Clone the repository to your local machine:
```bash
git clone <repository-url>
cd todo-app
```

### 2. Build the Frontend
Navigate to the frontend directory and install dependencies:
```bash
cd frontend
npm install
```
Build the Angular application:
```bash
ng build --prod
```

### 3. Deploy the Backend
Navigate to the backend directory and install dependencies:
```bash
cd ../backend
go mod tidy
```

### 4. Configure AWS CloudFormation
Edit the `deploy/cloudformation.yaml` file to set any necessary parameters, such as DynamoDB table names or IAM roles.

### 5. Deploy Using CloudFormation
Use the AWS CLI to deploy the CloudFormation stack:
```bash
aws cloudformation create-stack --stack-name todo-app-stack --template-body file://deploy/cloudformation.yaml --capabilities CAPABILITY_IAM
```

### 6. Monitor the Deployment
You can monitor the progress of the stack creation in the AWS Management Console under CloudFormation.

### 7. Access the Application
Once the stack is created successfully, you will receive an output with the URL of the deployed application. Use this URL to access your To-Do List application.

## Cost Management
To ensure that your AWS costs remain low:
- Use the free tier services where applicable.
- Delete the CloudFormation stack when you are done testing to avoid incurring charges:
```bash
aws cloudformation delete-stack --stack-name todo-app-stack
```

## Local Development
For local development, refer to the `README.md` in the backend and frontend directories for instructions on running the application locally.

## Conclusion
This guide provides a step-by-step approach to deploying your To-Do List application on AWS. Ensure to monitor your usage to keep costs within the desired limits.