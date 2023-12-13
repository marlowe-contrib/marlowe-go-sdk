# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAccount** | [**Party**](Party.md) |  | 
**Pay** | [**Value**](Value.md) |  | 
**Then** | [**Contract**](Contract.md) |  | 
**To** | [**Payee**](Payee.md) |  | 
**Token** | [**Token**](Token.md) |  | 
**Else** | [**Contract**](Contract.md) |  | 
**If** | [**Observation**](Observation.md) |  | 
**Timeout** | **int64** |  | 
**TimeoutContinuation** | [**Contract**](Contract.md) |  | 
**When** | [**[]Case**](Case.md) |  | 
**Be** | [**Value**](Value.md) |  | 
**Let** | **string** |  | 
**Assert** | [**Observation**](Observation.md) |  | 

## Methods

### NewContract

`func NewContract(fromAccount Party, pay Value, then Contract, to Payee, token Token, else_ Contract, if_ Observation, timeout int64, timeoutContinuation Contract, when []Case, be Value, let string, assert Observation, ) *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAccount

`func (o *Contract) GetFromAccount() Party`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *Contract) GetFromAccountOk() (*Party, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *Contract) SetFromAccount(v Party)`

SetFromAccount sets FromAccount field to given value.


### GetPay

`func (o *Contract) GetPay() Value`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *Contract) GetPayOk() (*Value, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *Contract) SetPay(v Value)`

SetPay sets Pay field to given value.


### GetThen

`func (o *Contract) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *Contract) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *Contract) SetThen(v Contract)`

SetThen sets Then field to given value.


### GetTo

`func (o *Contract) GetTo() Payee`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Contract) GetToOk() (*Payee, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Contract) SetTo(v Payee)`

SetTo sets To field to given value.


### GetToken

`func (o *Contract) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Contract) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Contract) SetToken(v Token)`

SetToken sets Token field to given value.


### GetElse

`func (o *Contract) GetElse() Contract`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *Contract) GetElseOk() (*Contract, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *Contract) SetElse(v Contract)`

SetElse sets Else field to given value.


### GetIf

`func (o *Contract) GetIf() Observation`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *Contract) GetIfOk() (*Observation, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *Contract) SetIf(v Observation)`

SetIf sets If field to given value.


### GetTimeout

`func (o *Contract) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Contract) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Contract) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.


### GetTimeoutContinuation

`func (o *Contract) GetTimeoutContinuation() Contract`

GetTimeoutContinuation returns the TimeoutContinuation field if non-nil, zero value otherwise.

### GetTimeoutContinuationOk

`func (o *Contract) GetTimeoutContinuationOk() (*Contract, bool)`

GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutContinuation

`func (o *Contract) SetTimeoutContinuation(v Contract)`

SetTimeoutContinuation sets TimeoutContinuation field to given value.


### GetWhen

`func (o *Contract) GetWhen() []Case`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *Contract) GetWhenOk() (*[]Case, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *Contract) SetWhen(v []Case)`

SetWhen sets When field to given value.


### GetBe

`func (o *Contract) GetBe() Value`

GetBe returns the Be field if non-nil, zero value otherwise.

### GetBeOk

`func (o *Contract) GetBeOk() (*Value, bool)`

GetBeOk returns a tuple with the Be field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBe

`func (o *Contract) SetBe(v Value)`

SetBe sets Be field to given value.


### GetLet

`func (o *Contract) GetLet() string`

GetLet returns the Let field if non-nil, zero value otherwise.

### GetLetOk

`func (o *Contract) GetLetOk() (*string, bool)`

GetLetOk returns a tuple with the Let field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLet

`func (o *Contract) SetLet(v string)`

SetLet sets Let field to given value.


### GetAssert

`func (o *Contract) GetAssert() Observation`

GetAssert returns the Assert field if non-nil, zero value otherwise.

### GetAssertOk

`func (o *Contract) GetAssertOk() (*Observation, bool)`

GetAssertOk returns a tuple with the Assert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssert

`func (o *Contract) SetAssert(v Observation)`

SetAssert sets Assert field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


