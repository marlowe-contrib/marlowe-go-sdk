# ContractObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAccount** | [**PartyObject**](PartyObject.md) |  | 
**Pay** | [**ValueObject**](ValueObject.md) |  | 
**Then** | [**ContractObject**](ContractObject.md) |  | 
**To** | [**PayeeObject**](PayeeObject.md) |  | 
**Token** | [**TokenObject**](TokenObject.md) |  | 
**Else** | [**ContractObject**](ContractObject.md) |  | 
**If** | [**ObservationObject**](ObservationObject.md) |  | 
**Timeout** | **int32** |  | 
**TimeoutContinuation** | [**ContractObject**](ContractObject.md) |  | 
**When** | [**[]CaseObject**](CaseObject.md) |  | 
**Be** | [**ValueObject**](ValueObject.md) |  | 
**Let** | **string** |  | 
**Assert** | [**ObservationObject**](ObservationObject.md) |  | 
**Ref** | **string** | An arbitrary text identifier for an object in a Marlowe object bundle. | 

## Methods

### NewContractObject

`func NewContractObject(fromAccount PartyObject, pay ValueObject, then ContractObject, to PayeeObject, token TokenObject, else_ ContractObject, if_ ObservationObject, timeout int32, timeoutContinuation ContractObject, when []CaseObject, be ValueObject, let string, assert ObservationObject, ref string, ) *ContractObject`

NewContractObject instantiates a new ContractObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractObjectWithDefaults

`func NewContractObjectWithDefaults() *ContractObject`

NewContractObjectWithDefaults instantiates a new ContractObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAccount

`func (o *ContractObject) GetFromAccount() PartyObject`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *ContractObject) GetFromAccountOk() (*PartyObject, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *ContractObject) SetFromAccount(v PartyObject)`

SetFromAccount sets FromAccount field to given value.


### GetPay

`func (o *ContractObject) GetPay() ValueObject`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *ContractObject) GetPayOk() (*ValueObject, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *ContractObject) SetPay(v ValueObject)`

SetPay sets Pay field to given value.


### GetThen

`func (o *ContractObject) GetThen() ContractObject`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *ContractObject) GetThenOk() (*ContractObject, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *ContractObject) SetThen(v ContractObject)`

SetThen sets Then field to given value.


### GetTo

`func (o *ContractObject) GetTo() PayeeObject`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContractObject) GetToOk() (*PayeeObject, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContractObject) SetTo(v PayeeObject)`

SetTo sets To field to given value.


### GetToken

`func (o *ContractObject) GetToken() TokenObject`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ContractObject) GetTokenOk() (*TokenObject, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ContractObject) SetToken(v TokenObject)`

SetToken sets Token field to given value.


### GetElse

`func (o *ContractObject) GetElse() ContractObject`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *ContractObject) GetElseOk() (*ContractObject, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *ContractObject) SetElse(v ContractObject)`

SetElse sets Else field to given value.


### GetIf

`func (o *ContractObject) GetIf() ObservationObject`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *ContractObject) GetIfOk() (*ObservationObject, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *ContractObject) SetIf(v ObservationObject)`

SetIf sets If field to given value.


### GetTimeout

`func (o *ContractObject) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ContractObject) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ContractObject) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetTimeoutContinuation

`func (o *ContractObject) GetTimeoutContinuation() ContractObject`

GetTimeoutContinuation returns the TimeoutContinuation field if non-nil, zero value otherwise.

### GetTimeoutContinuationOk

`func (o *ContractObject) GetTimeoutContinuationOk() (*ContractObject, bool)`

GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutContinuation

`func (o *ContractObject) SetTimeoutContinuation(v ContractObject)`

SetTimeoutContinuation sets TimeoutContinuation field to given value.


### GetWhen

`func (o *ContractObject) GetWhen() []CaseObject`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *ContractObject) GetWhenOk() (*[]CaseObject, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *ContractObject) SetWhen(v []CaseObject)`

SetWhen sets When field to given value.


### GetBe

`func (o *ContractObject) GetBe() ValueObject`

GetBe returns the Be field if non-nil, zero value otherwise.

### GetBeOk

`func (o *ContractObject) GetBeOk() (*ValueObject, bool)`

GetBeOk returns a tuple with the Be field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBe

`func (o *ContractObject) SetBe(v ValueObject)`

SetBe sets Be field to given value.


### GetLet

`func (o *ContractObject) GetLet() string`

GetLet returns the Let field if non-nil, zero value otherwise.

### GetLetOk

`func (o *ContractObject) GetLetOk() (*string, bool)`

GetLetOk returns a tuple with the Let field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLet

`func (o *ContractObject) SetLet(v string)`

SetLet sets Let field to given value.


### GetAssert

`func (o *ContractObject) GetAssert() ObservationObject`

GetAssert returns the Assert field if non-nil, zero value otherwise.

### GetAssertOk

`func (o *ContractObject) GetAssertOk() (*ObservationObject, bool)`

GetAssertOk returns a tuple with the Assert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssert

`func (o *ContractObject) SetAssert(v ObservationObject)`

SetAssert sets Assert field to given value.


### GetRef

`func (o *ContractObject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ContractObject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ContractObject) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


