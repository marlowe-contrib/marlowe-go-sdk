# LabelledObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | An arbitrary text identifier for an object in a Marlowe object bundle. | 
**Type** | **string** |  | 
**Value** | [**LabelledObjectValue**](LabelledObjectValue.md) |  | 

## Methods

### NewLabelledObject

`func NewLabelledObject(label string, type_ string, value LabelledObjectValue, ) *LabelledObject`

NewLabelledObject instantiates a new LabelledObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelledObjectWithDefaults

`func NewLabelledObjectWithDefaults() *LabelledObject`

NewLabelledObjectWithDefaults instantiates a new LabelledObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LabelledObject) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LabelledObject) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LabelledObject) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *LabelledObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelledObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelledObject) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *LabelledObject) GetValue() LabelledObjectValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LabelledObject) GetValueOk() (*LabelledObjectValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LabelledObject) SetValue(v LabelledObjectValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


