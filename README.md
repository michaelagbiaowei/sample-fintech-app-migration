# Sample Fintech App Migration

This is a sample fintech web application built with Go. It provides basic functionality for user authentication, account management, and file uploads.

## Features

- User registration and login
- Account balance management (deposit and withdraw)
- File upload to DigitalOcean Spaces
- Responsive web interface

## Tech Stack

- Go (Golang) for backend
- HTML templates for frontend
- AWS RDS (PostgreSQL) for database
- DigitalOcean Spaces for file storage

## Prerequisites

- Go 1.16 or higher
- AWS RDS (PostgreSQL)
- DigitalOcean account (for Spaces)

## Environment Variables

The application uses the following environment variables:

- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port
- `DB_USER`: PostgreSQL user
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name
- `SESSION_KEY`: Secret key for session management
- `DO_SPACE_KEY`: DigitalOcean Spaces access key
- `DO_SPACE_SECRET`: DigitalOcean Spaces secret key
- `DO_SPACE_ENDPOINT`: DigitalOcean Spaces endpoint
- `DO_SPACE_REGION`: DigitalOcean Spaces region
- `DO_SPACE_BUCKET`: DigitalOcean Spaces bucket name
- `PORT`: Application port (default: 8080)

## Running Locally

1. Clone the repository
2. Set the required environment variables
3. Run the application:

```bash
go run main
 ```

## Deployment
### DigitalOcean App Platform
This project is configured to automatically deploy to DigitalOcean App Platform using GitHub Actions. The workflow is triggered on pushes to the main branch.
To set up the deployment:

Configure the necessary secrets in your GitHub repository settings:

`DIGITALOCEAN_ACCESS_TOKEN`
Any other required environment variables


Push changes to the main branch to trigger the deployment workflow.

## Migrating to Azure App Service
We are planning to migrate this Application to Azure App Service. 

Steps for migration will include:

- Create an Azure App Service plan
- Set up an Azure Database for PostgreSQL
- Configure Azure Blob Storage to replace DigitalOcean Spaces
- Update environment variables in Azure App Service configuration
- Set up continuous deployment from GitHub to Azure App Service

(Detailed migration steps will be added once the migration is complete)

### Files in Directory

sample-fintech-app-migration/

    ├── internal/
    │   ├── auth/
    │   │   └── auth.go
    │   ├── handlers/
    │   │   ├── deposit.go
    │   │   ├── home.go
    │   │   ├── login.go
    │   │   ├── signup.go
    │   │   ├── upload.go
    │   │   └── withdraw.go
    │   ├── models/
    │   │   ├── account.go
    │   │   └── user.go
    │   ├── storage/
    │   │   └── postgres.go
    │   └── upload/
    │       └── digitalocean.go
    ├── static/
    │   └── style.css
    ├── templates/
    │   ├── deposit.html
    │   ├── home.html
    │   ├── layout.html
    │   ├── login.html
    │   ├── signup.html
    │   ├── upload.html
    │   └── withdraw.html
    ├── tests/
    │   ├── auth_test.go
    ├── .env
    ├── go.mod
    └── go.sum

### Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
License

This project is licensed under the MIT License.