# WithdrawTxBodyEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxBody** | [**TextEnvelope**](TextEnvelope.md) |  | 
**WithdrawalId** | **string** | The hex-encoded identifier of a Cardano transaction | 

## Methods

### NewWithdrawTxBodyEnvelope

`func NewWithdrawTxBodyEnvelope(txBody TextEnvelope, withdrawalId string, ) *WithdrawTxBodyEnvelope`

NewWithdrawTxBodyEnvelope instantiates a new WithdrawTxBodyEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawTxBodyEnvelopeWithDefaults

`func NewWithdrawTxBodyEnvelopeWithDefaults() *WithdrawTxBodyEnvelope`

NewWithdrawTxBodyEnvelopeWithDefaults instantiates a new WithdrawTxBodyEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxBody

`func (o *WithdrawTxBodyEnvelope) GetTxBody() TextEnvelope`

GetTxBody returns the TxBody field if non-nil, zero value otherwise.

### GetTxBodyOk

`func (o *WithdrawTxBodyEnvelope) GetTxBodyOk() (*TextEnvelope, bool)`

GetTxBodyOk returns a tuple with the TxBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBody

`func (o *WithdrawTxBodyEnvelope) SetTxBody(v TextEnvelope)`

SetTxBody sets TxBody field to given value.


### GetWithdrawalId

`func (o *WithdrawTxBodyEnvelope) GetWithdrawalId() string`

GetWithdrawalId returns the WithdrawalId field if non-nil, zero value otherwise.

### GetWithdrawalIdOk

`func (o *WithdrawTxBodyEnvelope) GetWithdrawalIdOk() (*string, bool)`

GetWithdrawalIdOk returns a tuple with the WithdrawalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalId

`func (o *WithdrawTxBodyEnvelope) SetWithdrawalId(v string)`

SetWithdrawalId sets WithdrawalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


