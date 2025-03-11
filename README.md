# Ghostbuster: Your GitHub Follower Auditor 👻

Tired of following people who don't follow you back on GitHub? **Ghostbuster** is here to help! This simple Go program audits your GitHub following list and unfollows those who haven't reciprocated the follow. Free up your feed and keep your network clean!

## ✨ Features

-   **Easy Setup:** Just a few steps to get started.
-   **Automated Unfollowing:** Identifies and unfollows non-followers.
-   **Clean Output:** Clear messages indicating who was unfollowed and any errors encountered.
-   **Secure:** Uses your GitHub token for authentication.

## 🛠️ Installation

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

2.  **Set up your `.env` file:**

    -   Create a `.env` file in the project root.
    -   Add your GitHub token:

        ```
        GITHUB_TOKEN=YOUR_GITHUB_TOKEN
        ```

        Replace `YOUR_GITHUB_TOKEN` with your personal access token. You can generate one [here](https://github.com/settings/tokens).  Make sure the token has the `public_repo` or `repo` scope.

    -   An example `.env` file (`.env_example`) is provided for reference.

3.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

## 🚀 Usage

1.  **Run the program:**

    ```bash
    go run cmd/main.go
    ```

2.  **Observe the output:**

    The program will list the people you follow who aren't following you back and attempt to unfollow them.  It will display success or failure messages for each unfollow attempt.

## ⚙️ Configuration

| Variable       | Description                                                                       |
| -------------- | --------------------------------------------------------------------------------- |
| `MY_GITHUB_TOKEN` | Your GitHub Personal Access Token.  **Required**.  Keep this secure!               |

## 📜 Contributing

Contributions are welcome! Here's how you can contribute:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Make your changes.
4.  Submit a pull request.

Please ensure your code follows the project's coding standards and includes appropriate tests.

## 📄 License

This project has no license.

[![Built with Dokugen](https://img.shields.io/badge/Built%20with-Dokugen-brightgreen)](https://github.com/samueltuoyo15/Dokugen)
