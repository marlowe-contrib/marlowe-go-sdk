# PayoutHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**PayoutId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**Role** | [**AssetId**](AssetId.md) |  | 
**Status** | [**PayoutStatus**](PayoutStatus.md) |  | 
**WithdrawalId** | Pointer to **string** | The hex-encoded identifier of a Cardano transaction | [optional] 

## Methods

### NewPayoutHeader

`func NewPayoutHeader(contractId string, payoutId string, role AssetId, status PayoutStatus, ) *PayoutHeader`

NewPayoutHeader instantiates a new PayoutHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutHeaderWithDefaults

`func NewPayoutHeaderWithDefaults() *PayoutHeader`

NewPayoutHeaderWithDefaults instantiates a new PayoutHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *PayoutHeader) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *PayoutHeader) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *PayoutHeader) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetPayoutId

`func (o *PayoutHeader) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PayoutHeader) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PayoutHeader) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetRole

`func (o *PayoutHeader) GetRole() AssetId`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PayoutHeader) GetRoleOk() (*AssetId, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PayoutHeader) SetRole(v AssetId)`

SetRole sets Role field to given value.


### GetStatus

`func (o *PayoutHeader) GetStatus() PayoutStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutHeader) GetStatusOk() (*PayoutStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutHeader) SetStatus(v PayoutStatus)`

SetStatus sets Status field to given value.


### GetWithdrawalId

`func (o *PayoutHeader) GetWithdrawalId() string`

GetWithdrawalId returns the WithdrawalId field if non-nil, zero value otherwise.

### GetWithdrawalIdOk

`func (o *PayoutHeader) GetWithdrawalIdOk() (*string, bool)`

GetWithdrawalIdOk returns a tuple with the WithdrawalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalId

`func (o *PayoutHeader) SetWithdrawalId(v string)`

SetWithdrawalId sets WithdrawalId field to given value.

### HasWithdrawalId

`func (o *PayoutHeader) HasWithdrawalId() bool`

HasWithdrawalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


