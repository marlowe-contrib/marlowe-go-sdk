# Tx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | [**Assets**](Assets.md) |  | 
**Block** | Pointer to [**BlockHeader**](BlockHeader.md) |  | [optional] 
**ConsumingTx** | Pointer to **string** | The hex-encoded identifier of a Cardano transaction | [optional] 
**Continuations** | Pointer to **string** |  | [optional] 
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**InputUtxo** | **string** | A reference to a transaction output with a transaction ID and index. | 
**Inputs** | [**[]Input**](Input.md) |  | 
**InvalidBefore** | **string** |  | 
**InvalidHereafter** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**OutputContract** | Pointer to [**Contract**](Contract.md) |  | [optional] 
**OutputState** | Pointer to [**MarloweState**](MarloweState.md) |  | [optional] 
**OutputUtxo** | Pointer to **string** | A reference to a transaction output with a transaction ID and index. | [optional] 
**Payouts** | [**[]Payout**](Payout.md) |  | 
**Status** | [**TxStatus**](TxStatus.md) |  | 
**Tags** | **map[string]interface{}** |  | 
**TransactionId** | **string** | The hex-encoded identifier of a Cardano transaction | 
**TxBody** | Pointer to [**TextEnvelope**](TextEnvelope.md) |  | [optional] 

## Methods

### NewTx

`func NewTx(assets Assets, contractId string, inputUtxo string, inputs []Input, invalidBefore string, invalidHereafter string, metadata map[string]interface{}, payouts []Payout, status TxStatus, tags map[string]interface{}, transactionId string, ) *Tx`

NewTx instantiates a new Tx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxWithDefaults

`func NewTxWithDefaults() *Tx`

NewTxWithDefaults instantiates a new Tx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *Tx) GetAssets() Assets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Tx) GetAssetsOk() (*Assets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Tx) SetAssets(v Assets)`

SetAssets sets Assets field to given value.


### GetBlock

`func (o *Tx) GetBlock() BlockHeader`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *Tx) GetBlockOk() (*BlockHeader, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *Tx) SetBlock(v BlockHeader)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *Tx) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetConsumingTx

`func (o *Tx) GetConsumingTx() string`

GetConsumingTx returns the ConsumingTx field if non-nil, zero value otherwise.

### GetConsumingTxOk

`func (o *Tx) GetConsumingTxOk() (*string, bool)`

GetConsumingTxOk returns a tuple with the ConsumingTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumingTx

`func (o *Tx) SetConsumingTx(v string)`

SetConsumingTx sets ConsumingTx field to given value.

### HasConsumingTx

`func (o *Tx) HasConsumingTx() bool`

HasConsumingTx returns a boolean if a field has been set.

### GetContinuations

`func (o *Tx) GetContinuations() string`

GetContinuations returns the Continuations field if non-nil, zero value otherwise.

### GetContinuationsOk

`func (o *Tx) GetContinuationsOk() (*string, bool)`

GetContinuationsOk returns a tuple with the Continuations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuations

`func (o *Tx) SetContinuations(v string)`

SetContinuations sets Continuations field to given value.

### HasContinuations

`func (o *Tx) HasContinuations() bool`

HasContinuations returns a boolean if a field has been set.

### GetContractId

`func (o *Tx) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *Tx) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *Tx) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetInputUtxo

`func (o *Tx) GetInputUtxo() string`

GetInputUtxo returns the InputUtxo field if non-nil, zero value otherwise.

### GetInputUtxoOk

`func (o *Tx) GetInputUtxoOk() (*string, bool)`

GetInputUtxoOk returns a tuple with the InputUtxo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputUtxo

`func (o *Tx) SetInputUtxo(v string)`

SetInputUtxo sets InputUtxo field to given value.


### GetInputs

`func (o *Tx) GetInputs() []Input`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *Tx) GetInputsOk() (*[]Input, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *Tx) SetInputs(v []Input)`

SetInputs sets Inputs field to given value.


### GetInvalidBefore

`func (o *Tx) GetInvalidBefore() string`

GetInvalidBefore returns the InvalidBefore field if non-nil, zero value otherwise.

### GetInvalidBeforeOk

`func (o *Tx) GetInvalidBeforeOk() (*string, bool)`

GetInvalidBeforeOk returns a tuple with the InvalidBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidBefore

`func (o *Tx) SetInvalidBefore(v string)`

SetInvalidBefore sets InvalidBefore field to given value.


### GetInvalidHereafter

`func (o *Tx) GetInvalidHereafter() string`

GetInvalidHereafter returns the InvalidHereafter field if non-nil, zero value otherwise.

### GetInvalidHereafterOk

`func (o *Tx) GetInvalidHereafterOk() (*string, bool)`

GetInvalidHereafterOk returns a tuple with the InvalidHereafter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidHereafter

`func (o *Tx) SetInvalidHereafter(v string)`

SetInvalidHereafter sets InvalidHereafter field to given value.


### GetMetadata

`func (o *Tx) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Tx) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Tx) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetOutputContract

`func (o *Tx) GetOutputContract() Contract`

GetOutputContract returns the OutputContract field if non-nil, zero value otherwise.

### GetOutputContractOk

`func (o *Tx) GetOutputContractOk() (*Contract, bool)`

GetOutputContractOk returns a tuple with the OutputContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputContract

`func (o *Tx) SetOutputContract(v Contract)`

SetOutputContract sets OutputContract field to given value.

### HasOutputContract

`func (o *Tx) HasOutputContract() bool`

HasOutputContract returns a boolean if a field has been set.

### GetOutputState

`func (o *Tx) GetOutputState() MarloweState`

GetOutputState returns the OutputState field if non-nil, zero value otherwise.

### GetOutputStateOk

`func (o *Tx) GetOutputStateOk() (*MarloweState, bool)`

GetOutputStateOk returns a tuple with the OutputState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputState

`func (o *Tx) SetOutputState(v MarloweState)`

SetOutputState sets OutputState field to given value.

### HasOutputState

`func (o *Tx) HasOutputState() bool`

HasOutputState returns a boolean if a field has been set.

### GetOutputUtxo

`func (o *Tx) GetOutputUtxo() string`

GetOutputUtxo returns the OutputUtxo field if non-nil, zero value otherwise.

### GetOutputUtxoOk

`func (o *Tx) GetOutputUtxoOk() (*string, bool)`

GetOutputUtxoOk returns a tuple with the OutputUtxo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputUtxo

`func (o *Tx) SetOutputUtxo(v string)`

SetOutputUtxo sets OutputUtxo field to given value.

### HasOutputUtxo

`func (o *Tx) HasOutputUtxo() bool`

HasOutputUtxo returns a boolean if a field has been set.

### GetPayouts

`func (o *Tx) GetPayouts() []Payout`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *Tx) GetPayoutsOk() (*[]Payout, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *Tx) SetPayouts(v []Payout)`

SetPayouts sets Payouts field to given value.


### GetStatus

`func (o *Tx) GetStatus() TxStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Tx) GetStatusOk() (*TxStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Tx) SetStatus(v TxStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *Tx) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Tx) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Tx) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetTransactionId

`func (o *Tx) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Tx) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Tx) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTxBody

`func (o *Tx) GetTxBody() TextEnvelope`

GetTxBody returns the TxBody field if non-nil, zero value otherwise.

### GetTxBodyOk

`func (o *Tx) GetTxBodyOk() (*TextEnvelope, bool)`

GetTxBodyOk returns a tuple with the TxBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBody

`func (o *Tx) SetTxBody(v TextEnvelope)`

SetTxBody sets TxBody field to given value.

### HasTxBody

`func (o *Tx) HasTxBody() bool`

HasTxBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


