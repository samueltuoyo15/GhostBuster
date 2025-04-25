package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/google/go-github/v50/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const logFile = "unfollow_log.txt"

func appendToLog(message string) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}

func gitCommitLog() {
	cmds := [][]string{
		{"git", "add", logFile},
		{"git", "commit", "-m", "Update unfollow log"},
		{"git", "push"},
	}

	for _, cmdArgs := range cmds {
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			appendToLog(fmt.Sprintf("Git command failed: %v", err))
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("MY_GITHUB_TOKEN")
	if token == "" {
		log.Fatal("No Token Found in .env")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Fatalf("Error getting user: %v", err)
	}

	username := *user.Login
	appendToLog(fmt.Sprintf("Starting unfollow check for user: %s at %s", username, time.Now().Format(time.RFC1123)))

	var allFollowers []*github.User
	followerOpt := &github.ListOptions{PerPage: 100}
	for {
		followers, resp, err := client.Users.ListFollowers(ctx, username, followerOpt)
		if err != nil {
			log.Fatalf("Error getting followers: %v", err)
		}
		allFollowers = append(allFollowers, followers...)
		if resp.NextPage == 0 {
			break
		}
		followerOpt.Page = resp.NextPage
	}

	var allFollowing []*github.User
	followingOpt := &github.ListOptions{PerPage: 100}
	for {
		following, resp, err := client.Users.ListFollowing(ctx, username, followingOpt)
		if err != nil {
			log.Fatalf("Error getting following: %v", err)
		}
		allFollowing = append(allFollowing, following...)
		if resp.NextPage == 0 {
			break
		}
		followingOpt.Page = resp.NextPage
	}

	followerMap := make(map[string]bool)
	for _, follower := range allFollowers {
		followerMap[*follower.Login] = true
	}

	unfollowCount := 0
	for _, followee := range allFollowing {
		if !followerMap[*followee.Login] {
			appendToLog(fmt.Sprintf("Not following back: %s", *followee.Login))

			_, err := client.Users.Unfollow(ctx, *followee.Login)
			if err != nil {
				appendToLog(fmt.Sprintf("Failed to unfollow %s: %v", *followee.Login, err))
			} else {
				appendToLog(fmt.Sprintf("Successfully unfollowed %s", *followee.Login))
				unfollowCount++
			}

			time.Sleep(500 * time.Millisecond)
		}
	}

	appendToLog(fmt.Sprintf("Finished unfollowing %d users", unfollowCount))

	// Auto commit and push the log
	gitCommitLog()
}
