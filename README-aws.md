# README-aws.md

# To-Do List Application Deployment on AWS

This document provides instructions for deploying the To-Do List application to AWS using CloudFormation. The application consists of a frontend built with Angular, a backend written in Go, and uses AWS DynamoDB for data storage.

## Prerequisites

1. **AWS Account**: Ensure you have an active AWS account.
2. **AWS CLI**: Install and configure the AWS Command Line Interface (CLI) with your credentials.
3. **Docker**: Install Docker to build the application images.
4. **Node.js and npm**: Ensure Node.js and npm are installed for building the frontend.

## Deployment Steps

### Step 1: Clone the Repository

Clone the repository to your local machine:

```bash
git clone <repository-url>
cd todo-app
```

### Step 2: Build Docker Images

Navigate to the `infrastructure/scripts` directory and run the build script to create Docker images for the frontend and backend:

```bash
cd infrastructure/scripts
./build.sh
```

### Step 3: Configure Environment Variables

Create a `.env` file in the root directory based on the `.env.example` file. Make sure to set the necessary environment variables, including the password for authentication.

```bash
cp .env.example .env
```

Edit the `.env` file to include your secrets.

### Step 4: Deploy to AWS

Run the deployment script to create the necessary AWS resources using CloudFormation:

```bash
./deploy.sh
```

This script will deploy the following resources:
- AWS DynamoDB table for storing to-do items.
- AWS Lambda functions for the backend API.
- API Gateway to expose the Lambda functions.

### Step 5: Access the Application

Once the deployment is complete, you will receive an API endpoint URL. Use this URL to access the application. Open your browser and navigate to the frontend application.

### Step 6: Monitor Costs

To ensure that your AWS costs do not exceed $0.5 for 2 days, consider the following:
- Use the AWS Free Tier where applicable.
- Monitor your usage through the AWS Billing Dashboard.
- Delete the resources after testing to avoid incurring additional charges.

## Cleanup

To remove the deployed resources and avoid further charges, run the following command:

```bash
aws cloudformation delete-stack --stack-name <your-stack-name>
```

Replace `<your-stack-name>` with the name of the CloudFormation stack created during deployment.

## Conclusion

You have successfully deployed the To-Do List application on AWS. For local development instructions, refer to `README-local.md`. For any issues or questions, please refer to the main `README.md` or contact the project maintainers.