# Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOfToken** | [**Token**](Token.md) |  | 
**InAccount** | [**Party**](Party.md) |  | 
**Negate** | [**Value**](Value.md) |  | 
**Add** | [**Value**](Value.md) |  | 
**And** | [**Value**](Value.md) |  | 
**Minus** | [**Value**](Value.md) |  | 
**Value** | [**Value**](Value.md) |  | 
**Multiply** | [**Value**](Value.md) |  | 
**Times** | [**Value**](Value.md) |  | 
**By** | [**Value**](Value.md) |  | 
**Divide** | [**Value**](Value.md) |  | 
**ValueOfChoice** | [**ChoiceId**](ChoiceId.md) |  | 
**UseValue** | **string** |  | 
**Else** | [**Value**](Value.md) |  | 
**If** | [**Observation**](Observation.md) |  | 
**Then** | [**Value**](Value.md) |  | 

## Methods

### NewValue

`func NewValue(amountOfToken Token, inAccount Party, negate Value, add Value, and Value, minus Value, value Value, multiply Value, times Value, by Value, divide Value, valueOfChoice ChoiceId, useValue string, else_ Value, if_ Observation, then Value, ) *Value`

NewValue instantiates a new Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithDefaults

`func NewValueWithDefaults() *Value`

NewValueWithDefaults instantiates a new Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOfToken

`func (o *Value) GetAmountOfToken() Token`

GetAmountOfToken returns the AmountOfToken field if non-nil, zero value otherwise.

### GetAmountOfTokenOk

`func (o *Value) GetAmountOfTokenOk() (*Token, bool)`

GetAmountOfTokenOk returns a tuple with the AmountOfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfToken

`func (o *Value) SetAmountOfToken(v Token)`

SetAmountOfToken sets AmountOfToken field to given value.


### GetInAccount

`func (o *Value) GetInAccount() Party`

GetInAccount returns the InAccount field if non-nil, zero value otherwise.

### GetInAccountOk

`func (o *Value) GetInAccountOk() (*Party, bool)`

GetInAccountOk returns a tuple with the InAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAccount

`func (o *Value) SetInAccount(v Party)`

SetInAccount sets InAccount field to given value.


### GetNegate

`func (o *Value) GetNegate() Value`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *Value) GetNegateOk() (*Value, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *Value) SetNegate(v Value)`

SetNegate sets Negate field to given value.


### GetAdd

`func (o *Value) GetAdd() Value`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *Value) GetAddOk() (*Value, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *Value) SetAdd(v Value)`

SetAdd sets Add field to given value.


### GetAnd

`func (o *Value) GetAnd() Value`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *Value) GetAndOk() (*Value, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *Value) SetAnd(v Value)`

SetAnd sets And field to given value.


### GetMinus

`func (o *Value) GetMinus() Value`

GetMinus returns the Minus field if non-nil, zero value otherwise.

### GetMinusOk

`func (o *Value) GetMinusOk() (*Value, bool)`

GetMinusOk returns a tuple with the Minus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinus

`func (o *Value) SetMinus(v Value)`

SetMinus sets Minus field to given value.


### GetValue

`func (o *Value) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Value) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Value) SetValue(v Value)`

SetValue sets Value field to given value.


### GetMultiply

`func (o *Value) GetMultiply() Value`

GetMultiply returns the Multiply field if non-nil, zero value otherwise.

### GetMultiplyOk

`func (o *Value) GetMultiplyOk() (*Value, bool)`

GetMultiplyOk returns a tuple with the Multiply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiply

`func (o *Value) SetMultiply(v Value)`

SetMultiply sets Multiply field to given value.


### GetTimes

`func (o *Value) GetTimes() Value`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *Value) GetTimesOk() (*Value, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *Value) SetTimes(v Value)`

SetTimes sets Times field to given value.


### GetBy

`func (o *Value) GetBy() Value`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *Value) GetByOk() (*Value, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *Value) SetBy(v Value)`

SetBy sets By field to given value.


### GetDivide

`func (o *Value) GetDivide() Value`

GetDivide returns the Divide field if non-nil, zero value otherwise.

### GetDivideOk

`func (o *Value) GetDivideOk() (*Value, bool)`

GetDivideOk returns a tuple with the Divide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivide

`func (o *Value) SetDivide(v Value)`

SetDivide sets Divide field to given value.


### GetValueOfChoice

`func (o *Value) GetValueOfChoice() ChoiceId`

GetValueOfChoice returns the ValueOfChoice field if non-nil, zero value otherwise.

### GetValueOfChoiceOk

`func (o *Value) GetValueOfChoiceOk() (*ChoiceId, bool)`

GetValueOfChoiceOk returns a tuple with the ValueOfChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOfChoice

`func (o *Value) SetValueOfChoice(v ChoiceId)`

SetValueOfChoice sets ValueOfChoice field to given value.


### GetUseValue

`func (o *Value) GetUseValue() string`

GetUseValue returns the UseValue field if non-nil, zero value otherwise.

### GetUseValueOk

`func (o *Value) GetUseValueOk() (*string, bool)`

GetUseValueOk returns a tuple with the UseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValue

`func (o *Value) SetUseValue(v string)`

SetUseValue sets UseValue field to given value.


### GetElse

`func (o *Value) GetElse() Value`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *Value) GetElseOk() (*Value, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *Value) SetElse(v Value)`

SetElse sets Else field to given value.


### GetIf

`func (o *Value) GetIf() Observation`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *Value) GetIfOk() (*Observation, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *Value) SetIf(v Observation)`

SetIf sets If field to given value.


### GetThen

`func (o *Value) GetThen() Value`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *Value) GetThenOk() (*Value, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *Value) SetThen(v Value)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


