# ApplicableInputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | [**[]CanChoose**](CanChoose.md) |  | 
**Deposits** | [**[]CanDeposit**](CanDeposit.md) |  | 
**Notify** | Pointer to [**CanNotify**](CanNotify.md) |  | [optional] 

## Methods

### NewApplicableInputs

`func NewApplicableInputs(choices []CanChoose, deposits []CanDeposit, ) *ApplicableInputs`

NewApplicableInputs instantiates a new ApplicableInputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicableInputsWithDefaults

`func NewApplicableInputsWithDefaults() *ApplicableInputs`

NewApplicableInputsWithDefaults instantiates a new ApplicableInputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *ApplicableInputs) GetChoices() []CanChoose`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicableInputs) GetChoicesOk() (*[]CanChoose, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicableInputs) SetChoices(v []CanChoose)`

SetChoices sets Choices field to given value.


### GetDeposits

`func (o *ApplicableInputs) GetDeposits() []CanDeposit`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *ApplicableInputs) GetDepositsOk() (*[]CanDeposit, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *ApplicableInputs) SetDeposits(v []CanDeposit)`

SetDeposits sets Deposits field to given value.


### GetNotify

`func (o *ApplicableInputs) GetNotify() CanNotify`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *ApplicableInputs) GetNotifyOk() (*CanNotify, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *ApplicableInputs) SetNotify(v CanNotify)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *ApplicableInputs) HasNotify() bool`

HasNotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


