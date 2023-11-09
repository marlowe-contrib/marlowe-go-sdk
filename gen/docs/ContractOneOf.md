# ContractOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAccount** | [**Party**](Party.md) |  | 
**Pay** | [**Value**](Value.md) |  | 
**Then** | [**Contract**](Contract.md) |  | 
**To** | [**Payee**](Payee.md) |  | 
**Token** | [**Token**](Token.md) |  | 

## Methods

### NewContractOneOf

`func NewContractOneOf(fromAccount Party, pay Value, then Contract, to Payee, token Token, ) *ContractOneOf`

NewContractOneOf instantiates a new ContractOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractOneOfWithDefaults

`func NewContractOneOfWithDefaults() *ContractOneOf`

NewContractOneOfWithDefaults instantiates a new ContractOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAccount

`func (o *ContractOneOf) GetFromAccount() Party`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *ContractOneOf) GetFromAccountOk() (*Party, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *ContractOneOf) SetFromAccount(v Party)`

SetFromAccount sets FromAccount field to given value.


### GetPay

`func (o *ContractOneOf) GetPay() Value`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *ContractOneOf) GetPayOk() (*Value, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *ContractOneOf) SetPay(v Value)`

SetPay sets Pay field to given value.


### GetThen

`func (o *ContractOneOf) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *ContractOneOf) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *ContractOneOf) SetThen(v Contract)`

SetThen sets Then field to given value.


### GetTo

`func (o *ContractOneOf) GetTo() Payee`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContractOneOf) GetToOk() (*Payee, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContractOneOf) SetTo(v Payee)`

SetTo sets To field to given value.


### GetToken

`func (o *ContractOneOf) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ContractOneOf) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ContractOneOf) SetToken(v Token)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


