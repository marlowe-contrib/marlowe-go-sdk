# ContractObjectOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAccount** | [**PartyObject**](PartyObject.md) |  | 
**Pay** | [**ValueObject**](ValueObject.md) |  | 
**Then** | [**ContractObject**](ContractObject.md) |  | 
**To** | [**PayeeObject**](PayeeObject.md) |  | 
**Token** | [**TokenObject**](TokenObject.md) |  | 

## Methods

### NewContractObjectOneOf

`func NewContractObjectOneOf(fromAccount PartyObject, pay ValueObject, then ContractObject, to PayeeObject, token TokenObject, ) *ContractObjectOneOf`

NewContractObjectOneOf instantiates a new ContractObjectOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractObjectOneOfWithDefaults

`func NewContractObjectOneOfWithDefaults() *ContractObjectOneOf`

NewContractObjectOneOfWithDefaults instantiates a new ContractObjectOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAccount

`func (o *ContractObjectOneOf) GetFromAccount() PartyObject`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *ContractObjectOneOf) GetFromAccountOk() (*PartyObject, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *ContractObjectOneOf) SetFromAccount(v PartyObject)`

SetFromAccount sets FromAccount field to given value.


### GetPay

`func (o *ContractObjectOneOf) GetPay() ValueObject`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *ContractObjectOneOf) GetPayOk() (*ValueObject, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *ContractObjectOneOf) SetPay(v ValueObject)`

SetPay sets Pay field to given value.


### GetThen

`func (o *ContractObjectOneOf) GetThen() ContractObject`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *ContractObjectOneOf) GetThenOk() (*ContractObject, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *ContractObjectOneOf) SetThen(v ContractObject)`

SetThen sets Then field to given value.


### GetTo

`func (o *ContractObjectOneOf) GetTo() PayeeObject`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContractObjectOneOf) GetToOk() (*PayeeObject, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContractObjectOneOf) SetTo(v PayeeObject)`

SetTo sets To field to given value.


### GetToken

`func (o *ContractObjectOneOf) GetToken() TokenObject`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ContractObjectOneOf) GetTokenOk() (*TokenObject, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ContractObjectOneOf) SetToken(v TokenObject)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


