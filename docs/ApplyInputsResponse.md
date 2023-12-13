# ApplyInputsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**ApplyInputsResponseLinks**](ApplyInputsResponseLinks.md) |  | 
**Resource** | [**ApplyInputsTxEnvelope**](ApplyInputsTxEnvelope.md) |  | 

## Methods

### NewApplyInputsResponse

`func NewApplyInputsResponse(links ApplyInputsResponseLinks, resource ApplyInputsTxEnvelope, ) *ApplyInputsResponse`

NewApplyInputsResponse instantiates a new ApplyInputsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyInputsResponseWithDefaults

`func NewApplyInputsResponseWithDefaults() *ApplyInputsResponse`

NewApplyInputsResponseWithDefaults instantiates a new ApplyInputsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplyInputsResponse) GetLinks() ApplyInputsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplyInputsResponse) GetLinksOk() (*ApplyInputsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplyInputsResponse) SetLinks(v ApplyInputsResponseLinks)`

SetLinks sets Links field to given value.


### GetResource

`func (o *ApplyInputsResponse) GetResource() ApplyInputsTxEnvelope`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ApplyInputsResponse) GetResourceOk() (*ApplyInputsTxEnvelope, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ApplyInputsResponse) SetResource(v ApplyInputsTxEnvelope)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


