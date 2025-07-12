package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cars/sfss-go-client/sfss"
)

func main() {
	cfg := sfss.NewConfiguration()
	cfg.Host = "https://api.example.com"
	cfg.Scheme = "https"

	client := sfss.NewAPIClient(cfg)
	client.DefaultAPI.GetIDRedfishV1SFSSApp()
	resp, _, err := client.SFSSAppApi.GetSosReports(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SFSSAppApi.GetSosReports`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response from `SFSSAppApi.GetSosReports`: %v\n", resp)
}
