# GetTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**GetTransactionResponseLinks**](GetTransactionResponseLinks.md) |  | 
**Resource** | [**Tx**](Tx.md) |  | 

## Methods

### NewGetTransactionResponse

`func NewGetTransactionResponse(links GetTransactionResponseLinks, resource Tx, ) *GetTransactionResponse`

NewGetTransactionResponse instantiates a new GetTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionResponseWithDefaults

`func NewGetTransactionResponseWithDefaults() *GetTransactionResponse`

NewGetTransactionResponseWithDefaults instantiates a new GetTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetTransactionResponse) GetLinks() GetTransactionResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTransactionResponse) GetLinksOk() (*GetTransactionResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTransactionResponse) SetLinks(v GetTransactionResponseLinks)`

SetLinks sets Links field to given value.


### GetResource

`func (o *GetTransactionResponse) GetResource() Tx`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetTransactionResponse) GetResourceOk() (*Tx, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetTransactionResponse) SetResource(v Tx)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


