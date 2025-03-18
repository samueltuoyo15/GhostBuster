package main

import (
  "context"
  "fmt"
  "log"
  "os"
  "github.com/google/go-github/v50/github"
  "golang.org/x/oauth2"
  "github.com/joho/godotenv"
)

func main(){
  err := godotenv.Load()
  if err != nil{
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
    log.Fatalf("Error getting user:", err)
  }
  
  username := *user.Login
  
  followers, _, err := client.Users.ListFollowers(ctx, username, nil)
  if err != nil {
    log.Fatalf("Error getting followers: %v", err)
  }
  
  following, _, err := client.Users.ListFollowing(ctx, username, nil)
  if err != nil {
    log.Fatalf("Error getting following: %v", err)
  }
  
  followerMap := make(map[string]bool)
 
  for _, follower := range followers{
    followerMap[*follower.Login] = true
  }
  
  followeeMap := make(map[string]bool)
  for _, followee := range following{
    followeeMap[*followee.Login] = true
  }
  
  for _, follower := range followers{
    if !followeeMap[*follower.Login] {
      _, err := client.Users.Follow(ctx, *follower.Login)
      if err != nil {
        fmt.Printf("Failed to follow %s: %v\n", *follower.Login, err)
      } else{
        fmt.Printf("Successfully Followed %s\n", *follower.Login)
      }
    }
  }
  
  for _, followee := range following{
    if !followerMap[*followee.Login] {
      _, err := client.Users.Unfollow(ctx, *followee.Login)
      if err != nil {
        fmt.Printf("Failed to unfollow %s: %v\n", *followee.Login, err)
      } else{
        fmt.Printf("Successfully Unfollowed %s\n", *followee.Login)
      }
    }
  }
}