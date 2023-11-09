# TransactionWarningOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Party**](Party.md) |  | 
**AskedToPay** | **int32** |  | 
**OfToken** | [**Token**](Token.md) |  | 
**ToPayee** | [**Payee**](Payee.md) |  | 

## Methods

### NewTransactionWarningOneOf1

`func NewTransactionWarningOneOf1(account Party, askedToPay int32, ofToken Token, toPayee Payee, ) *TransactionWarningOneOf1`

NewTransactionWarningOneOf1 instantiates a new TransactionWarningOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWarningOneOf1WithDefaults

`func NewTransactionWarningOneOf1WithDefaults() *TransactionWarningOneOf1`

NewTransactionWarningOneOf1WithDefaults instantiates a new TransactionWarningOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TransactionWarningOneOf1) GetAccount() Party`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransactionWarningOneOf1) GetAccountOk() (*Party, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransactionWarningOneOf1) SetAccount(v Party)`

SetAccount sets Account field to given value.


### GetAskedToPay

`func (o *TransactionWarningOneOf1) GetAskedToPay() int32`

GetAskedToPay returns the AskedToPay field if non-nil, zero value otherwise.

### GetAskedToPayOk

`func (o *TransactionWarningOneOf1) GetAskedToPayOk() (*int32, bool)`

GetAskedToPayOk returns a tuple with the AskedToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedToPay

`func (o *TransactionWarningOneOf1) SetAskedToPay(v int32)`

SetAskedToPay sets AskedToPay field to given value.


### GetOfToken

`func (o *TransactionWarningOneOf1) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *TransactionWarningOneOf1) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *TransactionWarningOneOf1) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetToPayee

`func (o *TransactionWarningOneOf1) GetToPayee() Payee`

GetToPayee returns the ToPayee field if non-nil, zero value otherwise.

### GetToPayeeOk

`func (o *TransactionWarningOneOf1) GetToPayeeOk() (*Payee, bool)`

GetToPayeeOk returns a tuple with the ToPayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPayee

`func (o *TransactionWarningOneOf1) SetToPayee(v Payee)`

SetToPayee sets ToPayee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


