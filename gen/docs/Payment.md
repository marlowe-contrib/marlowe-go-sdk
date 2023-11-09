# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**PaymentFrom** | [**Party**](Party.md) |  | 
**To** | [**Payee**](Payee.md) |  | 
**Token** | [**Token**](Token.md) |  | 

## Methods

### NewPayment

`func NewPayment(amount int32, paymentFrom Party, to Payee, token Token, ) *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Payment) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetPaymentFrom

`func (o *Payment) GetPaymentFrom() Party`

GetPaymentFrom returns the PaymentFrom field if non-nil, zero value otherwise.

### GetPaymentFromOk

`func (o *Payment) GetPaymentFromOk() (*Party, bool)`

GetPaymentFromOk returns a tuple with the PaymentFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrom

`func (o *Payment) SetPaymentFrom(v Party)`

SetPaymentFrom sets PaymentFrom field to given value.


### GetTo

`func (o *Payment) GetTo() Payee`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Payment) GetToOk() (*Payee, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Payment) SetTo(v Payee)`

SetTo sets To field to given value.


### GetToken

`func (o *Payment) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Payment) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Payment) SetToken(v Token)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


