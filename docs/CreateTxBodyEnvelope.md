# CreateTxBodyEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | A reference to a transaction output with a transaction ID and index. | 
**SafetyErrors** | Pointer to [**[]SafetyError**](SafetyError.md) |  | [optional] 
**TxBody** | [**TextEnvelope**](TextEnvelope.md) |  | 

## Methods

### NewCreateTxBodyEnvelope

`func NewCreateTxBodyEnvelope(contractId string, txBody TextEnvelope, ) *CreateTxBodyEnvelope`

NewCreateTxBodyEnvelope instantiates a new CreateTxBodyEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTxBodyEnvelopeWithDefaults

`func NewCreateTxBodyEnvelopeWithDefaults() *CreateTxBodyEnvelope`

NewCreateTxBodyEnvelopeWithDefaults instantiates a new CreateTxBodyEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *CreateTxBodyEnvelope) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *CreateTxBodyEnvelope) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *CreateTxBodyEnvelope) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetSafetyErrors

`func (o *CreateTxBodyEnvelope) GetSafetyErrors() []SafetyError`

GetSafetyErrors returns the SafetyErrors field if non-nil, zero value otherwise.

### GetSafetyErrorsOk

`func (o *CreateTxBodyEnvelope) GetSafetyErrorsOk() (*[]SafetyError, bool)`

GetSafetyErrorsOk returns a tuple with the SafetyErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyErrors

`func (o *CreateTxBodyEnvelope) SetSafetyErrors(v []SafetyError)`

SetSafetyErrors sets SafetyErrors field to given value.

### HasSafetyErrors

`func (o *CreateTxBodyEnvelope) HasSafetyErrors() bool`

HasSafetyErrors returns a boolean if a field has been set.

### GetTxBody

`func (o *CreateTxBodyEnvelope) GetTxBody() TextEnvelope`

GetTxBody returns the TxBody field if non-nil, zero value otherwise.

### GetTxBodyOk

`func (o *CreateTxBodyEnvelope) GetTxBodyOk() (*TextEnvelope, bool)`

GetTxBodyOk returns a tuple with the TxBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBody

`func (o *CreateTxBodyEnvelope) SetTxBody(v TextEnvelope)`

SetTxBody sets TxBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


