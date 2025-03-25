# Go Real Estate Backend API

[![Go](https://img.shields.io/badge/Go-v1.21-blue)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-v1.9-critical)](https://github.com/gin-gonic/gin)
[![MySQL](https://img.shields.io/badge/MySQL-8.0-green)](https://www.mysql.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This repository contains the backend API for a real estate application, built using Go, the Gin web framework, and MySQL.

## Features

*   **User Authentication:** Secure user registration, login, and authentication.
*   **Hero Banner Management:**  Retrieval of a hero banner with a title, description, and associated images for the website's front page.
*   **API Documentation:** Clear and concise documentation to help developers integrate with the API.
*   **Database Migrations:** Automated database schema management using GORM.
*   **Configuration Management:** Using `.env` files for storing sensitive configuration information.

## Technologies Used

*   **Go:** Programming language
*   **Gin:** Web framework for building APIs
*   **GORM:** ORM for interacting with the MySQL database
*   **MySQL:** Relational database management system
*   **dotenv:** For loading environment variables from `.env` files
*   **Testify:** Testing suite for writing effective tests

## Prerequisites

Before you begin, ensure you have the following installed:

*   [Go](https://go.dev/dl/) (version 1.21 or higher)
*   [MySQL](https://www.mysql.com/downloads/) (version 8.0 or higher)

## Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/HariKrishnan-Ramesh/Go-Project-second.git
    cd Go-Project-second
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Create a `.env` file:**

    Create a `.env` file in the root directory of the project and configure the following environment variables:

    ```
    DB_USER=<your_mysql_username>
    DB_PASS=<your_mysql_password>
    DB_HOST=<your_mysql_host>
    DB_NAME=<your_mysql_database_name>
    DB_PORT=<your_mysql_port> # usually 3306
    PORT=<port_for_the_server> #usually 8080

    # Optional: Test database configuration
    TEST_DB_USER=<your_test_mysql_username>
    TEST_DB_PASS=<your_test_mysql_password>
    TEST_DB_HOST=<your_test_mysql_host>
    TEST_DB_NAME=<your_test_mysql_database_name>
    TEST_DB_PORT=<your_test_mysql_port>
    ```

4.  **Run database migrations:**

    The application automatically migrates the database schema on startup. No manual migration command needed.

## Running the Application

1.  **Start the server:**

    ```bash
    go run main.go
    ```

2.  **Access the API:**

    The API will be available at `http://localhost:<PORT>`, where `<PORT>` is the port specified in your `.env` file (default is 8080).

## API Endpoints

| Method | Endpoint          | Description                                                                                                                                                             |
| :----- | :---------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| GET    | `/api/herobanner` | Retrieves the hero banner data, including title, description, and a list of image URLs.  Returns a 404 if no hero banner is found.                                      |

## Testing

1.  **Configure the test database:**

    Ensure that you have configured the test database environment variables in your `.env` file.

2.  **Run the tests:**

    ```bash
    go test ./handlers
    ```

## Contributing

Contributions are welcome! Please follow these steps:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Make your changes and commit them with clear, concise commit messages.
4.  Push your branch to your fork.
5.  Create a pull request to the main repository.


## Contact

[HariKrishna Ramesh](mailto:harikrishnan.ramesh@example.com)

## Acknowledgements

*   Thanks to the Go community for providing excellent tools and resources.
*   Thanks to the developers of Gin and GORM for creating powerful and easy-to-use libraries.
