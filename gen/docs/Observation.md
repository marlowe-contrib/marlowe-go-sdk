# Observation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | [**Observation**](Observation.md) |  | 
**Both** | [**Observation**](Observation.md) |  | 
**Either** | [**Observation**](Observation.md) |  | 
**Or** | [**Observation**](Observation.md) |  | 
**Not** | [**Observation**](Observation.md) |  | 
**ChoseSomethingFor** | [**ChoiceId**](ChoiceId.md) |  | 
**GeThan** | [**Value**](Value.md) |  | 
**Value** | [**Value**](Value.md) |  | 
**Gt** | [**Value**](Value.md) |  | 
**Lt** | [**Value**](Value.md) |  | 
**LeThan** | [**Value**](Value.md) |  | 
**EqualTo** | [**Value**](Value.md) |  | 

## Methods

### NewObservation

`func NewObservation(and Observation, both Observation, either Observation, or Observation, not Observation, choseSomethingFor ChoiceId, geThan Value, value Value, gt Value, lt Value, leThan Value, equalTo Value, ) *Observation`

NewObservation instantiates a new Observation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationWithDefaults

`func NewObservationWithDefaults() *Observation`

NewObservationWithDefaults instantiates a new Observation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *Observation) GetAnd() Observation`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *Observation) GetAndOk() (*Observation, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *Observation) SetAnd(v Observation)`

SetAnd sets And field to given value.


### GetBoth

`func (o *Observation) GetBoth() Observation`

GetBoth returns the Both field if non-nil, zero value otherwise.

### GetBothOk

`func (o *Observation) GetBothOk() (*Observation, bool)`

GetBothOk returns a tuple with the Both field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoth

`func (o *Observation) SetBoth(v Observation)`

SetBoth sets Both field to given value.


### GetEither

`func (o *Observation) GetEither() Observation`

GetEither returns the Either field if non-nil, zero value otherwise.

### GetEitherOk

`func (o *Observation) GetEitherOk() (*Observation, bool)`

GetEitherOk returns a tuple with the Either field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEither

`func (o *Observation) SetEither(v Observation)`

SetEither sets Either field to given value.


### GetOr

`func (o *Observation) GetOr() Observation`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *Observation) GetOrOk() (*Observation, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *Observation) SetOr(v Observation)`

SetOr sets Or field to given value.


### GetNot

`func (o *Observation) GetNot() Observation`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *Observation) GetNotOk() (*Observation, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *Observation) SetNot(v Observation)`

SetNot sets Not field to given value.


### GetChoseSomethingFor

`func (o *Observation) GetChoseSomethingFor() ChoiceId`

GetChoseSomethingFor returns the ChoseSomethingFor field if non-nil, zero value otherwise.

### GetChoseSomethingForOk

`func (o *Observation) GetChoseSomethingForOk() (*ChoiceId, bool)`

GetChoseSomethingForOk returns a tuple with the ChoseSomethingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoseSomethingFor

`func (o *Observation) SetChoseSomethingFor(v ChoiceId)`

SetChoseSomethingFor sets ChoseSomethingFor field to given value.


### GetGeThan

`func (o *Observation) GetGeThan() Value`

GetGeThan returns the GeThan field if non-nil, zero value otherwise.

### GetGeThanOk

`func (o *Observation) GetGeThanOk() (*Value, bool)`

GetGeThanOk returns a tuple with the GeThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeThan

`func (o *Observation) SetGeThan(v Value)`

SetGeThan sets GeThan field to given value.


### GetValue

`func (o *Observation) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Observation) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Observation) SetValue(v Value)`

SetValue sets Value field to given value.


### GetGt

`func (o *Observation) GetGt() Value`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *Observation) GetGtOk() (*Value, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *Observation) SetGt(v Value)`

SetGt sets Gt field to given value.


### GetLt

`func (o *Observation) GetLt() Value`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *Observation) GetLtOk() (*Value, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *Observation) SetLt(v Value)`

SetLt sets Lt field to given value.


### GetLeThan

`func (o *Observation) GetLeThan() Value`

GetLeThan returns the LeThan field if non-nil, zero value otherwise.

### GetLeThanOk

`func (o *Observation) GetLeThanOk() (*Value, bool)`

GetLeThanOk returns a tuple with the LeThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeThan

`func (o *Observation) SetLeThan(v Value)`

SetLeThan sets LeThan field to given value.


### GetEqualTo

`func (o *Observation) GetEqualTo() Value`

GetEqualTo returns the EqualTo field if non-nil, zero value otherwise.

### GetEqualToOk

`func (o *Observation) GetEqualToOk() (*Value, bool)`

GetEqualToOk returns a tuple with the EqualTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualTo

`func (o *Observation) SetEqualTo(v Value)`

SetEqualTo sets EqualTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


