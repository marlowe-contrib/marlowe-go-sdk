# PartialPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Party**](Party.md) |  | 
**AskedToPay** | **int32** |  | 
**ButOnlyPaid** | **int32** |  | 
**OfToken** | [**Token**](Token.md) |  | 
**ToPayee** | [**Payee**](Payee.md) |  | 

## Methods

### NewPartialPayment

`func NewPartialPayment(account Party, askedToPay int32, butOnlyPaid int32, ofToken Token, toPayee Payee, ) *PartialPayment`

NewPartialPayment instantiates a new PartialPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialPaymentWithDefaults

`func NewPartialPaymentWithDefaults() *PartialPayment`

NewPartialPaymentWithDefaults instantiates a new PartialPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PartialPayment) GetAccount() Party`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartialPayment) GetAccountOk() (*Party, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartialPayment) SetAccount(v Party)`

SetAccount sets Account field to given value.


### GetAskedToPay

`func (o *PartialPayment) GetAskedToPay() int32`

GetAskedToPay returns the AskedToPay field if non-nil, zero value otherwise.

### GetAskedToPayOk

`func (o *PartialPayment) GetAskedToPayOk() (*int32, bool)`

GetAskedToPayOk returns a tuple with the AskedToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedToPay

`func (o *PartialPayment) SetAskedToPay(v int32)`

SetAskedToPay sets AskedToPay field to given value.


### GetButOnlyPaid

`func (o *PartialPayment) GetButOnlyPaid() int32`

GetButOnlyPaid returns the ButOnlyPaid field if non-nil, zero value otherwise.

### GetButOnlyPaidOk

`func (o *PartialPayment) GetButOnlyPaidOk() (*int32, bool)`

GetButOnlyPaidOk returns a tuple with the ButOnlyPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButOnlyPaid

`func (o *PartialPayment) SetButOnlyPaid(v int32)`

SetButOnlyPaid sets ButOnlyPaid field to given value.


### GetOfToken

`func (o *PartialPayment) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *PartialPayment) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *PartialPayment) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetToPayee

`func (o *PartialPayment) GetToPayee() Payee`

GetToPayee returns the ToPayee field if non-nil, zero value otherwise.

### GetToPayeeOk

`func (o *PartialPayment) GetToPayeeOk() (*Payee, bool)`

GetToPayeeOk returns a tuple with the ToPayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPayee

`func (o *PartialPayment) SetToPayee(v Payee)`

SetToPayee sets ToPayee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


