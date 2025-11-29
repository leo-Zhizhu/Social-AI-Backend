# Google App Engine Deployment Guide

This guide will help you deploy the socialai project to Google App Engine using the gcloud CLI from your local terminal.

## Prerequisites

1. **Install Google Cloud SDK (gcloud CLI)**
   - Download from: https://cloud.google.com/sdk/docs/install
   - Or install via package manager:
     ```bash
     # macOS
     brew install google-cloud-sdk
     
     # Linux (Debian/Ubuntu)
     curl https://sdk.cloud.google.com | bash
     exec -l $SHELL
     ```

2. **Authenticate with Google Cloud**
   ```bash
   gcloud auth login
   ```

3. **Set your project**
   ```bash
   gcloud config set project YOUR_PROJECT_ID
   ```

4. **Enable required APIs**
   ```bash
   gcloud services enable appengine.googleapis.com
   gcloud services enable storage-component.googleapis.com
   ```

## Deployment Steps

### Step 1: Navigate to Project Directory
```bash
cd /home/HAL/go/src/socialai
```

### Step 2: Verify Configuration
Make sure `conf/deploy.yaml` contains the correct values:
- Elasticsearch address, username, and password
- GCS bucket name
- JWT token secret

### Step 3: Deploy to App Engine
```bash
gcloud app deploy app.yaml
```

This command will:
- Build your Go application
- Upload the application to Google Cloud
- Deploy it to App Engine Flex

### Step 4: Monitor Deployment
The deployment process will show progress. It may take 5-10 minutes for the first deployment.

### Step 5: Get Application URL
After deployment completes, get your app URL:
```bash
gcloud app browse
```

Or get the URL directly:
```bash
gcloud app describe --format="value(defaultHostname)"
```

## Important Notes

1. **Network Access**: Your Elasticsearch instance at `10.128.0.4:9200` must be accessible from App Engine. Make sure:
   - The Elasticsearch instance is in the same VPC network
   - Firewall rules allow App Engine to access it
   - Or use a public IP if Elasticsearch is configured for external access

2. **Service Account**: App Engine uses a default service account. Make sure it has:
   - Storage Admin role (for GCS access)
   - Network access to your Elasticsearch instance

3. **First Deployment**: The first deployment takes longer as it creates the App Engine environment.

## Troubleshooting

### Check Logs
```bash
gcloud app logs tail -s default
```

### View Service Status
```bash
gcloud app services list
```

### View Versions
```bash
gcloud app versions list
```

### Rollback if Needed
```bash
gcloud app versions list
gcloud app versions migrate PREVIOUS_VERSION
```

## Updating the Application

To deploy updates:
```bash
cd /home/HAL/go/src/socialai
gcloud app deploy app.yaml
```

## Configuration

The application reads configuration from `conf/deploy.yaml`. Make sure this file is included in your deployment (it should be by default).

## Cost Considerations

App Engine Flex has minimum costs:
- Minimum 1 instance running
- Charges for compute, network, and storage
- Consider using App Engine Standard for lower costs (requires code changes)


