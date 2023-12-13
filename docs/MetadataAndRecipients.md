# MetadataAndRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**TokenMetadata**](TokenMetadata.md) |  | [optional] 
**Recipients** | **map[string]int64** |  | 

## Methods

### NewMetadataAndRecipients

`func NewMetadataAndRecipients(recipients map[string]int64, ) *MetadataAndRecipients`

NewMetadataAndRecipients instantiates a new MetadataAndRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataAndRecipientsWithDefaults

`func NewMetadataAndRecipientsWithDefaults() *MetadataAndRecipients`

NewMetadataAndRecipientsWithDefaults instantiates a new MetadataAndRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *MetadataAndRecipients) GetMetadata() TokenMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetadataAndRecipients) GetMetadataOk() (*TokenMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetadataAndRecipients) SetMetadata(v TokenMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MetadataAndRecipients) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRecipients

`func (o *MetadataAndRecipients) GetRecipients() map[string]int64`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *MetadataAndRecipients) GetRecipientsOk() (*map[string]int64, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *MetadataAndRecipients) SetRecipients(v map[string]int64)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


