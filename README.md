# GitHub Unfollowers

## Description

This Go program helps you identify GitHub users who you are following but are not following you back. It leverages the GitHub API to fetch your followers and following lists and then compares them to find the "unfollowers." This is a useful tool for cleaning up your following list and ensuring you're following users who are also engaged with your content.

## Installation

To get started with this project, follow these steps:

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

2.  **Install dependencies:**

    ```bash
    go mod download
    ```

3.  **Set up environment variables:**

    *   Create a `.env` file in the root directory.
    *   Add your GitHub token to the `.env` file:

        ```
        GITHUB_TOKEN=<YOUR_GITHUB_TOKEN>
        ```

    *   You can generate a personal access token (PAT) from GitHub's settings page (`Settings` > `Developer settings` > `Personal access tokens`). The token needs the `read:user` scope.

## Usage

Run the program with the following command:

```bash
go run cmd/main.go
```

The program will output a list of GitHub usernames who you follow but who do not follow you back.

## Features

*   **Identifies Unfollowers:** Accurately determines which users you follow are not following you back.
*   **GitHub API Integration:** Uses the official GitHub API for reliable data fetching.
*   **Easy Setup:** Simple installation and configuration process.
*   **Clear Output:** Presents the list of unfollowers in a clean and readable format.

## Technologies Used

| Technology | Description                                      |
|------------|--------------------------------------------------|
| Go         | Programming language used for the application. |
| GitHub API | For fetching user and follower data.            |
| `go-github` | Go library for interacting with the GitHub API. |
| `godotenv`  | For loading environment variables from a `.env` file. |
| `oauth2`  | For authenticating with the Github API.          |

## Contributing

Contributions are welcome! Here's how you can contribute:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Make your changes and commit them with descriptive messages.
4.  Submit a pull request.

### Setting up the development environment:

1.  Install Go: [https://go.dev/doc/install](https://go.dev/doc/install)
2.  Clone the repository:

    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

3.  Run tests:

    ```bash
    go test ./...
    ```

## License

This project does not have a license.

[![Built with Dokugen](https://img.shields.io/badge/Built%20with-Dokugen-brightgreen)](https://github.com/samueltuoyo15/Dokugen)
