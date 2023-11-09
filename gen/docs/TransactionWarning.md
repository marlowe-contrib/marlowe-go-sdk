# TransactionWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskedToDeposit** | **int32** |  | 
**InAccount** | [**Party**](Party.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**Party** | [**Party**](Party.md) |  | 
**Account** | [**Party**](Party.md) |  | 
**AskedToPay** | **int32** |  | 
**ToPayee** | [**Payee**](Payee.md) |  | 
**ButOnlyPaid** | **int32** |  | 
**HadValue** | **int32** |  | 
**IsNowAssigned** | **int32** |  | 
**ValueId** | **string** |  | 

## Methods

### NewTransactionWarning

`func NewTransactionWarning(askedToDeposit int32, inAccount Party, ofToken Token, party Party, account Party, askedToPay int32, toPayee Payee, butOnlyPaid int32, hadValue int32, isNowAssigned int32, valueId string, ) *TransactionWarning`

NewTransactionWarning instantiates a new TransactionWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWarningWithDefaults

`func NewTransactionWarningWithDefaults() *TransactionWarning`

NewTransactionWarningWithDefaults instantiates a new TransactionWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskedToDeposit

`func (o *TransactionWarning) GetAskedToDeposit() int32`

GetAskedToDeposit returns the AskedToDeposit field if non-nil, zero value otherwise.

### GetAskedToDepositOk

`func (o *TransactionWarning) GetAskedToDepositOk() (*int32, bool)`

GetAskedToDepositOk returns a tuple with the AskedToDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedToDeposit

`func (o *TransactionWarning) SetAskedToDeposit(v int32)`

SetAskedToDeposit sets AskedToDeposit field to given value.


### GetInAccount

`func (o *TransactionWarning) GetInAccount() Party`

GetInAccount returns the InAccount field if non-nil, zero value otherwise.

### GetInAccountOk

`func (o *TransactionWarning) GetInAccountOk() (*Party, bool)`

GetInAccountOk returns a tuple with the InAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAccount

`func (o *TransactionWarning) SetInAccount(v Party)`

SetInAccount sets InAccount field to given value.


### GetOfToken

`func (o *TransactionWarning) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *TransactionWarning) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *TransactionWarning) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetParty

`func (o *TransactionWarning) GetParty() Party`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *TransactionWarning) GetPartyOk() (*Party, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *TransactionWarning) SetParty(v Party)`

SetParty sets Party field to given value.


### GetAccount

`func (o *TransactionWarning) GetAccount() Party`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransactionWarning) GetAccountOk() (*Party, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransactionWarning) SetAccount(v Party)`

SetAccount sets Account field to given value.


### GetAskedToPay

`func (o *TransactionWarning) GetAskedToPay() int32`

GetAskedToPay returns the AskedToPay field if non-nil, zero value otherwise.

### GetAskedToPayOk

`func (o *TransactionWarning) GetAskedToPayOk() (*int32, bool)`

GetAskedToPayOk returns a tuple with the AskedToPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedToPay

`func (o *TransactionWarning) SetAskedToPay(v int32)`

SetAskedToPay sets AskedToPay field to given value.


### GetToPayee

`func (o *TransactionWarning) GetToPayee() Payee`

GetToPayee returns the ToPayee field if non-nil, zero value otherwise.

### GetToPayeeOk

`func (o *TransactionWarning) GetToPayeeOk() (*Payee, bool)`

GetToPayeeOk returns a tuple with the ToPayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPayee

`func (o *TransactionWarning) SetToPayee(v Payee)`

SetToPayee sets ToPayee field to given value.


### GetButOnlyPaid

`func (o *TransactionWarning) GetButOnlyPaid() int32`

GetButOnlyPaid returns the ButOnlyPaid field if non-nil, zero value otherwise.

### GetButOnlyPaidOk

`func (o *TransactionWarning) GetButOnlyPaidOk() (*int32, bool)`

GetButOnlyPaidOk returns a tuple with the ButOnlyPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButOnlyPaid

`func (o *TransactionWarning) SetButOnlyPaid(v int32)`

SetButOnlyPaid sets ButOnlyPaid field to given value.


### GetHadValue

`func (o *TransactionWarning) GetHadValue() int32`

GetHadValue returns the HadValue field if non-nil, zero value otherwise.

### GetHadValueOk

`func (o *TransactionWarning) GetHadValueOk() (*int32, bool)`

GetHadValueOk returns a tuple with the HadValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadValue

`func (o *TransactionWarning) SetHadValue(v int32)`

SetHadValue sets HadValue field to given value.


### GetIsNowAssigned

`func (o *TransactionWarning) GetIsNowAssigned() int32`

GetIsNowAssigned returns the IsNowAssigned field if non-nil, zero value otherwise.

### GetIsNowAssignedOk

`func (o *TransactionWarning) GetIsNowAssignedOk() (*int32, bool)`

GetIsNowAssignedOk returns a tuple with the IsNowAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNowAssigned

`func (o *TransactionWarning) SetIsNowAssigned(v int32)`

SetIsNowAssigned sets IsNowAssigned field to given value.


### GetValueId

`func (o *TransactionWarning) GetValueId() string`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *TransactionWarning) GetValueIdOk() (*string, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *TransactionWarning) SetValueId(v string)`

SetValueId sets ValueId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


