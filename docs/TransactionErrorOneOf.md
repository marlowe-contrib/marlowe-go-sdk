# TransactionErrorOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**IntervalError**](IntervalError.md) |  | 
**Error** | **string** |  | 

## Methods

### NewTransactionErrorOneOf

`func NewTransactionErrorOneOf(context IntervalError, error_ string, ) *TransactionErrorOneOf`

NewTransactionErrorOneOf instantiates a new TransactionErrorOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionErrorOneOfWithDefaults

`func NewTransactionErrorOneOfWithDefaults() *TransactionErrorOneOf`

NewTransactionErrorOneOfWithDefaults instantiates a new TransactionErrorOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TransactionErrorOneOf) GetContext() IntervalError`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransactionErrorOneOf) GetContextOk() (*IntervalError, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransactionErrorOneOf) SetContext(v IntervalError)`

SetContext sets Context field to given value.


### GetError

`func (o *TransactionErrorOneOf) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TransactionErrorOneOf) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TransactionErrorOneOf) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


