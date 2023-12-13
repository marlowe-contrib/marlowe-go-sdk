# ObservationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | [**ObservationObject**](ObservationObject.md) |  | 
**Both** | [**ObservationObject**](ObservationObject.md) |  | 
**Either** | [**ObservationObject**](ObservationObject.md) |  | 
**Or** | [**ObservationObject**](ObservationObject.md) |  | 
**Not** | [**ObservationObject**](ObservationObject.md) |  | 
**ChoseSomethingFor** | [**ChoiceIdObject**](ChoiceIdObject.md) |  | 
**GeThan** | [**ValueObject**](ValueObject.md) |  | 
**Value** | [**ValueObject**](ValueObject.md) |  | 
**Gt** | [**ValueObject**](ValueObject.md) |  | 
**Lt** | [**ValueObject**](ValueObject.md) |  | 
**LeThan** | [**ValueObject**](ValueObject.md) |  | 
**EqualTo** | [**ValueObject**](ValueObject.md) |  | 
**Ref** | **string** | An arbitrary text identifier for an object in a Marlowe object bundle. | 

## Methods

### NewObservationObject

`func NewObservationObject(and ObservationObject, both ObservationObject, either ObservationObject, or ObservationObject, not ObservationObject, choseSomethingFor ChoiceIdObject, geThan ValueObject, value ValueObject, gt ValueObject, lt ValueObject, leThan ValueObject, equalTo ValueObject, ref string, ) *ObservationObject`

NewObservationObject instantiates a new ObservationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationObjectWithDefaults

`func NewObservationObjectWithDefaults() *ObservationObject`

NewObservationObjectWithDefaults instantiates a new ObservationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *ObservationObject) GetAnd() ObservationObject`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ObservationObject) GetAndOk() (*ObservationObject, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ObservationObject) SetAnd(v ObservationObject)`

SetAnd sets And field to given value.


### GetBoth

`func (o *ObservationObject) GetBoth() ObservationObject`

GetBoth returns the Both field if non-nil, zero value otherwise.

### GetBothOk

`func (o *ObservationObject) GetBothOk() (*ObservationObject, bool)`

GetBothOk returns a tuple with the Both field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoth

`func (o *ObservationObject) SetBoth(v ObservationObject)`

SetBoth sets Both field to given value.


### GetEither

`func (o *ObservationObject) GetEither() ObservationObject`

GetEither returns the Either field if non-nil, zero value otherwise.

### GetEitherOk

`func (o *ObservationObject) GetEitherOk() (*ObservationObject, bool)`

GetEitherOk returns a tuple with the Either field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEither

`func (o *ObservationObject) SetEither(v ObservationObject)`

SetEither sets Either field to given value.


### GetOr

`func (o *ObservationObject) GetOr() ObservationObject`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *ObservationObject) GetOrOk() (*ObservationObject, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *ObservationObject) SetOr(v ObservationObject)`

SetOr sets Or field to given value.


### GetNot

`func (o *ObservationObject) GetNot() ObservationObject`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *ObservationObject) GetNotOk() (*ObservationObject, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *ObservationObject) SetNot(v ObservationObject)`

SetNot sets Not field to given value.


### GetChoseSomethingFor

`func (o *ObservationObject) GetChoseSomethingFor() ChoiceIdObject`

GetChoseSomethingFor returns the ChoseSomethingFor field if non-nil, zero value otherwise.

### GetChoseSomethingForOk

`func (o *ObservationObject) GetChoseSomethingForOk() (*ChoiceIdObject, bool)`

GetChoseSomethingForOk returns a tuple with the ChoseSomethingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoseSomethingFor

`func (o *ObservationObject) SetChoseSomethingFor(v ChoiceIdObject)`

SetChoseSomethingFor sets ChoseSomethingFor field to given value.


### GetGeThan

`func (o *ObservationObject) GetGeThan() ValueObject`

GetGeThan returns the GeThan field if non-nil, zero value otherwise.

### GetGeThanOk

`func (o *ObservationObject) GetGeThanOk() (*ValueObject, bool)`

GetGeThanOk returns a tuple with the GeThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeThan

`func (o *ObservationObject) SetGeThan(v ValueObject)`

SetGeThan sets GeThan field to given value.


### GetValue

`func (o *ObservationObject) GetValue() ValueObject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ObservationObject) GetValueOk() (*ValueObject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ObservationObject) SetValue(v ValueObject)`

SetValue sets Value field to given value.


### GetGt

`func (o *ObservationObject) GetGt() ValueObject`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *ObservationObject) GetGtOk() (*ValueObject, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *ObservationObject) SetGt(v ValueObject)`

SetGt sets Gt field to given value.


### GetLt

`func (o *ObservationObject) GetLt() ValueObject`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *ObservationObject) GetLtOk() (*ValueObject, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *ObservationObject) SetLt(v ValueObject)`

SetLt sets Lt field to given value.


### GetLeThan

`func (o *ObservationObject) GetLeThan() ValueObject`

GetLeThan returns the LeThan field if non-nil, zero value otherwise.

### GetLeThanOk

`func (o *ObservationObject) GetLeThanOk() (*ValueObject, bool)`

GetLeThanOk returns a tuple with the LeThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeThan

`func (o *ObservationObject) SetLeThan(v ValueObject)`

SetLeThan sets LeThan field to given value.


### GetEqualTo

`func (o *ObservationObject) GetEqualTo() ValueObject`

GetEqualTo returns the EqualTo field if non-nil, zero value otherwise.

### GetEqualToOk

`func (o *ObservationObject) GetEqualToOk() (*ValueObject, bool)`

GetEqualToOk returns a tuple with the EqualTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualTo

`func (o *ObservationObject) SetEqualTo(v ValueObject)`

SetEqualTo sets EqualTo field to given value.


### GetRef

`func (o *ObservationObject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ObservationObject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ObservationObject) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


