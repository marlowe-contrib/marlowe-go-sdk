# GetContractResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**GetContractResponseLinks**](GetContractResponseLinks.md) |  | 
**Resource** | [**ContractState**](ContractState.md) |  | 

## Methods

### NewGetContractResponse

`func NewGetContractResponse(links GetContractResponseLinks, resource ContractState, ) *GetContractResponse`

NewGetContractResponse instantiates a new GetContractResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractResponseWithDefaults

`func NewGetContractResponseWithDefaults() *GetContractResponse`

NewGetContractResponseWithDefaults instantiates a new GetContractResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetContractResponse) GetLinks() GetContractResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetContractResponse) GetLinksOk() (*GetContractResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetContractResponse) SetLinks(v GetContractResponseLinks)`

SetLinks sets Links field to given value.


### GetResource

`func (o *GetContractResponse) GetResource() ContractState`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetContractResponse) GetResourceOk() (*ContractState, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetContractResponse) SetResource(v ContractState)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


