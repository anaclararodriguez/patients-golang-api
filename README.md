# patients-golang-api
A simple API built with Golang that provides endpoints to persist patient's info.
## Table of Contents
- [Requirements](#requirements)
- [Environment Variables](#environment-variables)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Installation](#installation)
---
## Requirements
- Go version 1.20 or higher.
- A MySQL database.
  
## Environment Variables

The application requires the following environment variables:

- `MYSQL_PASSWORD`: The password for the MySQL user.
- `MYSQL_ROOT_PASSWORD`: The root password for MySQL.
- `MYSQL_USER`: The user to connect to the MySQL database.
- `MYSQL_HOST`: The host where the MySQL database is running.
- `MYSQL_PORT`: The port on which MySQL is running (default is set to `3306`).
- `MYSQL_DB`: The name of the database to connect to.


- `SERVER_PORT`: The port on which the application server will run (default is set to `8080`).


- `FROM_ADDRESS`: The sender's email address.
- `SMTP_SERVER`: The SMTP server to send emails.
- `SMTP_PORT`: The port of the SMTP server.
- `SMTP_USER`: The SMTP user (usually the email address).
- `SMTP_PASS`: The SMTP password or app-specific password.



## API Endpoints

| Endpoint             | Method | Description                                      | Request Parameters                              | Response                                                                 |
|----------------------|--------|--------------------------------------------------|------------------------------------------------|-------------------------------------------------------------------------|
| `/`                  | GET    | Retrieves a simple message confirming the API is up. | None                                           | `{ "message": "patients-golang-api is up and running !" }`                              |
| `/patients`          | GET    | Retrieves a list of all patients.                | None                                           | JSON array of patients |
| `/patients`          | POST   | Creates a new patient.                           | `name`, `email`, `address`, `contact_number`, `document_photo` (all required) | `{ "message": "Patient created successfully", "patient_id": 1 }` |
| `/{anything:.*}`     | GET    | Handles invalid paths (404).                     | None                                           | `{ "error": "Invalid path" }`                                               |

## Project structure

```plaintext
patients-golang-api
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── db
│   │   └── db.go
│   └── handler
│       ├── patient.go
│       ├── response.go
│       ├── root.go
│       └── server.go
│   ├── logic
│       ├── notifier.go
│   │   └── logic.go
│   └── model
│       └── model.go
├── postman
│   └── patients-golang-api.postman_collection.json
├── .env.example
├── docker-compose.yml
├── Dockerfile
└── init.sql

```
- `cmd/api`: Contains the entry point of the application.
- ``internal/config``: Handles configuration loading.
- ``internal/db``: Contains database connection and related functions.
- ``internal/handler``: Contains the HTTP handlers and related functions.
- ``internal/logic``: Includes business logic and utilities.
- ``internal/model``: Defines the data models used across the project.
- `postman`: Contains the Postman collection for API testing and documentation.
- `.env.example`: Provides an example of the environment variables.
- `docker-compose.yml`: Specifies the services to run in Docker containers, along with network and environment variable configurations.
- `Dockerfile`: Defines how the application is built and packaged into a Docker container.
- `init.sql`: Contains the initial SQL script to set up the database.



## Installation

1. Clone the repository:
   ```bash
   git clone <repository-path>
   ```

2. Navigate to the project directory:
   ```bash
   cd patients-golang-api
   ```

3. Set up the `.env` file in the root directory. Example:
   ```env
    MYSQL_PASSWORD=your_mysql_password
    MYSQL_ROOT_PASSWORD=your_mysql_root_password
    MYSQL_USER=your_mysql_user
    MYSQL_HOST=your_mysql_host
    MYSQL_PORT=your_mysql_port
    MYSQL_DB=your_database_name

    SERVER_PORT=your_server_port

    FROM_ADDRESS=your_email_address
    SMTP_SERVER=your_smtp_server
    SMTP_PORT=your_smtp_port
    SMTP_USER=your_smtp_user
    SMTP_PASS=your_smtp_password
   ```

4. Start the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

5. The application should now be running at `http://localhost:<SERVER_PORT>`.
