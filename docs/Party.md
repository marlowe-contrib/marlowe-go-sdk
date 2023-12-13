# Party

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleToken** | **string** |  | 
**Address** | **string** | A cardano address, in Bech32 format | 

## Methods

### NewParty

`func NewParty(roleToken string, address string, ) *Party`

NewParty instantiates a new Party object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyWithDefaults

`func NewPartyWithDefaults() *Party`

NewPartyWithDefaults instantiates a new Party object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleToken

`func (o *Party) GetRoleToken() string`

GetRoleToken returns the RoleToken field if non-nil, zero value otherwise.

### GetRoleTokenOk

`func (o *Party) GetRoleTokenOk() (*string, bool)`

GetRoleTokenOk returns a tuple with the RoleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleToken

`func (o *Party) SetRoleToken(v string)`

SetRoleToken sets RoleToken field to given value.


### GetAddress

`func (o *Party) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Party) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Party) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


