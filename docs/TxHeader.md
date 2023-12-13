# TxHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | Pointer to [**BlockHeader**](BlockHeader.md) |  | [optional] 
**Continuations** | Pointer to **string** |  | [optional] 
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**Metadata** | [**map[string]Metadata**](Metadata.md) |  | 
**Status** | [**TxStatus**](TxStatus.md) |  | 
**Tags** | [**map[string]Metadata**](Metadata.md) |  | 
**TransactionId** | **string** | The hex-encoded identifier of a Cardano transaction | 
**Utxo** | Pointer to **string** | A reference to a transaction output with a transaction ID and index. | [optional] 

## Methods

### NewTxHeader

`func NewTxHeader(contractId string, metadata map[string]Metadata, status TxStatus, tags map[string]Metadata, transactionId string, ) *TxHeader`

NewTxHeader instantiates a new TxHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxHeaderWithDefaults

`func NewTxHeaderWithDefaults() *TxHeader`

NewTxHeaderWithDefaults instantiates a new TxHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *TxHeader) GetBlock() BlockHeader`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *TxHeader) GetBlockOk() (*BlockHeader, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *TxHeader) SetBlock(v BlockHeader)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *TxHeader) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetContinuations

`func (o *TxHeader) GetContinuations() string`

GetContinuations returns the Continuations field if non-nil, zero value otherwise.

### GetContinuationsOk

`func (o *TxHeader) GetContinuationsOk() (*string, bool)`

GetContinuationsOk returns a tuple with the Continuations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuations

`func (o *TxHeader) SetContinuations(v string)`

SetContinuations sets Continuations field to given value.

### HasContinuations

`func (o *TxHeader) HasContinuations() bool`

HasContinuations returns a boolean if a field has been set.

### GetContractId

`func (o *TxHeader) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *TxHeader) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *TxHeader) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetMetadata

`func (o *TxHeader) GetMetadata() map[string]Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TxHeader) GetMetadataOk() (*map[string]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TxHeader) SetMetadata(v map[string]Metadata)`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *TxHeader) GetStatus() TxStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TxHeader) GetStatusOk() (*TxStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TxHeader) SetStatus(v TxStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *TxHeader) GetTags() map[string]Metadata`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TxHeader) GetTagsOk() (*map[string]Metadata, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TxHeader) SetTags(v map[string]Metadata)`

SetTags sets Tags field to given value.


### GetTransactionId

`func (o *TxHeader) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TxHeader) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TxHeader) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetUtxo

`func (o *TxHeader) GetUtxo() string`

GetUtxo returns the Utxo field if non-nil, zero value otherwise.

### GetUtxoOk

`func (o *TxHeader) GetUtxoOk() (*string, bool)`

GetUtxoOk returns a tuple with the Utxo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxo

`func (o *TxHeader) SetUtxo(v string)`

SetUtxo sets Utxo field to given value.

### HasUtxo

`func (o *TxHeader) HasUtxo() bool`

HasUtxo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


