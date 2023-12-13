# NonPositivePayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Party**](Party.md) |  | 
**AskedToPay** | **int32** |  | 
**OfToken** | [**Token**](Token.md) |  | 
**ToPayee** | [**Payee**](Payee.md) |  | 

## Methods

### NewNonPositivePayment

`func NewNonPositivePayment(account Party, askedToPay int32, ofToken Token, toPayee Payee, ) *NonPositivePayment`

NewNonPositivePayment instantiates a new NonPositivePayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonPositivePaymentWithDefaults

`func NewNonPositivePaymentWithDefaults() *NonPositivePayment`

NewNonPositivePaymentWithDefaults instantiates a new NonPositivePayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *NonPositivePayment) GetAccount() Party`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NonPositivePayment) GetAccountOk() (*Party, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NonPositivePayment) SetAccount(v Party)`

SetAccount sets Account field to given value.


### GetAskedToPay

`func (o *NonPositivePayment) GetAskedToPay() int32`

GetAskedToPay returns the AskedToPay field if non-nil, zero value otherwise.

### GetAskedToPayOk

`func (o *NonPositivePayment) GetAskedToPayOk() (*int32, bool)`

GetAskedToPayOk returns a tuple with the AskedToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedToPay

`func (o *NonPositivePayment) SetAskedToPay(v int32)`

SetAskedToPay sets AskedToPay field to given value.


### GetOfToken

`func (o *NonPositivePayment) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *NonPositivePayment) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *NonPositivePayment) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetToPayee

`func (o *NonPositivePayment) GetToPayee() Payee`

GetToPayee returns the ToPayee field if non-nil, zero value otherwise.

### GetToPayeeOk

`func (o *NonPositivePayment) GetToPayeeOk() (*Payee, bool)`

GetToPayeeOk returns a tuple with the ToPayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPayee

`func (o *NonPositivePayment) SetToPayee(v Payee)`

SetToPayee sets ToPayee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


