# GetPayoutsResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**GetPayoutsResponseResultsInnerLinks**](GetPayoutsResponseResultsInnerLinks.md) |  | 
**Resource** | [**PayoutHeader**](PayoutHeader.md) |  | 

## Methods

### NewGetPayoutsResponseResultsInner

`func NewGetPayoutsResponseResultsInner(links GetPayoutsResponseResultsInnerLinks, resource PayoutHeader, ) *GetPayoutsResponseResultsInner`

NewGetPayoutsResponseResultsInner instantiates a new GetPayoutsResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayoutsResponseResultsInnerWithDefaults

`func NewGetPayoutsResponseResultsInnerWithDefaults() *GetPayoutsResponseResultsInner`

NewGetPayoutsResponseResultsInnerWithDefaults instantiates a new GetPayoutsResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetPayoutsResponseResultsInner) GetLinks() GetPayoutsResponseResultsInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPayoutsResponseResultsInner) GetLinksOk() (*GetPayoutsResponseResultsInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPayoutsResponseResultsInner) SetLinks(v GetPayoutsResponseResultsInnerLinks)`

SetLinks sets Links field to given value.


### GetResource

`func (o *GetPayoutsResponseResultsInner) GetResource() PayoutHeader`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetPayoutsResponseResultsInner) GetResourceOk() (*PayoutHeader, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetPayoutsResponseResultsInner) SetResource(v PayoutHeader)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


