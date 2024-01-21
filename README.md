#  video-vault

This RESTful API is built using Go and the Gin framework, providing a plataform for managing your saved videos in one place. 

## Features

- **User Authentication:** Secure your platform with user authentication to ensure that only authorized users can access and manage their saved videos.

- **Video CRUD Operations:** Perform CRUD (Create, Read, Update, Delete) operations on saved videos. Add new videos, retrieve video details, update video information, and delete unwanted videos.

- **Search and Filters:** Implement powerful search and filtering capabilities to help users find specific videos quickly.

- **Authorization and Permissions:** Control access to various API endpoints based on user roles and permissions.

- **Create and set tags:** Create and set tags to various videos to better categorize them.

## Usage

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/)
- [Gin](https://github.com/gin-gonic/gin) - Golang web framework

### Installation

1. Clone the repository:

   ```
   git clone https://gitlab.com/olooeez/video-vault.git
   ```

2. Navigate to the project directory:

   ```
   cd video-vault
   ```

3. Install dependencies:

   ```
   go mod tidy
   ```

4. Set up your database and configure the database connection with `Docker Compose`:
  ```
  docker-compose up -d
  ```

5. Run the application:

   ```
   go run main.go
   ```

By default, the API will be accessible at `http://localhost:8080`.

## Contributing

If you wish to contribute to this project, feel free to open a merge request. We welcome all forms of contribution!

## License

This project is licensed under the [MIT License](https://gitlab.com/olooeez/video-vault/-/blob/main/LICENSE). Refer to the LICENSE file for more details.
