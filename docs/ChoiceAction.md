# ChoiceAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChooseBetween** | [**[]Bound**](Bound.md) |  | 
**ForChoice** | [**ChoiceId**](ChoiceId.md) |  | 

## Methods

### NewChoiceAction

`func NewChoiceAction(chooseBetween []Bound, forChoice ChoiceId, ) *ChoiceAction`

NewChoiceAction instantiates a new ChoiceAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceActionWithDefaults

`func NewChoiceActionWithDefaults() *ChoiceAction`

NewChoiceActionWithDefaults instantiates a new ChoiceAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChooseBetween

`func (o *ChoiceAction) GetChooseBetween() []Bound`

GetChooseBetween returns the ChooseBetween field if non-nil, zero value otherwise.

### GetChooseBetweenOk

`func (o *ChoiceAction) GetChooseBetweenOk() (*[]Bound, bool)`

GetChooseBetweenOk returns a tuple with the ChooseBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChooseBetween

`func (o *ChoiceAction) SetChooseBetween(v []Bound)`

SetChooseBetween sets ChooseBetween field to given value.


### GetForChoice

`func (o *ChoiceAction) GetForChoice() ChoiceId`

GetForChoice returns the ForChoice field if non-nil, zero value otherwise.

### GetForChoiceOk

`func (o *ChoiceAction) GetForChoiceOk() (*ChoiceId, bool)`

GetForChoiceOk returns a tuple with the ForChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChoice

`func (o *ChoiceAction) SetForChoice(v ChoiceId)`

SetForChoice sets ForChoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


