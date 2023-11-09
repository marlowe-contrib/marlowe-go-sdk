package main

import (
	"context"
	"fmt"
	openapiclient "marlowe/gen"
	"os"
)

func main() {
	configuration := openapiclient.NewConfiguration()
	configuration.Servers[0].URL = "https://marlowe-runtime-preprod-web.demo.scdev.aws.iohkdev.io"
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ContractsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContractsGet`: ListObjectContractHeader
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsGet`: %v\n", resp)
}
