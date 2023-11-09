# TransactionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**IntervalError**](IntervalError.md) |  | 
**Error** | **string** |  | 

## Methods

### NewTransactionError

`func NewTransactionError(context IntervalError, error_ string, ) *TransactionError`

NewTransactionError instantiates a new TransactionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionErrorWithDefaults

`func NewTransactionErrorWithDefaults() *TransactionError`

NewTransactionErrorWithDefaults instantiates a new TransactionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TransactionError) GetContext() IntervalError`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransactionError) GetContextOk() (*IntervalError, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransactionError) SetContext(v IntervalError)`

SetContext sets Context field to given value.


### GetError

`func (o *TransactionError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TransactionError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TransactionError) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


