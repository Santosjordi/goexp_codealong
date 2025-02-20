# Golang REST API

This repository is a Golang REST API structured following the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) guidelines. It provides a modular, scalable, and maintainable project architecture for building RESTful services in Go.

## Project Structure

```
├── api             # OpenAPI/Swagger specifications, protocol definitions, and API definitions
├── cmd             # Main applications of the project (entry points)
│   ├── server      # The main HTTP server application
│   ├── worker      # Background processing worker (if applicable)
│   └── cli         # CLI application (if applicable)
├── configs         # Configuration files for different environments
├── internal        # Private application and library code
│   ├── handlers    # HTTP request handlers (controllers)
│   ├── services    # Business logic layer
│   ├── repository  # Database access layer
│   ├── middlewares # Custom middleware implementations
│   ├── models      # Data models and domain logic
│   ├── utils       # Utility functions and helpers
│   └── app         # Application initialization and dependency injection
├── pkg             # Public packages that can be used by external projects
├── migrations      # Database migration files
├── scripts         # Deployment and automation scripts
├── test            # Additional test data and integration tests
├── docs            # Documentation files
├── web             # Frontend or web-related files (if applicable)
├── .env            # Environment variable file
├── .gitignore      # Git ignore file
├── Dockerfile      # Docker build file
├── Makefile        # Makefile for common commands
├── go.mod          # Go module file
├── go.sum          # Go dependencies checksum file
└── README.md       # Project documentation
```

## Getting Started

### Prerequisites

- [Golang](https://go.dev/dl/) 1.21+
- [Docker](https://www.docker.com/) (optional for containerized deployment)
- [Make](https://www.gnu.org/software/make/) (optional for task automation)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/yourproject.git
   cd yourproject
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

### Running the API

#### Using Go
```sh
 go run cmd/server/main.go
```

#### Using Docker
```sh
docker build -t yourproject .
docker run -p 8080:8080 yourproject
```

#### Using Makefile
```sh
make run
```

### Configuration

The application configuration is managed through environment variables and config files inside the `configs/` directory. You can define variables in a `.env` file.

### API Documentation

The API documentation is available in the `api/` directory as an OpenAPI specification. Use tools like Swagger UI to visualize it.

## Testing

Run unit tests using:
```sh
go test ./...
```

## Contributing

1. Fork the repository
2. Create a new feature branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

