# AssetId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetName** | **string** |  | 
**PolicyId** | **string** | The hex-encoded minting policy ID for a native Cardano token | 

## Methods

### NewAssetId

`func NewAssetId(assetName string, policyId string, ) *AssetId`

NewAssetId instantiates a new AssetId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetIdWithDefaults

`func NewAssetIdWithDefaults() *AssetId`

NewAssetIdWithDefaults instantiates a new AssetId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetName

`func (o *AssetId) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *AssetId) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *AssetId) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.


### GetPolicyId

`func (o *AssetId) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *AssetId) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *AssetId) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


