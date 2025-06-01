#!/bin/bash

# Set the AWS region and stack name
AWS_REGION="us-east-1"
STACK_NAME="todo-app-stack"

# Check if the necessary CloudFormation templates exist
if [ ! -f "../cloudformation/main.yaml" ]; then
  echo "Main CloudFormation template not found!"
  exit 1
fi

if [ ! -f "../cloudformation/dynamodb.yaml" ]; then
  echo "DynamoDB CloudFormation template not found!"
  exit 1
fi

if [ ! -f "../cloudformation/lambda.yaml" ]; then
  echo "Lambda CloudFormation template not found!"
  exit 1
fi

# Deploy the CloudFormation stack
echo "Deploying CloudFormation stack..."
aws cloudformation deploy \
  --template-file ../cloudformation/main.yaml \
  --stack-name $STACK_NAME \
  --region $AWS_REGION \
  --capabilities CAPABILITY_IAM

# Output the stack information
echo "Deployment complete. Stack name: $STACK_NAME"