# PostContractsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | [**PostContractsRequestContract**](PostContractsRequestContract.md) |  | 
**Metadata** | **map[string]interface{}** |  | 
**MinUTxODeposit** | Pointer to **int64** |  | [optional] 
**Roles** | Pointer to [**RolesConfig**](RolesConfig.md) |  | [optional] 
**Tags** | **map[string]interface{}** |  | 
**Version** | [**MarloweVersion**](MarloweVersion.md) |  | 

## Methods

### NewPostContractsRequest

`func NewPostContractsRequest(contract PostContractsRequestContract, metadata map[string]interface{}, tags map[string]interface{}, version MarloweVersion, ) *PostContractsRequest`

NewPostContractsRequest instantiates a new PostContractsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostContractsRequestWithDefaults

`func NewPostContractsRequestWithDefaults() *PostContractsRequest`

NewPostContractsRequestWithDefaults instantiates a new PostContractsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *PostContractsRequest) GetContract() PostContractsRequestContract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *PostContractsRequest) GetContractOk() (*PostContractsRequestContract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *PostContractsRequest) SetContract(v PostContractsRequestContract)`

SetContract sets Contract field to given value.


### GetMetadata

`func (o *PostContractsRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostContractsRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostContractsRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetMinUTxODeposit

`func (o *PostContractsRequest) GetMinUTxODeposit() int64`

GetMinUTxODeposit returns the MinUTxODeposit field if non-nil, zero value otherwise.

### GetMinUTxODepositOk

`func (o *PostContractsRequest) GetMinUTxODepositOk() (*int64, bool)`

GetMinUTxODepositOk returns a tuple with the MinUTxODeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUTxODeposit

`func (o *PostContractsRequest) SetMinUTxODeposit(v int64)`

SetMinUTxODeposit sets MinUTxODeposit field to given value.

### HasMinUTxODeposit

`func (o *PostContractsRequest) HasMinUTxODeposit() bool`

HasMinUTxODeposit returns a boolean if a field has been set.

### GetRoles

`func (o *PostContractsRequest) GetRoles() RolesConfig`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PostContractsRequest) GetRolesOk() (*RolesConfig, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PostContractsRequest) SetRoles(v RolesConfig)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PostContractsRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTags

`func (o *PostContractsRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostContractsRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostContractsRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetVersion

`func (o *PostContractsRequest) GetVersion() MarloweVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostContractsRequest) GetVersionOk() (*MarloweVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostContractsRequest) SetVersion(v MarloweVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


