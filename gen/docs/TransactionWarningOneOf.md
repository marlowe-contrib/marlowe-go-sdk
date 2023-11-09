# TransactionWarningOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskedToDeposit** | **int32** |  | 
**InAccount** | [**Party**](Party.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**Party** | [**Party**](Party.md) |  | 

## Methods

### NewTransactionWarningOneOf

`func NewTransactionWarningOneOf(askedToDeposit int32, inAccount Party, ofToken Token, party Party, ) *TransactionWarningOneOf`

NewTransactionWarningOneOf instantiates a new TransactionWarningOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWarningOneOfWithDefaults

`func NewTransactionWarningOneOfWithDefaults() *TransactionWarningOneOf`

NewTransactionWarningOneOfWithDefaults instantiates a new TransactionWarningOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskedToDeposit

`func (o *TransactionWarningOneOf) GetAskedToDeposit() int32`

GetAskedToDeposit returns the AskedToDeposit field if non-nil, zero value otherwise.

### GetAskedToDepositOk

`func (o *TransactionWarningOneOf) GetAskedToDepositOk() (*int32, bool)`

GetAskedToDepositOk returns a tuple with the AskedToDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskedToDeposit

`func (o *TransactionWarningOneOf) SetAskedToDeposit(v int32)`

SetAskedToDeposit sets AskedToDeposit field to given value.


### GetInAccount

`func (o *TransactionWarningOneOf) GetInAccount() Party`

GetInAccount returns the InAccount field if non-nil, zero value otherwise.

### GetInAccountOk

`func (o *TransactionWarningOneOf) GetInAccountOk() (*Party, bool)`

GetInAccountOk returns a tuple with the InAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAccount

`func (o *TransactionWarningOneOf) SetInAccount(v Party)`

SetInAccount sets InAccount field to given value.


### GetOfToken

`func (o *TransactionWarningOneOf) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *TransactionWarningOneOf) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *TransactionWarningOneOf) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetParty

`func (o *TransactionWarningOneOf) GetParty() Party`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *TransactionWarningOneOf) GetPartyOk() (*Party, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *TransactionWarningOneOf) SetParty(v Party)`

SetParty sets Party field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


