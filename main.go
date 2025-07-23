package main

import (
	"context"
	"fmt"

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
	req := client.DefaultAPI.GetRedfishV1SFSSApp(context.Background())
	resp1 := client.DefaultAPI.GetRedfishV1SFSSAppExecute(context.Background())
	if resp1 == nil {
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
	fmt.Printf("Response from `SFSSAppApi.GetSosReports`: %v\n", resp)
}
