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
  
  token := os.Getenv("GITHUB_TOKEN")
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
  
   fmt.Printf("These are the people rhat aren't following u back:")
  for _, followee := range following{
    if !followerMap[*followee.Login] {
      fmt.Printf("- %s\n", *followee.Login)
    }
  }
}