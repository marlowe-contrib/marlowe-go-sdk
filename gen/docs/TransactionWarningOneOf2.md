# TransactionWarningOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Party**](Party.md) |  | 
**AskedToPay** | **int32** |  | 
**ButOnlyPaid** | **int32** |  | 
**OfToken** | [**Token**](Token.md) |  | 
**ToPayee** | [**Payee**](Payee.md) |  | 

## Methods

### NewTransactionWarningOneOf2

`func NewTransactionWarningOneOf2(account Party, askedToPay int32, butOnlyPaid int32, ofToken Token, toPayee Payee, ) *TransactionWarningOneOf2`

NewTransactionWarningOneOf2 instantiates a new TransactionWarningOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWarningOneOf2WithDefaults

`func NewTransactionWarningOneOf2WithDefaults() *TransactionWarningOneOf2`

NewTransactionWarningOneOf2WithDefaults instantiates a new TransactionWarningOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TransactionWarningOneOf2) GetAccount() Party`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransactionWarningOneOf2) GetAccountOk() (*Party, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransactionWarningOneOf2) SetAccount(v Party)`

SetAccount sets Account field to given value.


### GetAskedToPay

`func (o *TransactionWarningOneOf2) GetAskedToPay() int32`

GetAskedToPay returns the AskedToPay field if non-nil, zero value otherwise.

### GetAskedToPayOk

`func (o *TransactionWarningOneOf2) GetAskedToPayOk() (*int32, bool)`

GetAskedToPayOk returns a tuple with the AskedToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedToPay

`func (o *TransactionWarningOneOf2) SetAskedToPay(v int32)`

SetAskedToPay sets AskedToPay field to given value.


### GetButOnlyPaid

`func (o *TransactionWarningOneOf2) GetButOnlyPaid() int32`

GetButOnlyPaid returns the ButOnlyPaid field if non-nil, zero value otherwise.

### GetButOnlyPaidOk

`func (o *TransactionWarningOneOf2) GetButOnlyPaidOk() (*int32, bool)`

GetButOnlyPaidOk returns a tuple with the ButOnlyPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButOnlyPaid

`func (o *TransactionWarningOneOf2) SetButOnlyPaid(v int32)`

SetButOnlyPaid sets ButOnlyPaid field to given value.


### GetOfToken

`func (o *TransactionWarningOneOf2) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *TransactionWarningOneOf2) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *TransactionWarningOneOf2) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetToPayee

`func (o *TransactionWarningOneOf2) GetToPayee() Payee`

GetToPayee returns the ToPayee field if non-nil, zero value otherwise.

### GetToPayeeOk

`func (o *TransactionWarningOneOf2) GetToPayeeOk() (*Payee, bool)`

GetToPayeeOk returns a tuple with the ToPayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPayee

`func (o *TransactionWarningOneOf2) SetToPayee(v Payee)`

SetToPayee sets ToPayee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

