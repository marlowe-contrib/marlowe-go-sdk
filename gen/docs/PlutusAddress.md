# PlutusAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressCredential** | [**PlutusCredential**](PlutusCredential.md) |  | 
**AddressStakingCredential** | Pointer to [**PlutusStakingCredential**](PlutusStakingCredential.md) |  | [optional] 

## Methods

### NewPlutusAddress

`func NewPlutusAddress(addressCredential PlutusCredential, ) *PlutusAddress`

NewPlutusAddress instantiates a new PlutusAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlutusAddressWithDefaults

`func NewPlutusAddressWithDefaults() *PlutusAddress`

NewPlutusAddressWithDefaults instantiates a new PlutusAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressCredential

`func (o *PlutusAddress) GetAddressCredential() PlutusCredential`

GetAddressCredential returns the AddressCredential field if non-nil, zero value otherwise.

### GetAddressCredentialOk

`func (o *PlutusAddress) GetAddressCredentialOk() (*PlutusCredential, bool)`

GetAddressCredentialOk returns a tuple with the AddressCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCredential

`func (o *PlutusAddress) SetAddressCredential(v PlutusCredential)`

SetAddressCredential sets AddressCredential field to given value.


### GetAddressStakingCredential

`func (o *PlutusAddress) GetAddressStakingCredential() PlutusStakingCredential`

GetAddressStakingCredential returns the AddressStakingCredential field if non-nil, zero value otherwise.

### GetAddressStakingCredentialOk

`func (o *PlutusAddress) GetAddressStakingCredentialOk() (*PlutusStakingCredential, bool)`

GetAddressStakingCredentialOk returns a tuple with the AddressStakingCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStakingCredential

`func (o *PlutusAddress) SetAddressStakingCredential(v PlutusStakingCredential)`

SetAddressStakingCredential sets AddressStakingCredential field to given value.

### HasAddressStakingCredential

`func (o *PlutusAddress) HasAddressStakingCredential() bool`

HasAddressStakingCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


