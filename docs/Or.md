# Or

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Either** | [**Observation**](Observation.md) |  | 
**Or** | [**Observation**](Observation.md) |  | 

## Methods

### NewOr

`func NewOr(either Observation, or Observation, ) *Or`

NewOr instantiates a new Or object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrWithDefaults

`func NewOrWithDefaults() *Or`

NewOrWithDefaults instantiates a new Or object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEither

`func (o *Or) GetEither() Observation`

GetEither returns the Either field if non-nil, zero value otherwise.

### GetEitherOk

`func (o *Or) GetEitherOk() (*Observation, bool)`

GetEitherOk returns a tuple with the Either field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEither

`func (o *Or) SetEither(v Observation)`

SetEither sets Either field to given value.


### GetOr

`func (o *Or) GetOr() Observation`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *Or) GetOrOk() (*Observation, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *Or) SetOr(v Observation)`

SetOr sets Or field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


