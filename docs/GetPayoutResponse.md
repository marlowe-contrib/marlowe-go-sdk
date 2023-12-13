# GetPayoutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**GetPayoutResponseLinks**](GetPayoutResponseLinks.md) |  | 
**Resource** | [**PayoutState**](PayoutState.md) |  | 

## Methods

### NewGetPayoutResponse

`func NewGetPayoutResponse(links GetPayoutResponseLinks, resource PayoutState, ) *GetPayoutResponse`

NewGetPayoutResponse instantiates a new GetPayoutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayoutResponseWithDefaults

`func NewGetPayoutResponseWithDefaults() *GetPayoutResponse`

NewGetPayoutResponseWithDefaults instantiates a new GetPayoutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetPayoutResponse) GetLinks() GetPayoutResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPayoutResponse) GetLinksOk() (*GetPayoutResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPayoutResponse) SetLinks(v GetPayoutResponseLinks)`

SetLinks sets Links field to given value.


### GetResource

`func (o *GetPayoutResponse) GetResource() PayoutState`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetPayoutResponse) GetResourceOk() (*PayoutState, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetPayoutResponse) SetResource(v PayoutState)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


