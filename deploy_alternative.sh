#!/bin/bash

# Alternative deployment script for App Engine Flex
# This uses Cloud Build which doesn't require app-engine-go component

set -e

echo "=========================================="
echo "SocialAI - Alternative App Engine Deployment"
echo "=========================================="
echo ""

PROJECT=$(gcloud config get-value project)
if [ -z "$PROJECT" ]; then
    echo "❌ Error: No Google Cloud project set"
    exit 1
fi

echo "Project: $PROJECT"
echo ""

# Method 1: Try direct deployment (might work with custom runtime)
echo "Attempting Method 1: Direct gcloud app deploy..."
if gcloud app deploy app.yaml --quiet 2>&1 | tee /tmp/deploy.log; then
    echo "✅ Deployment successful!"
    exit 0
fi

# Check if it's the component error
if grep -q "app-engine-go" /tmp/deploy.log; then
    echo ""
    echo "⚠️  Direct deployment failed due to missing component"
    echo "Trying Method 2: Cloud Build deployment..."
    echo ""
    
    # Method 2: Use Cloud Build to build and deploy
    # First, build the image
    IMAGE_NAME="gcr.io/${PROJECT}/socialai"
    
    echo "Building Docker image with Cloud Build..."
    gcloud builds submit --tag "$IMAGE_NAME" .
    
    if [ $? -eq 0 ]; then
        echo "✅ Image built successfully"
        echo ""
        echo "To deploy the image to App Engine Flex, you can:"
        echo "1. Use the Cloud Console to deploy"
        echo "2. Or update app.yaml to reference the image"
        echo ""
        echo "Image location: $IMAGE_NAME"
    else
        echo "❌ Cloud Build failed"
        exit 1
    fi
else
    echo "❌ Deployment failed with a different error"
    cat /tmp/deploy.log
    exit 1
fi


