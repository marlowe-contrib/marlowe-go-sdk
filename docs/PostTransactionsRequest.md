# PostTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**[]Input**](Input.md) |  | 
**InvalidBefore** | Pointer to **string** |  | [optional] 
**InvalidHereafter** | Pointer to **string** |  | [optional] 
**Metadata** | [**map[string]Metadata**](Metadata.md) |  | 
**Tags** | [**map[string]Metadata**](Metadata.md) |  | 
**Version** | [**MarloweVersion**](MarloweVersion.md) |  | 

## Methods

### NewPostTransactionsRequest

`func NewPostTransactionsRequest(inputs []Input, metadata map[string]Metadata, tags map[string]Metadata, version MarloweVersion, ) *PostTransactionsRequest`

NewPostTransactionsRequest instantiates a new PostTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTransactionsRequestWithDefaults

`func NewPostTransactionsRequestWithDefaults() *PostTransactionsRequest`

NewPostTransactionsRequestWithDefaults instantiates a new PostTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *PostTransactionsRequest) GetInputs() []Input`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *PostTransactionsRequest) GetInputsOk() (*[]Input, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *PostTransactionsRequest) SetInputs(v []Input)`

SetInputs sets Inputs field to given value.


### GetInvalidBefore

`func (o *PostTransactionsRequest) GetInvalidBefore() string`

GetInvalidBefore returns the InvalidBefore field if non-nil, zero value otherwise.

### GetInvalidBeforeOk

`func (o *PostTransactionsRequest) GetInvalidBeforeOk() (*string, bool)`

GetInvalidBeforeOk returns a tuple with the InvalidBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidBefore

`func (o *PostTransactionsRequest) SetInvalidBefore(v string)`

SetInvalidBefore sets InvalidBefore field to given value.

### HasInvalidBefore

`func (o *PostTransactionsRequest) HasInvalidBefore() bool`

HasInvalidBefore returns a boolean if a field has been set.

### GetInvalidHereafter

`func (o *PostTransactionsRequest) GetInvalidHereafter() string`

GetInvalidHereafter returns the InvalidHereafter field if non-nil, zero value otherwise.

### GetInvalidHereafterOk

`func (o *PostTransactionsRequest) GetInvalidHereafterOk() (*string, bool)`

GetInvalidHereafterOk returns a tuple with the InvalidHereafter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidHereafter

`func (o *PostTransactionsRequest) SetInvalidHereafter(v string)`

SetInvalidHereafter sets InvalidHereafter field to given value.

### HasInvalidHereafter

`func (o *PostTransactionsRequest) HasInvalidHereafter() bool`

HasInvalidHereafter returns a boolean if a field has been set.

### GetMetadata

`func (o *PostTransactionsRequest) GetMetadata() map[string]Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostTransactionsRequest) GetMetadataOk() (*map[string]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostTransactionsRequest) SetMetadata(v map[string]Metadata)`

SetMetadata sets Metadata field to given value.


### GetTags

`func (o *PostTransactionsRequest) GetTags() map[string]Metadata`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostTransactionsRequest) GetTagsOk() (*map[string]Metadata, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostTransactionsRequest) SetTags(v map[string]Metadata)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *PostTransactionsRequest) GetVersion() MarloweVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostTransactionsRequest) GetVersionOk() (*MarloweVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostTransactionsRequest) SetVersion(v MarloweVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


