# ChoiceId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChoiceName** | **string** |  | 
**ChoiceOwner** | [**Party**](Party.md) |  | 

## Methods

### NewChoiceId

`func NewChoiceId(choiceName string, choiceOwner Party, ) *ChoiceId`

NewChoiceId instantiates a new ChoiceId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceIdWithDefaults

`func NewChoiceIdWithDefaults() *ChoiceId`

NewChoiceIdWithDefaults instantiates a new ChoiceId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoiceName

`func (o *ChoiceId) GetChoiceName() string`

GetChoiceName returns the ChoiceName field if non-nil, zero value otherwise.

### GetChoiceNameOk

`func (o *ChoiceId) GetChoiceNameOk() (*string, bool)`

GetChoiceNameOk returns a tuple with the ChoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceName

`func (o *ChoiceId) SetChoiceName(v string)`

SetChoiceName sets ChoiceName field to given value.


### GetChoiceOwner

`func (o *ChoiceId) GetChoiceOwner() Party`

GetChoiceOwner returns the ChoiceOwner field if non-nil, zero value otherwise.

### GetChoiceOwnerOk

`func (o *ChoiceId) GetChoiceOwnerOk() (*Party, bool)`

GetChoiceOwnerOk returns a tuple with the ChoiceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceOwner

`func (o *ChoiceId) SetChoiceOwner(v Party)`

SetChoiceOwner sets ChoiceOwner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


