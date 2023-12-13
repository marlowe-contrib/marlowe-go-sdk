# MetadataAndScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**TokenMetadata**](TokenMetadata.md) |  | [optional] 
**Script** | **string** | The type of script receiving the role token. | 

## Methods

### NewMetadataAndScript

`func NewMetadataAndScript(script string, ) *MetadataAndScript`

NewMetadataAndScript instantiates a new MetadataAndScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataAndScriptWithDefaults

`func NewMetadataAndScriptWithDefaults() *MetadataAndScript`

NewMetadataAndScriptWithDefaults instantiates a new MetadataAndScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *MetadataAndScript) GetMetadata() TokenMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetadataAndScript) GetMetadataOk() (*TokenMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetadataAndScript) SetMetadata(v TokenMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MetadataAndScript) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScript

`func (o *MetadataAndScript) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *MetadataAndScript) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *MetadataAndScript) SetScript(v string)`

SetScript sets Script field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


