package main

import (
	"context"
	"fmt"
	"os"

	sfssapp "github.com/cars/sfss-go-client/sfssapp"
)

func main() {
	cfg := sfssapp.NewConfiguration()
	cfg.Host = "https://api.example.com"
	cfg.Scheme = "https"

	client := sfssapp.NewAPIClient(cfg)
	if client == nil {
		fmt.Println("Failed to create API client")
		return
	}
	ctx := context.Background()
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
	resp2 := client.DefaultAPI.GetRedfishV1SFSSAppSosReports(context.Background())
	// Handle the error if it occurs
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error when calling `SFSSAppApi.GetSosReports`: %v\n", err)
	// 	os.Exit(1)
	// }
	fmt.Printf("Response from `SFSSAppApi.GetSosReports`: %v\n", resp2)
}
