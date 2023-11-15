# PartyObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleToken** | **string** |  | 
**Address** | **string** | A cardano address, in Bech32 format | 
**Ref** | **string** | An arbitrary text identifier for an object in a Marlowe object bundle. | 

## Methods

### NewPartyObject

`func NewPartyObject(roleToken string, address string, ref string, ) *PartyObject`

NewPartyObject instantiates a new PartyObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyObjectWithDefaults

`func NewPartyObjectWithDefaults() *PartyObject`

NewPartyObjectWithDefaults instantiates a new PartyObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleToken

`func (o *PartyObject) GetRoleToken() string`

GetRoleToken returns the RoleToken field if non-nil, zero value otherwise.

### GetRoleTokenOk

`func (o *PartyObject) GetRoleTokenOk() (*string, bool)`

GetRoleTokenOk returns a tuple with the RoleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleToken

`func (o *PartyObject) SetRoleToken(v string)`

SetRoleToken sets RoleToken field to given value.


### GetAddress

`func (o *PartyObject) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PartyObject) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PartyObject) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetRef

`func (o *PartyObject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PartyObject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PartyObject) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


