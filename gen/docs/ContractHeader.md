# ContractHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | Pointer to [**BlockHeader**](BlockHeader.md) |  | [optional] 
**Continuations** | Pointer to **string** |  | [optional] 
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**Metadata** | **map[string]interface{}** |  | 
**RoleTokenMintingPolicyId** | **string** | The hex-encoded minting policy ID for a native Cardano token | 
**Status** | [**TxStatus**](TxStatus.md) |  | 
**Tags** | **map[string]interface{}** |  | 
**Version** | [**MarloweVersion**](MarloweVersion.md) |  | 

## Methods

### NewContractHeader

`func NewContractHeader(contractId string, metadata map[string]interface{}, roleTokenMintingPolicyId string, status TxStatus, tags map[string]interface{}, version MarloweVersion, ) *ContractHeader`

NewContractHeader instantiates a new ContractHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractHeaderWithDefaults

`func NewContractHeaderWithDefaults() *ContractHeader`

NewContractHeaderWithDefaults instantiates a new ContractHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *ContractHeader) GetBlock() BlockHeader`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ContractHeader) GetBlockOk() (*BlockHeader, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ContractHeader) SetBlock(v BlockHeader)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *ContractHeader) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetContinuations

`func (o *ContractHeader) GetContinuations() string`

GetContinuations returns the Continuations field if non-nil, zero value otherwise.

### GetContinuationsOk

`func (o *ContractHeader) GetContinuationsOk() (*string, bool)`

GetContinuationsOk returns a tuple with the Continuations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuations

`func (o *ContractHeader) SetContinuations(v string)`

SetContinuations sets Continuations field to given value.

### HasContinuations

`func (o *ContractHeader) HasContinuations() bool`

HasContinuations returns a boolean if a field has been set.

### GetContractId

`func (o *ContractHeader) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractHeader) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractHeader) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetMetadata

`func (o *ContractHeader) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContractHeader) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContractHeader) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetRoleTokenMintingPolicyId

`func (o *ContractHeader) GetRoleTokenMintingPolicyId() string`

GetRoleTokenMintingPolicyId returns the RoleTokenMintingPolicyId field if non-nil, zero value otherwise.

### GetRoleTokenMintingPolicyIdOk

`func (o *ContractHeader) GetRoleTokenMintingPolicyIdOk() (*string, bool)`

GetRoleTokenMintingPolicyIdOk returns a tuple with the RoleTokenMintingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTokenMintingPolicyId

`func (o *ContractHeader) SetRoleTokenMintingPolicyId(v string)`

SetRoleTokenMintingPolicyId sets RoleTokenMintingPolicyId field to given value.


### GetStatus

`func (o *ContractHeader) GetStatus() TxStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractHeader) GetStatusOk() (*TxStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractHeader) SetStatus(v TxStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *ContractHeader) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContractHeader) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContractHeader) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetVersion

`func (o *ContractHeader) GetVersion() MarloweVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContractHeader) GetVersionOk() (*MarloweVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContractHeader) SetVersion(v MarloweVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


