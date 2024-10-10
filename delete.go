package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
)

func main() {
	consumerKey := os.Getenv("TWITTER_API_KEY")
	consumerSecret := os.Getenv("TWITTER_API_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Fatal("Twitter API credentials are not set in environment variables.")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Replace with your tweet ID to delete
	tweetID := "your-tweet-id"
	deleteURL := fmt.Sprintf("https://api.twitter.com/1.1/statuses/destroy/%s.json", tweetID)

	req, err := http.NewRequest("POST", deleteURL, nil)
	if err != nil {
		log.Fatalf("Failed to create delete request: %v", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to delete tweet: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Tweet deleted successfully!")
	} else {
		fmt.Printf("Failed to delete tweet, status code: %d\n", resp.StatusCode)
	}
}
