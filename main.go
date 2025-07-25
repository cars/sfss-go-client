package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	sfssapp "github.com/cars/sfss-go-client/sfssapp"
)

func main() {
	username := os.Getenv("SFSS_USERNAME")
	if username == "" {
		log.Fatalf("SFSS_USERNAME environment variable is not set")
		return
	}

	password := os.Getenv("SFSS_PASSWORD")
	if password == "" {
		log.Fatalf("SFSS_PASSWORD environment variable is not set")
		return
	}
	// userpw := fmt.Sprintf("%s:%s", username, password)
	// authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(userpw))

	host := os.Getenv("SFSS_HOST")
	if host == "" {
		log.Fatalf("SFSS_HOST environment variable is not set")
		return
	}
	cfg := sfssapp.NewConfiguration()
	cfg.Host = host
	cfg.Scheme = "https"
	xport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	tcli := &http.Client{Transport: xport}
	cfg.HTTPClient = tcli
	//cfg.DefaultHeader["Authorization"] = authHeader
	debug := os.Getenv("SFSS_DEBUG")
	if strings.ToLower(debug) == "true" {
		cfg.Debug = true
	} else {
		cfg.Debug = false
	}

	client := sfssapp.NewAPIClient(cfg)
	if client == nil {
		fmt.Println("Failed to create API client")
		return
	}
	//ctx := context.Background()
	ctx := context.WithValue(context.Background(), sfssapp.ContextBasicAuth, sfssapp.BasicAuth{
		UserName: username,
		Password: password,
	})

	req := client.DefaultAPI.GetRedfishV1SFSSApp(ctx)
	rf_resp, http_resp, err := client.DefaultAPI.GetRedfishV1SFSSAppExecute(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SFSSAppApi.GetRedfishV1SFSSApp`: %v\n", err)
		fmt.Printf("HTTP Response: %v\n", http_resp)
		return
	}
	if rf_resp == nil {
		fmt.Println("Failed to get SFSS application details")
		return
	} else {
		fmt.Printf("SFSS Application Details: %v\n", req)
	}

}
