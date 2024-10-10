package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

func main() {
	consumerKey := "1506GsOdJeMze8WJycdGiSVq4"
	consumerSecret := "eAGFwSP7xa1JnuuH5Vz7Fjii2ewCehkSO8j3aT1H31iSFezedq"
	accessToken := "1761252971705139200-ZCxlSIOP5bgfkSTTV4VMBJEgEJuc5r"
	accessTokenSecret := "IhYFG8w5Y3XJg6zsUBvuLaCKIzeQT1ScyBnNQPmAhqHf1"

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
