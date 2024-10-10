package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

func main() {
	consumerKey := "access key"
	consumerSecret := "secret access key"
	accessToken := "access token"
	accessTokenSecret := "secret access token"

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Fatal("Twitter API credentials are not set in environment variables.")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	tweetURL := "https://api.twitter.com/1.1/statuses/update.json"
	data := url.Values{}
	data.Set("status", " preetham from API with Go!")

	resp, err := httpClient.PostForm(tweetURL, data)
	if err != nil {
		log.Fatalf("Failed to post tweet: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Tweet posted successfully!")
	} else {
		fmt.Printf("Failed to post tweet, status code: %d\n", resp.StatusCode)
	}
}
