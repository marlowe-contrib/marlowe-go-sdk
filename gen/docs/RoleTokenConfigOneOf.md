# RoleTokenConfigOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | A cardano address, in Bech32 format | 
**Metadata** | [**TokenMetadata**](TokenMetadata.md) |  | 

## Methods

### NewRoleTokenConfigOneOf

`func NewRoleTokenConfigOneOf(address string, metadata TokenMetadata, ) *RoleTokenConfigOneOf`

NewRoleTokenConfigOneOf instantiates a new RoleTokenConfigOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleTokenConfigOneOfWithDefaults

`func NewRoleTokenConfigOneOfWithDefaults() *RoleTokenConfigOneOf`

NewRoleTokenConfigOneOfWithDefaults instantiates a new RoleTokenConfigOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RoleTokenConfigOneOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RoleTokenConfigOneOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RoleTokenConfigOneOf) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMetadata

`func (o *RoleTokenConfigOneOf) GetMetadata() TokenMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoleTokenConfigOneOf) GetMetadataOk() (*TokenMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoleTokenConfigOneOf) SetMetadata(v TokenMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


