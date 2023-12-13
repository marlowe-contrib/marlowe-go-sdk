# CreateTxEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**SafetyErrors** | Pointer to [**[]SafetyError**](SafetyError.md) |  | [optional] 
**Tx** | [**TextEnvelope**](TextEnvelope.md) |  | 

## Methods

### NewCreateTxEnvelope

`func NewCreateTxEnvelope(contractId string, tx TextEnvelope, ) *CreateTxEnvelope`

NewCreateTxEnvelope instantiates a new CreateTxEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTxEnvelopeWithDefaults

`func NewCreateTxEnvelopeWithDefaults() *CreateTxEnvelope`

NewCreateTxEnvelopeWithDefaults instantiates a new CreateTxEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *CreateTxEnvelope) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *CreateTxEnvelope) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *CreateTxEnvelope) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetSafetyErrors

`func (o *CreateTxEnvelope) GetSafetyErrors() []SafetyError`

GetSafetyErrors returns the SafetyErrors field if non-nil, zero value otherwise.

### GetSafetyErrorsOk

`func (o *CreateTxEnvelope) GetSafetyErrorsOk() (*[]SafetyError, bool)`

GetSafetyErrorsOk returns a tuple with the SafetyErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyErrors

`func (o *CreateTxEnvelope) SetSafetyErrors(v []SafetyError)`

SetSafetyErrors sets SafetyErrors field to given value.

### HasSafetyErrors

`func (o *CreateTxEnvelope) HasSafetyErrors() bool`

HasSafetyErrors returns a boolean if a field has been set.

### GetTx

`func (o *CreateTxEnvelope) GetTx() TextEnvelope`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *CreateTxEnvelope) GetTxOk() (*TextEnvelope, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *CreateTxEnvelope) SetTx(v TextEnvelope)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


