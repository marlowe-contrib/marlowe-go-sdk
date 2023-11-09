# LabelledObjectValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOfToken** | [**TokenObject**](TokenObject.md) |  | 
**InAccount** | [**PartyObject**](PartyObject.md) |  | 
**Negate** | [**ValueObject**](ValueObject.md) |  | 
**Add** | [**ValueObject**](ValueObject.md) |  | 
**And** | [**ObservationObject**](ObservationObject.md) |  | 
**Minus** | [**ValueObject**](ValueObject.md) |  | 
**Value** | [**ValueObject**](ValueObject.md) |  | 
**Multiply** | [**ValueObject**](ValueObject.md) |  | 
**Times** | [**ValueObject**](ValueObject.md) |  | 
**By** | [**ValueObject**](ValueObject.md) |  | 
**Divide** | [**ValueObject**](ValueObject.md) |  | 
**ValueOfChoice** | [**ChoiceIdObject**](ChoiceIdObject.md) |  | 
**UseValue** | **string** |  | 
**Else** | [**ContractObject**](ContractObject.md) |  | 
**If** | [**ObservationObject**](ObservationObject.md) |  | 
**Then** | [**ContractObject**](ContractObject.md) |  | 
**Ref** | **string** | An arbitrary text identifier for an object in a Marlowe object bundle. | 
**Both** | [**ObservationObject**](ObservationObject.md) |  | 
**Either** | [**ObservationObject**](ObservationObject.md) |  | 
**Or** | [**ObservationObject**](ObservationObject.md) |  | 
**Not** | [**ObservationObject**](ObservationObject.md) |  | 
**ChoseSomethingFor** | [**ChoiceIdObject**](ChoiceIdObject.md) |  | 
**GeThan** | [**ValueObject**](ValueObject.md) |  | 
**Gt** | [**ValueObject**](ValueObject.md) |  | 
**Lt** | [**ValueObject**](ValueObject.md) |  | 
**LeThan** | [**ValueObject**](ValueObject.md) |  | 
**EqualTo** | [**ValueObject**](ValueObject.md) |  | 
**FromAccount** | [**PartyObject**](PartyObject.md) |  | 
**Pay** | [**ValueObject**](ValueObject.md) |  | 
**To** | [**PayeeObject**](PayeeObject.md) |  | 
**Token** | [**TokenObject**](TokenObject.md) |  | 
**Timeout** | **int32** |  | 
**TimeoutContinuation** | [**ContractObject**](ContractObject.md) |  | 
**When** | [**[]CaseObject**](CaseObject.md) |  | 
**Be** | [**ValueObject**](ValueObject.md) |  | 
**Let** | **string** |  | 
**Assert** | [**ObservationObject**](ObservationObject.md) |  | 
**RoleToken** | **string** |  | 
**Address** | **string** | A cardano address | 
**CurrencySymbol** | **string** |  | 
**TokenName** | **string** |  | 
**Deposits** | [**Value**](Value.md) |  | 
**IntoAccount** | [**Party**](Party.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**Party** | [**Party**](Party.md) |  | 
**ChooseBetween** | [**[]Bound**](Bound.md) |  | 
**ForChoice** | [**ChoiceId**](ChoiceId.md) |  | 
**NotifyIf** | [**Observation**](Observation.md) |  | 

## Methods

### NewLabelledObjectValue

`func NewLabelledObjectValue(amountOfToken TokenObject, inAccount PartyObject, negate ValueObject, add ValueObject, and ObservationObject, minus ValueObject, value ValueObject, multiply ValueObject, times ValueObject, by ValueObject, divide ValueObject, valueOfChoice ChoiceIdObject, useValue string, else_ ContractObject, if_ ObservationObject, then ContractObject, ref string, both ObservationObject, either ObservationObject, or ObservationObject, not ObservationObject, choseSomethingFor ChoiceIdObject, geThan ValueObject, gt ValueObject, lt ValueObject, leThan ValueObject, equalTo ValueObject, fromAccount PartyObject, pay ValueObject, to PayeeObject, token TokenObject, timeout int32, timeoutContinuation ContractObject, when []CaseObject, be ValueObject, let string, assert ObservationObject, roleToken string, address string, currencySymbol string, tokenName string, deposits Value, intoAccount Party, ofToken Token, party Party, chooseBetween []Bound, forChoice ChoiceId, notifyIf Observation, ) *LabelledObjectValue`

NewLabelledObjectValue instantiates a new LabelledObjectValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelledObjectValueWithDefaults

`func NewLabelledObjectValueWithDefaults() *LabelledObjectValue`

NewLabelledObjectValueWithDefaults instantiates a new LabelledObjectValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOfToken

`func (o *LabelledObjectValue) GetAmountOfToken() TokenObject`

GetAmountOfToken returns the AmountOfToken field if non-nil, zero value otherwise.

### GetAmountOfTokenOk

`func (o *LabelledObjectValue) GetAmountOfTokenOk() (*TokenObject, bool)`

GetAmountOfTokenOk returns a tuple with the AmountOfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfToken

`func (o *LabelledObjectValue) SetAmountOfToken(v TokenObject)`

SetAmountOfToken sets AmountOfToken field to given value.


### GetInAccount

`func (o *LabelledObjectValue) GetInAccount() PartyObject`

GetInAccount returns the InAccount field if non-nil, zero value otherwise.

### GetInAccountOk

`func (o *LabelledObjectValue) GetInAccountOk() (*PartyObject, bool)`

GetInAccountOk returns a tuple with the InAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAccount

`func (o *LabelledObjectValue) SetInAccount(v PartyObject)`

SetInAccount sets InAccount field to given value.


### GetNegate

`func (o *LabelledObjectValue) GetNegate() ValueObject`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *LabelledObjectValue) GetNegateOk() (*ValueObject, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *LabelledObjectValue) SetNegate(v ValueObject)`

SetNegate sets Negate field to given value.


### GetAdd

`func (o *LabelledObjectValue) GetAdd() ValueObject`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *LabelledObjectValue) GetAddOk() (*ValueObject, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *LabelledObjectValue) SetAdd(v ValueObject)`

SetAdd sets Add field to given value.


### GetAnd

`func (o *LabelledObjectValue) GetAnd() ObservationObject`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *LabelledObjectValue) GetAndOk() (*ObservationObject, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *LabelledObjectValue) SetAnd(v ObservationObject)`

SetAnd sets And field to given value.


### GetMinus

`func (o *LabelledObjectValue) GetMinus() ValueObject`

GetMinus returns the Minus field if non-nil, zero value otherwise.

### GetMinusOk

`func (o *LabelledObjectValue) GetMinusOk() (*ValueObject, bool)`

GetMinusOk returns a tuple with the Minus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinus

`func (o *LabelledObjectValue) SetMinus(v ValueObject)`

SetMinus sets Minus field to given value.


### GetValue

`func (o *LabelledObjectValue) GetValue() ValueObject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LabelledObjectValue) GetValueOk() (*ValueObject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LabelledObjectValue) SetValue(v ValueObject)`

SetValue sets Value field to given value.


### GetMultiply

`func (o *LabelledObjectValue) GetMultiply() ValueObject`

GetMultiply returns the Multiply field if non-nil, zero value otherwise.

### GetMultiplyOk

`func (o *LabelledObjectValue) GetMultiplyOk() (*ValueObject, bool)`

GetMultiplyOk returns a tuple with the Multiply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiply

`func (o *LabelledObjectValue) SetMultiply(v ValueObject)`

SetMultiply sets Multiply field to given value.


### GetTimes

`func (o *LabelledObjectValue) GetTimes() ValueObject`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *LabelledObjectValue) GetTimesOk() (*ValueObject, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *LabelledObjectValue) SetTimes(v ValueObject)`

SetTimes sets Times field to given value.


### GetBy

`func (o *LabelledObjectValue) GetBy() ValueObject`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *LabelledObjectValue) GetByOk() (*ValueObject, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *LabelledObjectValue) SetBy(v ValueObject)`

SetBy sets By field to given value.


### GetDivide

`func (o *LabelledObjectValue) GetDivide() ValueObject`

GetDivide returns the Divide field if non-nil, zero value otherwise.

### GetDivideOk

`func (o *LabelledObjectValue) GetDivideOk() (*ValueObject, bool)`

GetDivideOk returns a tuple with the Divide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivide

`func (o *LabelledObjectValue) SetDivide(v ValueObject)`

SetDivide sets Divide field to given value.


### GetValueOfChoice

`func (o *LabelledObjectValue) GetValueOfChoice() ChoiceIdObject`

GetValueOfChoice returns the ValueOfChoice field if non-nil, zero value otherwise.

### GetValueOfChoiceOk

`func (o *LabelledObjectValue) GetValueOfChoiceOk() (*ChoiceIdObject, bool)`

GetValueOfChoiceOk returns a tuple with the ValueOfChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOfChoice

`func (o *LabelledObjectValue) SetValueOfChoice(v ChoiceIdObject)`

SetValueOfChoice sets ValueOfChoice field to given value.


### GetUseValue

`func (o *LabelledObjectValue) GetUseValue() string`

GetUseValue returns the UseValue field if non-nil, zero value otherwise.

### GetUseValueOk

`func (o *LabelledObjectValue) GetUseValueOk() (*string, bool)`

GetUseValueOk returns a tuple with the UseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValue

`func (o *LabelledObjectValue) SetUseValue(v string)`

SetUseValue sets UseValue field to given value.


### GetElse

`func (o *LabelledObjectValue) GetElse() ContractObject`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *LabelledObjectValue) GetElseOk() (*ContractObject, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *LabelledObjectValue) SetElse(v ContractObject)`

SetElse sets Else field to given value.


### GetIf

`func (o *LabelledObjectValue) GetIf() ObservationObject`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *LabelledObjectValue) GetIfOk() (*ObservationObject, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *LabelledObjectValue) SetIf(v ObservationObject)`

SetIf sets If field to given value.


### GetThen

`func (o *LabelledObjectValue) GetThen() ContractObject`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *LabelledObjectValue) GetThenOk() (*ContractObject, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *LabelledObjectValue) SetThen(v ContractObject)`

SetThen sets Then field to given value.


### GetRef

`func (o *LabelledObjectValue) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *LabelledObjectValue) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *LabelledObjectValue) SetRef(v string)`

SetRef sets Ref field to given value.


### GetBoth

`func (o *LabelledObjectValue) GetBoth() ObservationObject`

GetBoth returns the Both field if non-nil, zero value otherwise.

### GetBothOk

`func (o *LabelledObjectValue) GetBothOk() (*ObservationObject, bool)`

GetBothOk returns a tuple with the Both field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoth

`func (o *LabelledObjectValue) SetBoth(v ObservationObject)`

SetBoth sets Both field to given value.


### GetEither

`func (o *LabelledObjectValue) GetEither() ObservationObject`

GetEither returns the Either field if non-nil, zero value otherwise.

### GetEitherOk

`func (o *LabelledObjectValue) GetEitherOk() (*ObservationObject, bool)`

GetEitherOk returns a tuple with the Either field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEither

`func (o *LabelledObjectValue) SetEither(v ObservationObject)`

SetEither sets Either field to given value.


### GetOr

`func (o *LabelledObjectValue) GetOr() ObservationObject`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *LabelledObjectValue) GetOrOk() (*ObservationObject, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *LabelledObjectValue) SetOr(v ObservationObject)`

SetOr sets Or field to given value.


### GetNot

`func (o *LabelledObjectValue) GetNot() ObservationObject`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *LabelledObjectValue) GetNotOk() (*ObservationObject, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *LabelledObjectValue) SetNot(v ObservationObject)`

SetNot sets Not field to given value.


### GetChoseSomethingFor

`func (o *LabelledObjectValue) GetChoseSomethingFor() ChoiceIdObject`

GetChoseSomethingFor returns the ChoseSomethingFor field if non-nil, zero value otherwise.

### GetChoseSomethingForOk

`func (o *LabelledObjectValue) GetChoseSomethingForOk() (*ChoiceIdObject, bool)`

GetChoseSomethingForOk returns a tuple with the ChoseSomethingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoseSomethingFor

`func (o *LabelledObjectValue) SetChoseSomethingFor(v ChoiceIdObject)`

SetChoseSomethingFor sets ChoseSomethingFor field to given value.


### GetGeThan

`func (o *LabelledObjectValue) GetGeThan() ValueObject`

GetGeThan returns the GeThan field if non-nil, zero value otherwise.

### GetGeThanOk

`func (o *LabelledObjectValue) GetGeThanOk() (*ValueObject, bool)`

GetGeThanOk returns a tuple with the GeThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeThan

`func (o *LabelledObjectValue) SetGeThan(v ValueObject)`

SetGeThan sets GeThan field to given value.


### GetGt

`func (o *LabelledObjectValue) GetGt() ValueObject`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *LabelledObjectValue) GetGtOk() (*ValueObject, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *LabelledObjectValue) SetGt(v ValueObject)`

SetGt sets Gt field to given value.


### GetLt

`func (o *LabelledObjectValue) GetLt() ValueObject`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *LabelledObjectValue) GetLtOk() (*ValueObject, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *LabelledObjectValue) SetLt(v ValueObject)`

SetLt sets Lt field to given value.


### GetLeThan

`func (o *LabelledObjectValue) GetLeThan() ValueObject`

GetLeThan returns the LeThan field if non-nil, zero value otherwise.

### GetLeThanOk

`func (o *LabelledObjectValue) GetLeThanOk() (*ValueObject, bool)`

GetLeThanOk returns a tuple with the LeThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeThan

`func (o *LabelledObjectValue) SetLeThan(v ValueObject)`

SetLeThan sets LeThan field to given value.


### GetEqualTo

`func (o *LabelledObjectValue) GetEqualTo() ValueObject`

GetEqualTo returns the EqualTo field if non-nil, zero value otherwise.

### GetEqualToOk

`func (o *LabelledObjectValue) GetEqualToOk() (*ValueObject, bool)`

GetEqualToOk returns a tuple with the EqualTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualTo

`func (o *LabelledObjectValue) SetEqualTo(v ValueObject)`

SetEqualTo sets EqualTo field to given value.


### GetFromAccount

`func (o *LabelledObjectValue) GetFromAccount() PartyObject`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *LabelledObjectValue) GetFromAccountOk() (*PartyObject, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *LabelledObjectValue) SetFromAccount(v PartyObject)`

SetFromAccount sets FromAccount field to given value.


### GetPay

`func (o *LabelledObjectValue) GetPay() ValueObject`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *LabelledObjectValue) GetPayOk() (*ValueObject, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *LabelledObjectValue) SetPay(v ValueObject)`

SetPay sets Pay field to given value.


### GetTo

`func (o *LabelledObjectValue) GetTo() PayeeObject`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LabelledObjectValue) GetToOk() (*PayeeObject, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LabelledObjectValue) SetTo(v PayeeObject)`

SetTo sets To field to given value.


### GetToken

`func (o *LabelledObjectValue) GetToken() TokenObject`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LabelledObjectValue) GetTokenOk() (*TokenObject, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LabelledObjectValue) SetToken(v TokenObject)`

SetToken sets Token field to given value.


### GetTimeout

`func (o *LabelledObjectValue) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *LabelledObjectValue) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *LabelledObjectValue) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetTimeoutContinuation

`func (o *LabelledObjectValue) GetTimeoutContinuation() ContractObject`

GetTimeoutContinuation returns the TimeoutContinuation field if non-nil, zero value otherwise.

### GetTimeoutContinuationOk

`func (o *LabelledObjectValue) GetTimeoutContinuationOk() (*ContractObject, bool)`

GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutContinuation

`func (o *LabelledObjectValue) SetTimeoutContinuation(v ContractObject)`

SetTimeoutContinuation sets TimeoutContinuation field to given value.


### GetWhen

`func (o *LabelledObjectValue) GetWhen() []CaseObject`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *LabelledObjectValue) GetWhenOk() (*[]CaseObject, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *LabelledObjectValue) SetWhen(v []CaseObject)`

SetWhen sets When field to given value.


### GetBe

`func (o *LabelledObjectValue) GetBe() ValueObject`

GetBe returns the Be field if non-nil, zero value otherwise.

### GetBeOk

`func (o *LabelledObjectValue) GetBeOk() (*ValueObject, bool)`

GetBeOk returns a tuple with the Be field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBe

`func (o *LabelledObjectValue) SetBe(v ValueObject)`

SetBe sets Be field to given value.


### GetLet

`func (o *LabelledObjectValue) GetLet() string`

GetLet returns the Let field if non-nil, zero value otherwise.

### GetLetOk

`func (o *LabelledObjectValue) GetLetOk() (*string, bool)`

GetLetOk returns a tuple with the Let field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLet

`func (o *LabelledObjectValue) SetLet(v string)`

SetLet sets Let field to given value.


### GetAssert

`func (o *LabelledObjectValue) GetAssert() ObservationObject`

GetAssert returns the Assert field if non-nil, zero value otherwise.

### GetAssertOk

`func (o *LabelledObjectValue) GetAssertOk() (*ObservationObject, bool)`

GetAssertOk returns a tuple with the Assert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssert

`func (o *LabelledObjectValue) SetAssert(v ObservationObject)`

SetAssert sets Assert field to given value.


### GetRoleToken

`func (o *LabelledObjectValue) GetRoleToken() string`

GetRoleToken returns the RoleToken field if non-nil, zero value otherwise.

### GetRoleTokenOk

`func (o *LabelledObjectValue) GetRoleTokenOk() (*string, bool)`

GetRoleTokenOk returns a tuple with the RoleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleToken

`func (o *LabelledObjectValue) SetRoleToken(v string)`

SetRoleToken sets RoleToken field to given value.


### GetAddress

`func (o *LabelledObjectValue) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LabelledObjectValue) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LabelledObjectValue) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCurrencySymbol

`func (o *LabelledObjectValue) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *LabelledObjectValue) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *LabelledObjectValue) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.


### GetTokenName

`func (o *LabelledObjectValue) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *LabelledObjectValue) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *LabelledObjectValue) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetDeposits

`func (o *LabelledObjectValue) GetDeposits() Value`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *LabelledObjectValue) GetDepositsOk() (*Value, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *LabelledObjectValue) SetDeposits(v Value)`

SetDeposits sets Deposits field to given value.


### GetIntoAccount

`func (o *LabelledObjectValue) GetIntoAccount() Party`

GetIntoAccount returns the IntoAccount field if non-nil, zero value otherwise.

### GetIntoAccountOk

`func (o *LabelledObjectValue) GetIntoAccountOk() (*Party, bool)`

GetIntoAccountOk returns a tuple with the IntoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntoAccount

`func (o *LabelledObjectValue) SetIntoAccount(v Party)`

SetIntoAccount sets IntoAccount field to given value.


### GetOfToken

`func (o *LabelledObjectValue) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *LabelledObjectValue) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *LabelledObjectValue) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetParty

`func (o *LabelledObjectValue) GetParty() Party`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *LabelledObjectValue) GetPartyOk() (*Party, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *LabelledObjectValue) SetParty(v Party)`

SetParty sets Party field to given value.


### GetChooseBetween

`func (o *LabelledObjectValue) GetChooseBetween() []Bound`

GetChooseBetween returns the ChooseBetween field if non-nil, zero value otherwise.

### GetChooseBetweenOk

`func (o *LabelledObjectValue) GetChooseBetweenOk() (*[]Bound, bool)`

GetChooseBetweenOk returns a tuple with the ChooseBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChooseBetween

`func (o *LabelledObjectValue) SetChooseBetween(v []Bound)`

SetChooseBetween sets ChooseBetween field to given value.


### GetForChoice

`func (o *LabelledObjectValue) GetForChoice() ChoiceId`

GetForChoice returns the ForChoice field if non-nil, zero value otherwise.

### GetForChoiceOk

`func (o *LabelledObjectValue) GetForChoiceOk() (*ChoiceId, bool)`

GetForChoiceOk returns a tuple with the ForChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChoice

`func (o *LabelledObjectValue) SetForChoice(v ChoiceId)`

SetForChoice sets ForChoice field to given value.


### GetNotifyIf

`func (o *LabelledObjectValue) GetNotifyIf() Observation`

GetNotifyIf returns the NotifyIf field if non-nil, zero value otherwise.

### GetNotifyIfOk

`func (o *LabelledObjectValue) GetNotifyIfOk() (*Observation, bool)`

GetNotifyIfOk returns a tuple with the NotifyIf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyIf

`func (o *LabelledObjectValue) SetNotifyIf(v Observation)`

SetNotifyIf sets NotifyIf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


