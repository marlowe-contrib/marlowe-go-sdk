# RoleTokenConfigOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**TokenMetadata**](TokenMetadata.md) |  | [optional] 
**Script** | **string** | The type of script receiving the role token. | 

## Methods

### NewRoleTokenConfigOneOf1

`func NewRoleTokenConfigOneOf1(script string, ) *RoleTokenConfigOneOf1`

NewRoleTokenConfigOneOf1 instantiates a new RoleTokenConfigOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleTokenConfigOneOf1WithDefaults

`func NewRoleTokenConfigOneOf1WithDefaults() *RoleTokenConfigOneOf1`

NewRoleTokenConfigOneOf1WithDefaults instantiates a new RoleTokenConfigOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RoleTokenConfigOneOf1) GetMetadata() TokenMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoleTokenConfigOneOf1) GetMetadataOk() (*TokenMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoleTokenConfigOneOf1) SetMetadata(v TokenMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RoleTokenConfigOneOf1) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScript

`func (o *RoleTokenConfigOneOf1) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *RoleTokenConfigOneOf1) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *RoleTokenConfigOneOf1) SetScript(v string)`

SetScript sets Script field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


