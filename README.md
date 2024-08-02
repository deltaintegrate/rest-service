# Holiday Service

This project is a Go service that provides information about holidays using a public API. The service is contained within a Docker container and supports responses in both JSON and XML formats. 
The service allows filtering holidays by type and date range.

## Project Structure

- `main.go`: The main file that initializes and starts the server.
- `handlers/holiday.go`: Defines routes and handles HTTP requests.
- `models/holiday.go`: Defines data models used in the service.
- `services/holiday_service.go`: Implements the logic to fetch and filter holidays from the external API.
- `utils/logger.go`: Configures and handles logging using `logrus`.
- `tests/holiday_test.go`: Contains unit tests for the service.
- `Dockerfile`: Configuration file for building the Docker image of the service.

## Installation

### Requirements

- Go 1.22
- Docker

### Clone the Repository

```sh
git clone <REPOSITORY_URL>
cd holiday-service
