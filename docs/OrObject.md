# OrObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Either** | [**ObservationObject**](ObservationObject.md) |  | 
**Or** | [**ObservationObject**](ObservationObject.md) |  | 

## Methods

### NewOrObject

`func NewOrObject(either ObservationObject, or ObservationObject, ) *OrObject`

NewOrObject instantiates a new OrObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrObjectWithDefaults

`func NewOrObjectWithDefaults() *OrObject`

NewOrObjectWithDefaults instantiates a new OrObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEither

`func (o *OrObject) GetEither() ObservationObject`

GetEither returns the Either field if non-nil, zero value otherwise.

### GetEitherOk

`func (o *OrObject) GetEitherOk() (*ObservationObject, bool)`

GetEitherOk returns a tuple with the Either field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEither

`func (o *OrObject) SetEither(v ObservationObject)`

SetEither sets Either field to given value.


### GetOr

`func (o *OrObject) GetOr() ObservationObject`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *OrObject) GetOrOk() (*ObservationObject, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *OrObject) SetOr(v ObservationObject)`

SetOr sets Or field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


