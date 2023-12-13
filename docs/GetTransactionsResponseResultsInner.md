# GetTransactionsResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**ApplyInputsResponseLinks**](ApplyInputsResponseLinks.md) |  | 
**Resource** | [**TxHeader**](TxHeader.md) |  | 

## Methods

### NewGetTransactionsResponseResultsInner

`func NewGetTransactionsResponseResultsInner(links ApplyInputsResponseLinks, resource TxHeader, ) *GetTransactionsResponseResultsInner`

NewGetTransactionsResponseResultsInner instantiates a new GetTransactionsResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionsResponseResultsInnerWithDefaults

`func NewGetTransactionsResponseResultsInnerWithDefaults() *GetTransactionsResponseResultsInner`

NewGetTransactionsResponseResultsInnerWithDefaults instantiates a new GetTransactionsResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetTransactionsResponseResultsInner) GetLinks() ApplyInputsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTransactionsResponseResultsInner) GetLinksOk() (*ApplyInputsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTransactionsResponseResultsInner) SetLinks(v ApplyInputsResponseLinks)`

SetLinks sets Links field to given value.


### GetResource

`func (o *GetTransactionsResponseResultsInner) GetResource() TxHeader`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetTransactionsResponseResultsInner) GetResourceOk() (*TxHeader, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetTransactionsResponseResultsInner) SetResource(v TxHeader)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


