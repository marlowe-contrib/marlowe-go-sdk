# PostContractSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractSourceId** | **string** | The hex-encoded identifier of a Marlowe contract source | 
**IntermediateIds** | **map[string]string** |  | 

## Methods

### NewPostContractSourceResponse

`func NewPostContractSourceResponse(contractSourceId string, intermediateIds map[string]string, ) *PostContractSourceResponse`

NewPostContractSourceResponse instantiates a new PostContractSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostContractSourceResponseWithDefaults

`func NewPostContractSourceResponseWithDefaults() *PostContractSourceResponse`

NewPostContractSourceResponseWithDefaults instantiates a new PostContractSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractSourceId

`func (o *PostContractSourceResponse) GetContractSourceId() string`

GetContractSourceId returns the ContractSourceId field if non-nil, zero value otherwise.

### GetContractSourceIdOk

`func (o *PostContractSourceResponse) GetContractSourceIdOk() (*string, bool)`

GetContractSourceIdOk returns a tuple with the ContractSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractSourceId

`func (o *PostContractSourceResponse) SetContractSourceId(v string)`

SetContractSourceId sets ContractSourceId field to given value.


### GetIntermediateIds

`func (o *PostContractSourceResponse) GetIntermediateIds() map[string]string`

GetIntermediateIds returns the IntermediateIds field if non-nil, zero value otherwise.

### GetIntermediateIdsOk

`func (o *PostContractSourceResponse) GetIntermediateIdsOk() (*map[string]string, bool)`

GetIntermediateIdsOk returns a tuple with the IntermediateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateIds

`func (o *PostContractSourceResponse) SetIntermediateIds(v map[string]string)`

SetIntermediateIds sets IntermediateIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


