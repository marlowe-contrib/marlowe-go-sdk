# ValueObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOfToken** | [**TokenObject**](TokenObject.md) |  | 
**InAccount** | [**PartyObject**](PartyObject.md) |  | 
**Negate** | [**ValueObject**](ValueObject.md) |  | 
**Add** | [**ValueObject**](ValueObject.md) |  | 
**And** | [**ValueObject**](ValueObject.md) |  | 
**Minus** | [**ValueObject**](ValueObject.md) |  | 
**Value** | [**ValueObject**](ValueObject.md) |  | 
**Multiply** | [**ValueObject**](ValueObject.md) |  | 
**Times** | [**ValueObject**](ValueObject.md) |  | 
**By** | [**ValueObject**](ValueObject.md) |  | 
**Divide** | [**ValueObject**](ValueObject.md) |  | 
**ValueOfChoice** | [**ChoiceIdObject**](ChoiceIdObject.md) |  | 
**UseValue** | **string** |  | 
**Else** | [**ValueObject**](ValueObject.md) |  | 
**If** | [**ObservationObject**](ObservationObject.md) |  | 
**Then** | [**ValueObject**](ValueObject.md) |  | 
**Ref** | **string** | An arbitrary text identifier for an object in a Marlowe object bundle. | 

## Methods

### NewValueObject

`func NewValueObject(amountOfToken TokenObject, inAccount PartyObject, negate ValueObject, add ValueObject, and ValueObject, minus ValueObject, value ValueObject, multiply ValueObject, times ValueObject, by ValueObject, divide ValueObject, valueOfChoice ChoiceIdObject, useValue string, else_ ValueObject, if_ ObservationObject, then ValueObject, ref string, ) *ValueObject`

NewValueObject instantiates a new ValueObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueObjectWithDefaults

`func NewValueObjectWithDefaults() *ValueObject`

NewValueObjectWithDefaults instantiates a new ValueObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOfToken

`func (o *ValueObject) GetAmountOfToken() TokenObject`

GetAmountOfToken returns the AmountOfToken field if non-nil, zero value otherwise.

### GetAmountOfTokenOk

`func (o *ValueObject) GetAmountOfTokenOk() (*TokenObject, bool)`

GetAmountOfTokenOk returns a tuple with the AmountOfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfToken

`func (o *ValueObject) SetAmountOfToken(v TokenObject)`

SetAmountOfToken sets AmountOfToken field to given value.


### GetInAccount

`func (o *ValueObject) GetInAccount() PartyObject`

GetInAccount returns the InAccount field if non-nil, zero value otherwise.

### GetInAccountOk

`func (o *ValueObject) GetInAccountOk() (*PartyObject, bool)`

GetInAccountOk returns a tuple with the InAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAccount

`func (o *ValueObject) SetInAccount(v PartyObject)`

SetInAccount sets InAccount field to given value.


### GetNegate

`func (o *ValueObject) GetNegate() ValueObject`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *ValueObject) GetNegateOk() (*ValueObject, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *ValueObject) SetNegate(v ValueObject)`

SetNegate sets Negate field to given value.


### GetAdd

`func (o *ValueObject) GetAdd() ValueObject`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *ValueObject) GetAddOk() (*ValueObject, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *ValueObject) SetAdd(v ValueObject)`

SetAdd sets Add field to given value.


### GetAnd

`func (o *ValueObject) GetAnd() ValueObject`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ValueObject) GetAndOk() (*ValueObject, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ValueObject) SetAnd(v ValueObject)`

SetAnd sets And field to given value.


### GetMinus

`func (o *ValueObject) GetMinus() ValueObject`

GetMinus returns the Minus field if non-nil, zero value otherwise.

### GetMinusOk

`func (o *ValueObject) GetMinusOk() (*ValueObject, bool)`

GetMinusOk returns a tuple with the Minus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinus

`func (o *ValueObject) SetMinus(v ValueObject)`

SetMinus sets Minus field to given value.


### GetValue

`func (o *ValueObject) GetValue() ValueObject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValueObject) GetValueOk() (*ValueObject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValueObject) SetValue(v ValueObject)`

SetValue sets Value field to given value.


### GetMultiply

`func (o *ValueObject) GetMultiply() ValueObject`

GetMultiply returns the Multiply field if non-nil, zero value otherwise.

### GetMultiplyOk

`func (o *ValueObject) GetMultiplyOk() (*ValueObject, bool)`

GetMultiplyOk returns a tuple with the Multiply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiply

`func (o *ValueObject) SetMultiply(v ValueObject)`

SetMultiply sets Multiply field to given value.


### GetTimes

`func (o *ValueObject) GetTimes() ValueObject`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *ValueObject) GetTimesOk() (*ValueObject, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *ValueObject) SetTimes(v ValueObject)`

SetTimes sets Times field to given value.


### GetBy

`func (o *ValueObject) GetBy() ValueObject`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *ValueObject) GetByOk() (*ValueObject, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *ValueObject) SetBy(v ValueObject)`

SetBy sets By field to given value.


### GetDivide

`func (o *ValueObject) GetDivide() ValueObject`

GetDivide returns the Divide field if non-nil, zero value otherwise.

### GetDivideOk

`func (o *ValueObject) GetDivideOk() (*ValueObject, bool)`

GetDivideOk returns a tuple with the Divide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivide

`func (o *ValueObject) SetDivide(v ValueObject)`

SetDivide sets Divide field to given value.


### GetValueOfChoice

`func (o *ValueObject) GetValueOfChoice() ChoiceIdObject`

GetValueOfChoice returns the ValueOfChoice field if non-nil, zero value otherwise.

### GetValueOfChoiceOk

`func (o *ValueObject) GetValueOfChoiceOk() (*ChoiceIdObject, bool)`

GetValueOfChoiceOk returns a tuple with the ValueOfChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOfChoice

`func (o *ValueObject) SetValueOfChoice(v ChoiceIdObject)`

SetValueOfChoice sets ValueOfChoice field to given value.


### GetUseValue

`func (o *ValueObject) GetUseValue() string`

GetUseValue returns the UseValue field if non-nil, zero value otherwise.

### GetUseValueOk

`func (o *ValueObject) GetUseValueOk() (*string, bool)`

GetUseValueOk returns a tuple with the UseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValue

`func (o *ValueObject) SetUseValue(v string)`

SetUseValue sets UseValue field to given value.


### GetElse

`func (o *ValueObject) GetElse() ValueObject`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *ValueObject) GetElseOk() (*ValueObject, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *ValueObject) SetElse(v ValueObject)`

SetElse sets Else field to given value.


### GetIf

`func (o *ValueObject) GetIf() ObservationObject`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *ValueObject) GetIfOk() (*ObservationObject, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *ValueObject) SetIf(v ObservationObject)`

SetIf sets If field to given value.


### GetThen

`func (o *ValueObject) GetThen() ValueObject`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *ValueObject) GetThenOk() (*ValueObject, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *ValueObject) SetThen(v ValueObject)`

SetThen sets Then field to given value.


### GetRef

`func (o *ValueObject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ValueObject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ValueObject) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


