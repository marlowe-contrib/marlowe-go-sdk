# TransactionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | [**Contract**](Contract.md) |  | 
**Payments** | [**[]Payment**](Payment.md) |  | 
**State** | [**MarloweState**](MarloweState.md) |  | 
**Warnings** | [**[]TransactionWarning**](TransactionWarning.md) |  | 
**TransactionError** | [**TransactionError**](TransactionError.md) |  | 

## Methods

### NewTransactionOutput

`func NewTransactionOutput(contract Contract, payments []Payment, state MarloweState, warnings []TransactionWarning, transactionError TransactionError, ) *TransactionOutput`

NewTransactionOutput instantiates a new TransactionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionOutputWithDefaults

`func NewTransactionOutputWithDefaults() *TransactionOutput`

NewTransactionOutputWithDefaults instantiates a new TransactionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *TransactionOutput) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *TransactionOutput) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *TransactionOutput) SetContract(v Contract)`

SetContract sets Contract field to given value.


### GetPayments

`func (o *TransactionOutput) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *TransactionOutput) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *TransactionOutput) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.


### GetState

`func (o *TransactionOutput) GetState() MarloweState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TransactionOutput) GetStateOk() (*MarloweState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TransactionOutput) SetState(v MarloweState)`

SetState sets State field to given value.


### GetWarnings

`func (o *TransactionOutput) GetWarnings() []TransactionWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TransactionOutput) GetWarningsOk() (*[]TransactionWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TransactionOutput) SetWarnings(v []TransactionWarning)`

SetWarnings sets Warnings field to given value.


### GetTransactionError

`func (o *TransactionOutput) GetTransactionError() TransactionError`

GetTransactionError returns the TransactionError field if non-nil, zero value otherwise.

### GetTransactionErrorOk

`func (o *TransactionOutput) GetTransactionErrorOk() (*TransactionError, bool)`

GetTransactionErrorOk returns a tuple with the TransactionError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionError

`func (o *TransactionOutput) SetTransactionError(v TransactionError)`

SetTransactionError sets TransactionError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


