# PayoutState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | [**Assets**](Assets.md) |  | 
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**PayoutId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**PayoutValidatorAddress** | **string** | A cardano address, in Bech32 format | 
**Role** | [**AssetId**](AssetId.md) |  | 
**Status** | [**PayoutStatus**](PayoutStatus.md) |  | 
**WithdrawalId** | Pointer to **string** | The hex-encoded identifier of a Cardano transaction | [optional] 

## Methods

### NewPayoutState

`func NewPayoutState(assets Assets, contractId string, payoutId string, payoutValidatorAddress string, role AssetId, status PayoutStatus, ) *PayoutState`

NewPayoutState instantiates a new PayoutState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutStateWithDefaults

`func NewPayoutStateWithDefaults() *PayoutState`

NewPayoutStateWithDefaults instantiates a new PayoutState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *PayoutState) GetAssets() Assets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *PayoutState) GetAssetsOk() (*Assets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *PayoutState) SetAssets(v Assets)`

SetAssets sets Assets field to given value.


### GetContractId

`func (o *PayoutState) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *PayoutState) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *PayoutState) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetPayoutId

`func (o *PayoutState) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PayoutState) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PayoutState) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetPayoutValidatorAddress

`func (o *PayoutState) GetPayoutValidatorAddress() string`

GetPayoutValidatorAddress returns the PayoutValidatorAddress field if non-nil, zero value otherwise.

### GetPayoutValidatorAddressOk

`func (o *PayoutState) GetPayoutValidatorAddressOk() (*string, bool)`

GetPayoutValidatorAddressOk returns a tuple with the PayoutValidatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutValidatorAddress

`func (o *PayoutState) SetPayoutValidatorAddress(v string)`

SetPayoutValidatorAddress sets PayoutValidatorAddress field to given value.


### GetRole

`func (o *PayoutState) GetRole() AssetId`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PayoutState) GetRoleOk() (*AssetId, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PayoutState) SetRole(v AssetId)`

SetRole sets Role field to given value.


### GetStatus

`func (o *PayoutState) GetStatus() PayoutStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutState) GetStatusOk() (*PayoutStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutState) SetStatus(v PayoutStatus)`

SetStatus sets Status field to given value.


### GetWithdrawalId

`func (o *PayoutState) GetWithdrawalId() string`

GetWithdrawalId returns the WithdrawalId field if non-nil, zero value otherwise.

### GetWithdrawalIdOk

`func (o *PayoutState) GetWithdrawalIdOk() (*string, bool)`

GetWithdrawalIdOk returns a tuple with the WithdrawalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalId

`func (o *PayoutState) SetWithdrawalId(v string)`

SetWithdrawalId sets WithdrawalId field to given value.

### HasWithdrawalId

`func (o *PayoutState) HasWithdrawalId() bool`

HasWithdrawalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


