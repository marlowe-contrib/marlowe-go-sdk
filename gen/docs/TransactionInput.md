# TransactionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxInputs** | [**[]Input**](Input.md) |  | 
**TxInterval** | [**TransactionInputTxInterval**](TransactionInputTxInterval.md) |  | 

## Methods

### NewTransactionInput

`func NewTransactionInput(txInputs []Input, txInterval TransactionInputTxInterval, ) *TransactionInput`

NewTransactionInput instantiates a new TransactionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInputWithDefaults

`func NewTransactionInputWithDefaults() *TransactionInput`

NewTransactionInputWithDefaults instantiates a new TransactionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxInputs

`func (o *TransactionInput) GetTxInputs() []Input`

GetTxInputs returns the TxInputs field if non-nil, zero value otherwise.

### GetTxInputsOk

`func (o *TransactionInput) GetTxInputsOk() (*[]Input, bool)`

GetTxInputsOk returns a tuple with the TxInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxInputs

`func (o *TransactionInput) SetTxInputs(v []Input)`

SetTxInputs sets TxInputs field to given value.


### GetTxInterval

`func (o *TransactionInput) GetTxInterval() TransactionInputTxInterval`

GetTxInterval returns the TxInterval field if non-nil, zero value otherwise.

### GetTxIntervalOk

`func (o *TransactionInput) GetTxIntervalOk() (*TransactionInputTxInterval, bool)`

GetTxIntervalOk returns a tuple with the TxInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxInterval

`func (o *TransactionInput) SetTxInterval(v TransactionInputTxInterval)`

SetTxInterval sets TxInterval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


