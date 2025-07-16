# **GhostBuster** 👻🧹

GhostBuster is my personal GitHub following cleaner! This sleek Go application intelligently identifies and unfollows users who aren't following me back. Keep my GitHub network tidy and reciprocal with ease ✨.

## Installation 🚀

Getting GhostBuster up and running on your local machine is straightforward. Follow these steps:

### Prerequisites

Before you begin, ensure you have [Go](https://go.dev/doc/install) (version 1.23.5 or newer, as specified in `go.mod`) installed on your system.

### Clone the Repository

First, clone the project repository to your local machine:

```bash
git clone https://github.com/samueltuoyo15/Github-Unfollow.git
cd Github-Unfollow
```

### Environment Setup

This application requires a GitHub Personal Access Token to interact with the GitHub API.

✨ Create a new file named `.env` in the root directory of the project:

```bash
touch .env
```

📝 Open the `.env` file and add your GitHub Personal Access Token:

```env
MY_GITHUB_TOKEN=YOUR_GITHUB_PERSONAL_ACCESS_TOKEN
```

*   **Important**: Replace `YOUR_GITHUB_PERSONAL_ACCESS_TOKEN` with your actual token. Ensure this token has `user:follow` scope to allow unfollowing. For security, never commit your `.env` file to version control!

### Install Dependencies

Navigate to the project root and fetch the necessary Go modules:

```bash
go mod tidy
```

### Run the Application

Once dependencies are installed, you can run the application from the `cmd` directory:

```bash
go run cmd/main.go
```

## Usage 🛠️

Once installed, running GhostBuster is simple. The application will automatically connect to the GitHub API using your provided token, identify non-followers, and proceed to unfollow them.

1.  **Execute the application:**
    Open your terminal, navigate to the `Github-Unfollow` directory, and run:
    ```bash
    go run cmd/main.go
    ```

2.  **Observe the output:**
    The application will first display your GitHub username, then proceed to fetch your followers and the users you are following.
    
    If there are no users to unfollow (everyone you follow also follows you back), you will see a message like:
    ```
    No one to unfollow. am all good!
    ```

    If there are users who are not following you back, the application will list them and then begin the unfollowing process:
    ```
    These are the people that aren't following me back:
    - user1
    Successfully Unfollowed user1
    - user2
    Successfully Unfollowed user2
    Failed to unfollow user3: API error message here
    ...
    Successfully unfollowed X users.
    ```
    
    *   **Note on rate limiting**: The application includes a small delay (`500ms`) between unfollow actions to respect GitHub API rate limits. This helps prevent potential issues with too many rapid requests.

## Key Features 🌟

GhostBuster is designed to be effective and user-friendly, offering the following capabilities:

*   **Automated Unfollowing**: Automatically scans your GitHub following list and unfollows accounts that do not follow you back, saving you manual effort.
*   **GitHub API Integration**: Built upon `github.com/google/go-github`, ensuring reliable and up-to-date interaction with the GitHub API.
*   **Secure Credential Management**: Utilizes `godotenv` to securely load your GitHub Personal Access Token from an environment file, keeping sensitive information out of your codebase.
*   **Rate Limit Compliance**: Implements strategic delays between API calls to gracefully handle GitHub's rate limits, preventing service interruptions.
*   **Informative Output**: Provides clear, real-time feedback on which accounts are being unfollowed, along with a final summary of the operation.

## Technologies Used 💻

This project is crafted with modern Go practices and relies on a few well-regarded libraries:

| Technology  | Description                                        | Link                                      |
| :---------- | :------------------------------------------------- | :---------------------------------------- |
| **Go**      | The primary programming language used.             | [go.dev](https://go.dev/)                 |
| **Go-GitHub** | Official Go client for the GitHub API.           | [GitHub](https://github.com/google/go-github) |
| **OAuth2**  | Go packages for OAuth 2.0.                       | [GitHub](https://golang.org/x/oauth2)     |
| **Godotenv**| Loads environment variables from a `.env` file.  | [GitHub](https://github.com/joho/godotenv) |

## License 📄

No specific license file was found within this repository. If you plan to use or distribute this project, please consider adding an explicit license.

## About the Author 👋

Hi there! I'm Samuel Tuoyo, the creator of GhostBuster. I'm passionate about building practical tools with Go that simplify everyday development tasks.

Feel free to connect with me:

*   **GitHub**: [@samueltuoyo15](https://github.com/samueltuoyo15)
*   **LinkedIn**: [Your LinkedIn Profile](https://www.linkedin.com/in/samuel-tuoyo-8568b62b6)
*   **Twitter**: [@your_twitter](https://x.com/TuoyoS26091)

---

<p align="center">
  <a href="https://go.dev/" target="_blank">
    <img src="https://img.shields.io/badge/Go-1.23.5-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
  </a>
  <a href="https://github.com/samueltuoyo15/Github-Unfollow" target="_blank">
    <img src="https://img.shields.io/github/stars/samueltuoyo15/Github-Unfollow?style=for-the-badge&color=FFE066&label=Stars&logo=github" alt="GitHub Stars">
  </a>
  <a href="https://github.com/samueltuoyo15/Github-Unfollow" target="_blank">
    <img src="https://img.shields.io/github/forks/samueltuoyo15/Github-Unfollow?style=for-the-badge&color=64FFDA&label=Forks&logo=github" alt="GitHub Forks">
  </a>
  <img src="https://img.shields.io/badge/License-Unlicensed-blue.svg?style=for-the-badge" alt="License">
</p>

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)