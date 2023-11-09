# ContractState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | [**Assets**](Assets.md) |  | 
**Block** | Pointer to [**BlockHeader**](BlockHeader.md) |  | [optional] 
**Continuations** | Pointer to **string** |  | [optional] 
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**CurrentContract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**InitialContract** | [**Contract**](Contract.md) |  | 
**Metadata** | **map[string]interface{}** |  | 
**RoleTokenMintingPolicyId** | **string** | The hex-encoded minting policy ID for a native Cardano token | 
**State** | Pointer to [**MarloweState**](MarloweState.md) |  | [optional] 
**Status** | [**TxStatus**](TxStatus.md) |  | 
**Tags** | **map[string]interface{}** |  | 
**TxBody** | Pointer to [**TextEnvelope**](TextEnvelope.md) |  | [optional] 
**UnclaimedPayouts** | [**[]Payout**](Payout.md) |  | 
**Utxo** | Pointer to **string** | A reference to a transaction output with a transaction ID and index. | [optional] 
**Version** | [**MarloweVersion**](MarloweVersion.md) |  | 

## Methods

### NewContractState

`func NewContractState(assets Assets, contractId string, initialContract Contract, metadata map[string]interface{}, roleTokenMintingPolicyId string, status TxStatus, tags map[string]interface{}, unclaimedPayouts []Payout, version MarloweVersion, ) *ContractState`

NewContractState instantiates a new ContractState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractStateWithDefaults

`func NewContractStateWithDefaults() *ContractState`

NewContractStateWithDefaults instantiates a new ContractState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *ContractState) GetAssets() Assets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ContractState) GetAssetsOk() (*Assets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ContractState) SetAssets(v Assets)`

SetAssets sets Assets field to given value.


### GetBlock

`func (o *ContractState) GetBlock() BlockHeader`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ContractState) GetBlockOk() (*BlockHeader, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ContractState) SetBlock(v BlockHeader)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *ContractState) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetContinuations

`func (o *ContractState) GetContinuations() string`

GetContinuations returns the Continuations field if non-nil, zero value otherwise.

### GetContinuationsOk

`func (o *ContractState) GetContinuationsOk() (*string, bool)`

GetContinuationsOk returns a tuple with the Continuations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuations

`func (o *ContractState) SetContinuations(v string)`

SetContinuations sets Continuations field to given value.

### HasContinuations

`func (o *ContractState) HasContinuations() bool`

HasContinuations returns a boolean if a field has been set.

### GetContractId

`func (o *ContractState) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractState) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractState) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetCurrentContract

`func (o *ContractState) GetCurrentContract() Contract`

GetCurrentContract returns the CurrentContract field if non-nil, zero value otherwise.

### GetCurrentContractOk

`func (o *ContractState) GetCurrentContractOk() (*Contract, bool)`

GetCurrentContractOk returns a tuple with the CurrentContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentContract

`func (o *ContractState) SetCurrentContract(v Contract)`

SetCurrentContract sets CurrentContract field to given value.

### HasCurrentContract

`func (o *ContractState) HasCurrentContract() bool`

HasCurrentContract returns a boolean if a field has been set.

### GetInitialContract

`func (o *ContractState) GetInitialContract() Contract`

GetInitialContract returns the InitialContract field if non-nil, zero value otherwise.

### GetInitialContractOk

`func (o *ContractState) GetInitialContractOk() (*Contract, bool)`

GetInitialContractOk returns a tuple with the InitialContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialContract

`func (o *ContractState) SetInitialContract(v Contract)`

SetInitialContract sets InitialContract field to given value.


### GetMetadata

`func (o *ContractState) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContractState) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContractState) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetRoleTokenMintingPolicyId

`func (o *ContractState) GetRoleTokenMintingPolicyId() string`

GetRoleTokenMintingPolicyId returns the RoleTokenMintingPolicyId field if non-nil, zero value otherwise.

### GetRoleTokenMintingPolicyIdOk

`func (o *ContractState) GetRoleTokenMintingPolicyIdOk() (*string, bool)`

GetRoleTokenMintingPolicyIdOk returns a tuple with the RoleTokenMintingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTokenMintingPolicyId

`func (o *ContractState) SetRoleTokenMintingPolicyId(v string)`

SetRoleTokenMintingPolicyId sets RoleTokenMintingPolicyId field to given value.


### GetState

`func (o *ContractState) GetState() MarloweState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContractState) GetStateOk() (*MarloweState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContractState) SetState(v MarloweState)`

SetState sets State field to given value.

### HasState

`func (o *ContractState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *ContractState) GetStatus() TxStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractState) GetStatusOk() (*TxStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractState) SetStatus(v TxStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *ContractState) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContractState) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContractState) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetTxBody

`func (o *ContractState) GetTxBody() TextEnvelope`

GetTxBody returns the TxBody field if non-nil, zero value otherwise.

### GetTxBodyOk

`func (o *ContractState) GetTxBodyOk() (*TextEnvelope, bool)`

GetTxBodyOk returns a tuple with the TxBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBody

`func (o *ContractState) SetTxBody(v TextEnvelope)`

SetTxBody sets TxBody field to given value.

### HasTxBody

`func (o *ContractState) HasTxBody() bool`

HasTxBody returns a boolean if a field has been set.

### GetUnclaimedPayouts

`func (o *ContractState) GetUnclaimedPayouts() []Payout`

GetUnclaimedPayouts returns the UnclaimedPayouts field if non-nil, zero value otherwise.

### GetUnclaimedPayoutsOk

`func (o *ContractState) GetUnclaimedPayoutsOk() (*[]Payout, bool)`

GetUnclaimedPayoutsOk returns a tuple with the UnclaimedPayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclaimedPayouts

`func (o *ContractState) SetUnclaimedPayouts(v []Payout)`

SetUnclaimedPayouts sets UnclaimedPayouts field to given value.


### GetUtxo

`func (o *ContractState) GetUtxo() string`

GetUtxo returns the Utxo field if non-nil, zero value otherwise.

### GetUtxoOk

`func (o *ContractState) GetUtxoOk() (*string, bool)`

GetUtxoOk returns a tuple with the Utxo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxo

`func (o *ContractState) SetUtxo(v string)`

SetUtxo sets Utxo field to given value.

### HasUtxo

`func (o *ContractState) HasUtxo() bool`

HasUtxo returns a boolean if a field has been set.

### GetVersion

`func (o *ContractState) GetVersion() MarloweVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContractState) GetVersionOk() (*MarloweVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContractState) SetVersion(v MarloweVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


