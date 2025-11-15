# video-vault

> A RESTful API built with Go and the Gin framework for managing a personal collection of saved videos.

## About the Project

Video Vault provides a backend platform for managing your saved videos in one place. It allows for secure user authentication, video management, and organization through tags.

Key features include:
*   **User Authentication:** Secure your platform with user authentication to ensure only authorized users can access and manage their content.
*   **Video CRUD Operations:** Perform Create, Read, Update, and Delete operations on saved videos.
*   **Tagging System:** Create and assign tags to videos for better categorization and filtering.
*   **Search and Filters:** Implement powerful search and filtering capabilities to help users find specific videos quickly.
*   **Authorization and Permissions:** Control access to various API endpoints based on user roles and permissions.

## Tech Stack

*   [Go](https://golang.org/)
*   [Gin](https://gin-gonic.com/)
*   [Docker Compose](https://docs.docker.com/compose/) (for database setup)

## Usage

Below are the instructions for you to set up and run the project locally.

### Prerequisites

You need to have the following software installed:

*   [Go](https://golang.org/dl/)
*   [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

### Installation and Setup

Follow the steps below:

1.  **Clone the repository**
    ```bash
    git clone https://github.com/luizvilasboas/video-vault.git
    ```

2.  **Navigate to the project directory**
    ```bash
    cd video-vault
    ```

3.  **Install Go dependencies**
    ```bash
    go mod tidy
    ```

4.  **Set up the database**
    Start the required database service using Docker Compose.
    ```bash
    docker-compose up -d
    ```

### Workflow

To run the application server, execute the following command:
```bash
go run main.go
```
By default, the API will be accessible at `http://localhost:8080`.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
