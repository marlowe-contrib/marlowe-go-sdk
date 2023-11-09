# IntervalError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidInterval** | [**IntervalErrorOneOfInvalidInterval**](IntervalErrorOneOfInvalidInterval.md) |  | 
**IntervalInPastError** | [**IntervalErrorOneOf1IntervalInPastError**](IntervalErrorOneOf1IntervalInPastError.md) |  | 

## Methods

### NewIntervalError

`func NewIntervalError(invalidInterval IntervalErrorOneOfInvalidInterval, intervalInPastError IntervalErrorOneOf1IntervalInPastError, ) *IntervalError`

NewIntervalError instantiates a new IntervalError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervalErrorWithDefaults

`func NewIntervalErrorWithDefaults() *IntervalError`

NewIntervalErrorWithDefaults instantiates a new IntervalError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidInterval

`func (o *IntervalError) GetInvalidInterval() IntervalErrorOneOfInvalidInterval`

GetInvalidInterval returns the InvalidInterval field if non-nil, zero value otherwise.

### GetInvalidIntervalOk

`func (o *IntervalError) GetInvalidIntervalOk() (*IntervalErrorOneOfInvalidInterval, bool)`

GetInvalidIntervalOk returns a tuple with the InvalidInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidInterval

`func (o *IntervalError) SetInvalidInterval(v IntervalErrorOneOfInvalidInterval)`

SetInvalidInterval sets InvalidInterval field to given value.


### GetIntervalInPastError

`func (o *IntervalError) GetIntervalInPastError() IntervalErrorOneOf1IntervalInPastError`

GetIntervalInPastError returns the IntervalInPastError field if non-nil, zero value otherwise.

### GetIntervalInPastErrorOk

`func (o *IntervalError) GetIntervalInPastErrorOk() (*IntervalErrorOneOf1IntervalInPastError, bool)`

GetIntervalInPastErrorOk returns a tuple with the IntervalInPastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalInPastError

`func (o *IntervalError) SetIntervalInPastError(v IntervalErrorOneOf1IntervalInPastError)`

SetIntervalInPastError sets IntervalInPastError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


