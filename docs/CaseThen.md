# CaseThen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Case** | [**Action**](Action.md) |  | 
**Then** | [**Contract**](Contract.md) |  | 

## Methods

### NewCaseThen

`func NewCaseThen(case_ Action, then Contract, ) *CaseThen`

NewCaseThen instantiates a new CaseThen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseThenWithDefaults

`func NewCaseThenWithDefaults() *CaseThen`

NewCaseThenWithDefaults instantiates a new CaseThen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCase

`func (o *CaseThen) GetCase() Action`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *CaseThen) GetCaseOk() (*Action, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *CaseThen) SetCase(v Action)`

SetCase sets Case field to given value.


### GetThen

`func (o *CaseThen) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *CaseThen) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *CaseThen) SetThen(v Contract)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


