# Withdrawal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | Pointer to [**BlockHeader**](BlockHeader.md) |  | [optional] 
**Payouts** | [**[]PayoutHeader**](PayoutHeader.md) |  | 
**Status** | [**TxStatus**](TxStatus.md) |  | 
**WithdrawalId** | **string** | The hex-encoded identifier of a Cardano transaction | 

## Methods

### NewWithdrawal

`func NewWithdrawal(payouts []PayoutHeader, status TxStatus, withdrawalId string, ) *Withdrawal`

NewWithdrawal instantiates a new Withdrawal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalWithDefaults

`func NewWithdrawalWithDefaults() *Withdrawal`

NewWithdrawalWithDefaults instantiates a new Withdrawal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *Withdrawal) GetBlock() BlockHeader`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *Withdrawal) GetBlockOk() (*BlockHeader, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *Withdrawal) SetBlock(v BlockHeader)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *Withdrawal) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetPayouts

`func (o *Withdrawal) GetPayouts() []PayoutHeader`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *Withdrawal) GetPayoutsOk() (*[]PayoutHeader, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *Withdrawal) SetPayouts(v []PayoutHeader)`

SetPayouts sets Payouts field to given value.


### GetStatus

`func (o *Withdrawal) GetStatus() TxStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Withdrawal) GetStatusOk() (*TxStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Withdrawal) SetStatus(v TxStatus)`

SetStatus sets Status field to given value.


### GetWithdrawalId

`func (o *Withdrawal) GetWithdrawalId() string`

GetWithdrawalId returns the WithdrawalId field if non-nil, zero value otherwise.

### GetWithdrawalIdOk

`func (o *Withdrawal) GetWithdrawalIdOk() (*string, bool)`

GetWithdrawalIdOk returns a tuple with the WithdrawalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalId

`func (o *Withdrawal) SetWithdrawalId(v string)`

SetWithdrawalId sets WithdrawalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


