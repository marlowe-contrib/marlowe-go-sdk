# CanChoose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanChooseBetween** | [**[]Bound**](Bound.md) |  | 
**CaseIndex** | **int32** | Index of a \&quot;Case Action\&quot; in a \&quot;When\&quot; | 
**ForChoice** | [**ChoiceId**](ChoiceId.md) |  | 
**IsMerkleizedContinuation** | **bool** | Indicates if a given contract continuation is merkleized | 

## Methods

### NewCanChoose

`func NewCanChoose(canChooseBetween []Bound, caseIndex int32, forChoice ChoiceId, isMerkleizedContinuation bool, ) *CanChoose`

NewCanChoose instantiates a new CanChoose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanChooseWithDefaults

`func NewCanChooseWithDefaults() *CanChoose`

NewCanChooseWithDefaults instantiates a new CanChoose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanChooseBetween

`func (o *CanChoose) GetCanChooseBetween() []Bound`

GetCanChooseBetween returns the CanChooseBetween field if non-nil, zero value otherwise.

### GetCanChooseBetweenOk

`func (o *CanChoose) GetCanChooseBetweenOk() (*[]Bound, bool)`

GetCanChooseBetweenOk returns a tuple with the CanChooseBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChooseBetween

`func (o *CanChoose) SetCanChooseBetween(v []Bound)`

SetCanChooseBetween sets CanChooseBetween field to given value.


### GetCaseIndex

`func (o *CanChoose) GetCaseIndex() int32`

GetCaseIndex returns the CaseIndex field if non-nil, zero value otherwise.

### GetCaseIndexOk

`func (o *CanChoose) GetCaseIndexOk() (*int32, bool)`

GetCaseIndexOk returns a tuple with the CaseIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseIndex

`func (o *CanChoose) SetCaseIndex(v int32)`

SetCaseIndex sets CaseIndex field to given value.


### GetForChoice

`func (o *CanChoose) GetForChoice() ChoiceId`

GetForChoice returns the ForChoice field if non-nil, zero value otherwise.

### GetForChoiceOk

`func (o *CanChoose) GetForChoiceOk() (*ChoiceId, bool)`

GetForChoiceOk returns a tuple with the ForChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChoice

`func (o *CanChoose) SetForChoice(v ChoiceId)`

SetForChoice sets ForChoice field to given value.


### GetIsMerkleizedContinuation

`func (o *CanChoose) GetIsMerkleizedContinuation() bool`

GetIsMerkleizedContinuation returns the IsMerkleizedContinuation field if non-nil, zero value otherwise.

### GetIsMerkleizedContinuationOk

`func (o *CanChoose) GetIsMerkleizedContinuationOk() (*bool, bool)`

GetIsMerkleizedContinuationOk returns a tuple with the IsMerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerkleizedContinuation

`func (o *CanChoose) SetIsMerkleizedContinuation(v bool)`

SetIsMerkleizedContinuation sets IsMerkleizedContinuation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


