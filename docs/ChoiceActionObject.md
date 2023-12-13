# ChoiceActionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChooseBetween** | [**[]Bound**](Bound.md) |  | 
**ForChoice** | [**ChoiceId**](ChoiceId.md) |  | 

## Methods

### NewChoiceActionObject

`func NewChoiceActionObject(chooseBetween []Bound, forChoice ChoiceId, ) *ChoiceActionObject`

NewChoiceActionObject instantiates a new ChoiceActionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceActionObjectWithDefaults

`func NewChoiceActionObjectWithDefaults() *ChoiceActionObject`

NewChoiceActionObjectWithDefaults instantiates a new ChoiceActionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChooseBetween

`func (o *ChoiceActionObject) GetChooseBetween() []Bound`

GetChooseBetween returns the ChooseBetween field if non-nil, zero value otherwise.

### GetChooseBetweenOk

`func (o *ChoiceActionObject) GetChooseBetweenOk() (*[]Bound, bool)`

GetChooseBetweenOk returns a tuple with the ChooseBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChooseBetween

`func (o *ChoiceActionObject) SetChooseBetween(v []Bound)`

SetChooseBetween sets ChooseBetween field to given value.


### GetForChoice

`func (o *ChoiceActionObject) GetForChoice() ChoiceId`

GetForChoice returns the ForChoice field if non-nil, zero value otherwise.

### GetForChoiceOk

`func (o *ChoiceActionObject) GetForChoiceOk() (*ChoiceId, bool)`

GetForChoiceOk returns a tuple with the ForChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChoice

`func (o *ChoiceActionObject) SetForChoice(v ChoiceId)`

SetForChoice sets ForChoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


