# Social AI Backend

A Go-based backend service for a social media application with image/video upload capabilities, user authentication, and search functionality.

## Features

- **User Authentication**: JWT-based authentication with signup and signin endpoints
- **Post Management**: Upload posts with media files (images/videos)
- **Search**: Search posts by user or keywords
- **Storage**: 
  - Elasticsearch for data storage and search
  - Google Cloud Storage (GCS) for media files
- **Deployment**: Ready for Google App Engine Flex deployment

## Architecture

- **Backend**: Go with Gorilla Mux router
- **Database**: Elasticsearch
- **Storage**: Google Cloud Storage
- **Authentication**: JWT tokens
- **Deployment**: Google App Engine Flex

## Project Structure

```
.
├── backend/          # Backend services (Elasticsearch, GCS)
├── conf/            # Configuration files
├── constants/       # Application constants
├── handler/         # HTTP handlers
├── model/           # Data models
├── service/         # Business logic
├── util/            # Utility functions
├── main.go          # Application entry point
├── app.yaml         # App Engine configuration
└── Dockerfile       # Docker build configuration
```

## Configuration

### Setting Up Configuration

1. Copy the example configuration file:
   ```bash
   cp conf/deploy.yaml.example conf/deploy.yaml
   ```

2. Edit `conf/deploy.yaml` with your actual credentials:
   ```yaml
   elasticsearch:
     address: "http://your-elasticsearch:9200"
     username: "your-username"
     password: "your-password"

   gcs:
     bucket: "your-bucket-name"

   token:
     secret: "your-jwt-secret"
   ```
   
## API Endpoints

### Public Endpoints
- `POST /signup` - User registration
- `POST /signin` - User login (returns JWT token)

### Protected Endpoints (Require JWT Token)
- `POST /upload` - Upload a post with media file
- `GET /search?keywords=<text>` - Search posts by keywords
- `GET /search?user=<username>` - Search posts by user
- `DELETE /post/{id}` - Delete a post

## Elasticsearch Refresh Fixes

This version includes fixes for Elasticsearch near real-time search inconsistencies:
- Writes are immediately searchable with `Refresh("wait_for")`
- Reads refresh the index before searching to ensure consistency
- Prevents stale results and inconsistent query behavior

## Deployment

### To Google App Engine

```bash
gcloud app deploy app.yaml
```

### Local Development

```bash
go run main.go
```

## Requirements

- Go 1.25+
- Elasticsearch 7.x
- Google Cloud Storage bucket
- Google Cloud SDK (for deployment)

## License

MIT

