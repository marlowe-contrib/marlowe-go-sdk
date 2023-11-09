# MarloweState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[][]MarloweStateAccountsInnerInner**]([]MarloweStateAccountsInnerInner.md) |  | 
**BoundValues** | [**[][]MarloweStateBoundValuesInnerInner**]([]MarloweStateBoundValuesInnerInner.md) |  | 
**Choices** | [**[][]MarloweStateChoicesInnerInner**]([]MarloweStateChoicesInnerInner.md) |  | 
**MinTime** | **int32** |  | 

## Methods

### NewMarloweState

`func NewMarloweState(accounts [][]MarloweStateAccountsInnerInner, boundValues [][]MarloweStateBoundValuesInnerInner, choices [][]MarloweStateChoicesInnerInner, minTime int32, ) *MarloweState`

NewMarloweState instantiates a new MarloweState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarloweStateWithDefaults

`func NewMarloweStateWithDefaults() *MarloweState`

NewMarloweStateWithDefaults instantiates a new MarloweState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *MarloweState) GetAccounts() [][]MarloweStateAccountsInnerInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *MarloweState) GetAccountsOk() (*[][]MarloweStateAccountsInnerInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *MarloweState) SetAccounts(v [][]MarloweStateAccountsInnerInner)`

SetAccounts sets Accounts field to given value.


### GetBoundValues

`func (o *MarloweState) GetBoundValues() [][]MarloweStateBoundValuesInnerInner`

GetBoundValues returns the BoundValues field if non-nil, zero value otherwise.

### GetBoundValuesOk

`func (o *MarloweState) GetBoundValuesOk() (*[][]MarloweStateBoundValuesInnerInner, bool)`

GetBoundValuesOk returns a tuple with the BoundValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundValues

`func (o *MarloweState) SetBoundValues(v [][]MarloweStateBoundValuesInnerInner)`

SetBoundValues sets BoundValues field to given value.


### GetChoices

`func (o *MarloweState) GetChoices() [][]MarloweStateChoicesInnerInner`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *MarloweState) GetChoicesOk() (*[][]MarloweStateChoicesInnerInner, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *MarloweState) SetChoices(v [][]MarloweStateChoicesInnerInner)`

SetChoices sets Choices field to given value.


### GetMinTime

`func (o *MarloweState) GetMinTime() int32`

GetMinTime returns the MinTime field if non-nil, zero value otherwise.

### GetMinTimeOk

`func (o *MarloweState) GetMinTimeOk() (*int32, bool)`

GetMinTimeOk returns a tuple with the MinTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTime

`func (o *MarloweState) SetMinTime(v int32)`

SetMinTime sets MinTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


