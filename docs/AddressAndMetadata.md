# AddressAndMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | A cardano address, in Bech32 format | 
**Metadata** | Pointer to [**TokenMetadata**](TokenMetadata.md) |  | [optional] 

## Methods

### NewAddressAndMetadata

`func NewAddressAndMetadata(address string, ) *AddressAndMetadata`

NewAddressAndMetadata instantiates a new AddressAndMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressAndMetadataWithDefaults

`func NewAddressAndMetadataWithDefaults() *AddressAndMetadata`

NewAddressAndMetadataWithDefaults instantiates a new AddressAndMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressAndMetadata) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressAndMetadata) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressAndMetadata) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMetadata

`func (o *AddressAndMetadata) GetMetadata() TokenMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddressAndMetadata) GetMetadataOk() (*TokenMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddressAndMetadata) SetMetadata(v TokenMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddressAndMetadata) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


