# RoleTokenConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | A cardano address, in Bech32 format | 
**Metadata** | Pointer to [**TokenMetadata**](TokenMetadata.md) |  | [optional] 
**Script** | **string** | The type of script receiving the role token. | 
**Recipients** | **map[string]int64** |  | 

## Methods

### NewRoleTokenConfig

`func NewRoleTokenConfig(address string, script string, recipients map[string]int64, ) *RoleTokenConfig`

NewRoleTokenConfig instantiates a new RoleTokenConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleTokenConfigWithDefaults

`func NewRoleTokenConfigWithDefaults() *RoleTokenConfig`

NewRoleTokenConfigWithDefaults instantiates a new RoleTokenConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RoleTokenConfig) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RoleTokenConfig) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RoleTokenConfig) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMetadata

`func (o *RoleTokenConfig) GetMetadata() TokenMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoleTokenConfig) GetMetadataOk() (*TokenMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoleTokenConfig) SetMetadata(v TokenMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RoleTokenConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScript

`func (o *RoleTokenConfig) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *RoleTokenConfig) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *RoleTokenConfig) SetScript(v string)`

SetScript sets Script field to given value.


### GetRecipients

`func (o *RoleTokenConfig) GetRecipients() map[string]int64`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *RoleTokenConfig) GetRecipientsOk() (*map[string]int64, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *RoleTokenConfig) SetRecipients(v map[string]int64)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


