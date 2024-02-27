# marlowe-go-sdk

This library is meant to make your interaction with Marlowe Contracts much easier through an abstraction of the Marlowe Runtime REST API.

## Usage

Refer to the `docs/` to see how to see how the package is meant to be used.

Here's a sample of the library usage from the `docs/DefaultAPI.md` file:

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/marlowe-contrib/marlowe-go-sdk"
)

func main() {
    contractId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string |
    xChangeAddress := "addr1w94f8ywk4fg672xasahtk4t9k6w3aql943uxz5rt62d4dvq8evxaf" // string |
    xAddress := "xAddress_example" // string |  (optional)
    xCollateralUTxO := "xCollateralUTxO_example" // string |  (optional)
    postTransactionsRequest := *openapiclient.NewPostTransactionsRequest([]openapiclient.Input{openapiclient.Input{ChoiceContinuationInput: openapiclient.NewChoiceContinuationInput("ContinuationHash_example", *openapiclient.NewChoiceId("ChoiceName_example", openapiclient.Party{PartyAddress: openapiclient.NewPartyAddress("addr1w94f8ywk4fg672xasahtk4t9k6w3aql943uxz5rt62d4dvq8evxaf")}), int32(123), openapiclient.Contract{Assert: openapiclient.NewAssert(openapiclient.Observation{And: openapiclient.NewAnd(openapiclient.Observation{And: openapiclient.NewAnd(openapiclient.Observation{And: }, openapiclient.Observation{And: })}, openapiclient.Observation{And: })}, openapiclient.Contract{Assert: openapiclient.NewAssert(openapiclient.Observation{And: }, openapiclient.Contract{Assert: })})})}}, map[string]Metadata{"key": openapiclient.Metadata{ArrayOfMetadata: new([]Metadata)}}, map[string]Metadata{"key": openapiclient.Metadata{ArrayOfMetadata: new([]Metadata)}}, openapiclient.MarloweVersion("v1")) // PostTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApplyInputsToContract(context.Background(), contractId).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostTransactionsRequest(postTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApplyInputsToContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyInputsToContract`: ApplyInputsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApplyInputsToContract`: %v\n", resp)
}

```

## Check OpenAPI specs validity

Run the following command to check if the `openapi.json` is valid.

```sh
npx @openapitools/openapi-generator-cli validate -i openapi.json
```

## Versioning

Refer to the releases page to see which is compatible with each Marlowe Runtime REST API version.

## License

This package is under the Apache-2.0 license.
