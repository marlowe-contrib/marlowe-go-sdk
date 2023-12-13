# Payout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | [**Assets**](Assets.md) |  | 
**PayoutId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**Role** | **string** |  | 

## Methods

### NewPayout

`func NewPayout(assets Assets, payoutId string, role string, ) *Payout`

NewPayout instantiates a new Payout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutWithDefaults

`func NewPayoutWithDefaults() *Payout`

NewPayoutWithDefaults instantiates a new Payout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *Payout) GetAssets() Assets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Payout) GetAssetsOk() (*Assets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Payout) SetAssets(v Assets)`

SetAssets sets Assets field to given value.


### GetPayoutId

`func (o *Payout) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *Payout) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *Payout) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetRole

`func (o *Payout) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Payout) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Payout) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


