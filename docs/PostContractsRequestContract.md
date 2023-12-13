# PostContractsRequestContract

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

### NewPostContractsRequestContract

`func NewPostContractsRequestContract(fromAccount Party, pay Value, then Contract, to Payee, token Token, else_ Contract, if_ Observation, timeout int64, timeoutContinuation Contract, when []Case, be Value, let string, assert Observation, ) *PostContractsRequestContract`

NewPostContractsRequestContract instantiates a new PostContractsRequestContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostContractsRequestContractWithDefaults

`func NewPostContractsRequestContractWithDefaults() *PostContractsRequestContract`

NewPostContractsRequestContractWithDefaults instantiates a new PostContractsRequestContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAccount

`func (o *PostContractsRequestContract) GetFromAccount() Party`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *PostContractsRequestContract) GetFromAccountOk() (*Party, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *PostContractsRequestContract) SetFromAccount(v Party)`

SetFromAccount sets FromAccount field to given value.


### GetPay

`func (o *PostContractsRequestContract) GetPay() Value`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *PostContractsRequestContract) GetPayOk() (*Value, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *PostContractsRequestContract) SetPay(v Value)`

SetPay sets Pay field to given value.


### GetThen

`func (o *PostContractsRequestContract) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *PostContractsRequestContract) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *PostContractsRequestContract) SetThen(v Contract)`

SetThen sets Then field to given value.


### GetTo

`func (o *PostContractsRequestContract) GetTo() Payee`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PostContractsRequestContract) GetToOk() (*Payee, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PostContractsRequestContract) SetTo(v Payee)`

SetTo sets To field to given value.


### GetToken

`func (o *PostContractsRequestContract) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PostContractsRequestContract) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PostContractsRequestContract) SetToken(v Token)`

SetToken sets Token field to given value.


### GetElse

`func (o *PostContractsRequestContract) GetElse() Contract`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *PostContractsRequestContract) GetElseOk() (*Contract, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *PostContractsRequestContract) SetElse(v Contract)`

SetElse sets Else field to given value.


### GetIf

`func (o *PostContractsRequestContract) GetIf() Observation`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *PostContractsRequestContract) GetIfOk() (*Observation, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *PostContractsRequestContract) SetIf(v Observation)`

SetIf sets If field to given value.


### GetTimeout

`func (o *PostContractsRequestContract) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PostContractsRequestContract) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PostContractsRequestContract) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.


### GetTimeoutContinuation

`func (o *PostContractsRequestContract) GetTimeoutContinuation() Contract`

GetTimeoutContinuation returns the TimeoutContinuation field if non-nil, zero value otherwise.

### GetTimeoutContinuationOk

`func (o *PostContractsRequestContract) GetTimeoutContinuationOk() (*Contract, bool)`

GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutContinuation

`func (o *PostContractsRequestContract) SetTimeoutContinuation(v Contract)`

SetTimeoutContinuation sets TimeoutContinuation field to given value.


### GetWhen

`func (o *PostContractsRequestContract) GetWhen() []Case`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *PostContractsRequestContract) GetWhenOk() (*[]Case, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *PostContractsRequestContract) SetWhen(v []Case)`

SetWhen sets When field to given value.


### GetBe

`func (o *PostContractsRequestContract) GetBe() Value`

GetBe returns the Be field if non-nil, zero value otherwise.

### GetBeOk

`func (o *PostContractsRequestContract) GetBeOk() (*Value, bool)`

GetBeOk returns a tuple with the Be field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBe

`func (o *PostContractsRequestContract) SetBe(v Value)`

SetBe sets Be field to given value.


### GetLet

`func (o *PostContractsRequestContract) GetLet() string`

GetLet returns the Let field if non-nil, zero value otherwise.

### GetLetOk

`func (o *PostContractsRequestContract) GetLetOk() (*string, bool)`

GetLetOk returns a tuple with the Let field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLet

`func (o *PostContractsRequestContract) SetLet(v string)`

SetLet sets Let field to given value.


### GetAssert

`func (o *PostContractsRequestContract) GetAssert() Observation`

GetAssert returns the Assert field if non-nil, zero value otherwise.

### GetAssertOk

`func (o *PostContractsRequestContract) GetAssertOk() (*Observation, bool)`

GetAssertOk returns a tuple with the Assert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssert

`func (o *PostContractsRequestContract) SetAssert(v Observation)`

SetAssert sets Assert field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


