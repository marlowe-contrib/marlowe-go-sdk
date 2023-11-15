# ApplyInputsTxEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**TransactionId** | **string** | The hex-encoded identifier of a Cardano transaction | 
**Tx** | [**TextEnvelope**](TextEnvelope.md) |  | 

## Methods

### NewApplyInputsTxEnvelope

`func NewApplyInputsTxEnvelope(contractId string, transactionId string, tx TextEnvelope, ) *ApplyInputsTxEnvelope`

NewApplyInputsTxEnvelope instantiates a new ApplyInputsTxEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyInputsTxEnvelopeWithDefaults

`func NewApplyInputsTxEnvelopeWithDefaults() *ApplyInputsTxEnvelope`

NewApplyInputsTxEnvelopeWithDefaults instantiates a new ApplyInputsTxEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *ApplyInputsTxEnvelope) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ApplyInputsTxEnvelope) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ApplyInputsTxEnvelope) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetTransactionId

`func (o *ApplyInputsTxEnvelope) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ApplyInputsTxEnvelope) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ApplyInputsTxEnvelope) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTx

`func (o *ApplyInputsTxEnvelope) GetTx() TextEnvelope`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ApplyInputsTxEnvelope) GetTxOk() (*TextEnvelope, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ApplyInputsTxEnvelope) SetTx(v TextEnvelope)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


