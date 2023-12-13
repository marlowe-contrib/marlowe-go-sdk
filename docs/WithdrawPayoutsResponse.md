# WithdrawPayoutsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**GetWithdrawalsResponseResultsInnerLinks**](GetWithdrawalsResponseResultsInnerLinks.md) |  | 
**Resource** | [**WithdrawTxEnvelope**](WithdrawTxEnvelope.md) |  | 

## Methods

### NewWithdrawPayoutsResponse

`func NewWithdrawPayoutsResponse(links GetWithdrawalsResponseResultsInnerLinks, resource WithdrawTxEnvelope, ) *WithdrawPayoutsResponse`

NewWithdrawPayoutsResponse instantiates a new WithdrawPayoutsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawPayoutsResponseWithDefaults

`func NewWithdrawPayoutsResponseWithDefaults() *WithdrawPayoutsResponse`

NewWithdrawPayoutsResponseWithDefaults instantiates a new WithdrawPayoutsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WithdrawPayoutsResponse) GetLinks() GetWithdrawalsResponseResultsInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WithdrawPayoutsResponse) GetLinksOk() (*GetWithdrawalsResponseResultsInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WithdrawPayoutsResponse) SetLinks(v GetWithdrawalsResponseResultsInnerLinks)`

SetLinks sets Links field to given value.


### GetResource

`func (o *WithdrawPayoutsResponse) GetResource() WithdrawTxEnvelope`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *WithdrawPayoutsResponse) GetResourceOk() (*WithdrawTxEnvelope, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *WithdrawPayoutsResponse) SetResource(v WithdrawTxEnvelope)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


