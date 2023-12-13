# ChoiceIdObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChoiceName** | **string** |  | 
**ChoiceOwner** | [**PartyObject**](PartyObject.md) |  | 

## Methods

### NewChoiceIdObject

`func NewChoiceIdObject(choiceName string, choiceOwner PartyObject, ) *ChoiceIdObject`

NewChoiceIdObject instantiates a new ChoiceIdObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceIdObjectWithDefaults

`func NewChoiceIdObjectWithDefaults() *ChoiceIdObject`

NewChoiceIdObjectWithDefaults instantiates a new ChoiceIdObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoiceName

`func (o *ChoiceIdObject) GetChoiceName() string`

GetChoiceName returns the ChoiceName field if non-nil, zero value otherwise.

### GetChoiceNameOk

`func (o *ChoiceIdObject) GetChoiceNameOk() (*string, bool)`

GetChoiceNameOk returns a tuple with the ChoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceName

`func (o *ChoiceIdObject) SetChoiceName(v string)`

SetChoiceName sets ChoiceName field to given value.


### GetChoiceOwner

`func (o *ChoiceIdObject) GetChoiceOwner() PartyObject`

GetChoiceOwner returns the ChoiceOwner field if non-nil, zero value otherwise.

### GetChoiceOwnerOk

`func (o *ChoiceIdObject) GetChoiceOwnerOk() (*PartyObject, bool)`

GetChoiceOwnerOk returns a tuple with the ChoiceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceOwner

`func (o *ChoiceIdObject) SetChoiceOwner(v PartyObject)`

SetChoiceOwner sets ChoiceOwner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


