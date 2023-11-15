# CreateContractResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**CreateContractResponseLinks**](CreateContractResponseLinks.md) |  | 
**Resource** | [**CreateTxEnvelope**](CreateTxEnvelope.md) |  | 

## Methods

### NewCreateContractResponse

`func NewCreateContractResponse(links CreateContractResponseLinks, resource CreateTxEnvelope, ) *CreateContractResponse`

NewCreateContractResponse instantiates a new CreateContractResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContractResponseWithDefaults

`func NewCreateContractResponseWithDefaults() *CreateContractResponse`

NewCreateContractResponseWithDefaults instantiates a new CreateContractResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CreateContractResponse) GetLinks() CreateContractResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateContractResponse) GetLinksOk() (*CreateContractResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateContractResponse) SetLinks(v CreateContractResponseLinks)`

SetLinks sets Links field to given value.


### GetResource

`func (o *CreateContractResponse) GetResource() CreateTxEnvelope`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CreateContractResponse) GetResourceOk() (*CreateTxEnvelope, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CreateContractResponse) SetResource(v CreateTxEnvelope)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


