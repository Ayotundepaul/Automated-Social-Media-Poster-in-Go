package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func postToTwitter(apiKey, apiSecret, accessToken, accessSecret, message string) {
	anaconda.SetConsumerKey(apiKey)
	anaconda.SetConsumerSecret(apiSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	_, err := api.PostTweet(message, nil)
	if err != nil {
		log.Fatalf("Failed to post tweet: %v", err)
	}
}

func scheduleTweet(delay time.Duration, apiKey, apiSecret, accessToken, accessSecret, message string) {
	fmt.Printf("Tweet scheduled for %v seconds later.\n", delay.Seconds())
	time.Sleep(delay)
	postToTwitter(apiKey, apiSecret, accessToken, accessSecret, message)
}

func main() {
	// Replace these with your credentials.
	apiKey := "YOUR_API_KEY"
	apiSecret := "YOUR_API_SECRET"
	accessToken := "YOUR_ACCESS_TOKEN"
	accessSecret := "YOUR_ACCESS_SECRET"

	var delaySeconds time.Duration
	var message string

	fmt.Println("Enter delay for the tweet (in seconds):")
	fmt.Scan(&delaySeconds)

	fmt.Println("Enter your message:")
	fmt.Scan(&message)

	scheduleTweet(time.Second*delaySeconds, apiKey, apiSecret, accessToken, accessSecret, message)
}
