package main

import (
  "context"
  "fmt"
  "log"
  "os"
  "time"
  "github.com/google/go-github/v50/github"
  "golang.org/x/oauth2"
  "github.com/joho/godotenv"
)

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
      unfollowCount++
    }
  }

  if unfollowCount == 0 {
    fmt.Println("No one to unfollow. am all good!")
    return
  }


  fmt.Println("These are the people that aren't following me back:")
  for _, followee := range allFollowing {
    if !followerMap[*followee.Login] {
      fmt.Printf("- %s\n", *followee.Login)

      _, err := client.Users.Unfollow(ctx, *followee.Login)
      if err != nil {
        fmt.Printf("Failed to unfollow %s: %v\n", *followee.Login, err)
      } else {
        fmt.Printf("Successfully Unfollowed %s\n", *followee.Login)
        unfollowCount++
      }

      time.Sleep(500 * time.Millisecond) 
    }
  }

  fmt.Printf("Successfully unfollowed %d users.\n", unfollowCount)
}