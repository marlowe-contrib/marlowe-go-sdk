# CanNotify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseIndex** | **int32** | Index of a \&quot;Case Action\&quot; in a \&quot;When\&quot; | 
**IsMerkleizedContinuation** | **bool** | Indicates if a given contract continuation is merkleized | 

## Methods

### NewCanNotify

`func NewCanNotify(caseIndex int32, isMerkleizedContinuation bool, ) *CanNotify`

NewCanNotify instantiates a new CanNotify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanNotifyWithDefaults

`func NewCanNotifyWithDefaults() *CanNotify`

NewCanNotifyWithDefaults instantiates a new CanNotify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseIndex

`func (o *CanNotify) GetCaseIndex() int32`

GetCaseIndex returns the CaseIndex field if non-nil, zero value otherwise.

### GetCaseIndexOk

`func (o *CanNotify) GetCaseIndexOk() (*int32, bool)`

GetCaseIndexOk returns a tuple with the CaseIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseIndex

`func (o *CanNotify) SetCaseIndex(v int32)`

SetCaseIndex sets CaseIndex field to given value.


### GetIsMerkleizedContinuation

`func (o *CanNotify) GetIsMerkleizedContinuation() bool`

GetIsMerkleizedContinuation returns the IsMerkleizedContinuation field if non-nil, zero value otherwise.

### GetIsMerkleizedContinuationOk

`func (o *CanNotify) GetIsMerkleizedContinuationOk() (*bool, bool)`

GetIsMerkleizedContinuationOk returns a tuple with the IsMerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerkleizedContinuation

`func (o *CanNotify) SetIsMerkleizedContinuation(v bool)`

SetIsMerkleizedContinuation sets IsMerkleizedContinuation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


