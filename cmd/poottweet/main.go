package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type BearerTokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func main() {
	// Get configuration.
	key := flag.String("consumer-key", "", "Twitter API consumer key")
	secret := flag.String("consumer-secret", "", "Twitter API consumer secret")
	flag.Parse()

	// Obtain a Bearer token.
	creds := base64.StdEncoding.EncodeToString([]byte(*key + ":" + *secret))

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.twitter.com/oauth2/token",
		strings.NewReader("grant_type=client_credentials"),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Authorization", "Basic "+creds)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var token BearerTokenResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		panic(err)
	}

	if token.TokenType != "bearer" {
		panic("invalid bearer token response: " + string(body))
	}

	fmt.Printf("%#v\n", token)

	// Drink from the firehose.
	fhreq, err := http.NewRequest(
		http.MethodGet,
		"https://api.twitter.com/labs/1/tweets/stream/sample",
		nil,
	)
	if err != nil {
		panic(err)
	}
	fhreq.Header.Add("Authorization", "Bearer "+token.AccessToken)

	firehose, err := http.DefaultClient.Do(fhreq)
	if err != nil {
		panic(err)
	}

	i := 0
	last := time.Now()
	scanner := bufio.NewScanner(firehose.Body)
	for scanner.Scan() {
		i++
		if i%100 == 0 {
			now := time.Now()
			delta := now.Sub(last)
			fmt.Println(i, delta.Seconds())
			last = now
		}
	}
}
