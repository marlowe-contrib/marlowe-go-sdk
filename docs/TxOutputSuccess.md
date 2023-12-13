# TxOutputSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | [**Contract**](Contract.md) |  | 
**Payments** | [**[]Payment**](Payment.md) |  | 
**State** | [**MarloweState**](MarloweState.md) |  | 
**Warnings** | [**[]TransactionWarning**](TransactionWarning.md) |  | 

## Methods

### NewTxOutputSuccess

`func NewTxOutputSuccess(contract Contract, payments []Payment, state MarloweState, warnings []TransactionWarning, ) *TxOutputSuccess`

NewTxOutputSuccess instantiates a new TxOutputSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxOutputSuccessWithDefaults

`func NewTxOutputSuccessWithDefaults() *TxOutputSuccess`

NewTxOutputSuccessWithDefaults instantiates a new TxOutputSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *TxOutputSuccess) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *TxOutputSuccess) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *TxOutputSuccess) SetContract(v Contract)`

SetContract sets Contract field to given value.


### GetPayments

`func (o *TxOutputSuccess) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *TxOutputSuccess) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *TxOutputSuccess) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.


### GetState

`func (o *TxOutputSuccess) GetState() MarloweState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TxOutputSuccess) GetStateOk() (*MarloweState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TxOutputSuccess) SetState(v MarloweState)`

SetState sets State field to given value.


### GetWarnings

`func (o *TxOutputSuccess) GetWarnings() []TransactionWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TxOutputSuccess) GetWarningsOk() (*[]TransactionWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TxOutputSuccess) SetWarnings(v []TransactionWarning)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


