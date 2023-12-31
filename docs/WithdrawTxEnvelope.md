# WithdrawTxEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | [**TextEnvelope**](TextEnvelope.md) |  | 
**WithdrawalId** | **string** | The hex-encoded identifier of a Cardano transaction | 

## Methods

### NewWithdrawTxEnvelope

`func NewWithdrawTxEnvelope(tx TextEnvelope, withdrawalId string, ) *WithdrawTxEnvelope`

NewWithdrawTxEnvelope instantiates a new WithdrawTxEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawTxEnvelopeWithDefaults

`func NewWithdrawTxEnvelopeWithDefaults() *WithdrawTxEnvelope`

NewWithdrawTxEnvelopeWithDefaults instantiates a new WithdrawTxEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *WithdrawTxEnvelope) GetTx() TextEnvelope`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *WithdrawTxEnvelope) GetTxOk() (*TextEnvelope, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *WithdrawTxEnvelope) SetTx(v TextEnvelope)`

SetTx sets Tx field to given value.


### GetWithdrawalId

`func (o *WithdrawTxEnvelope) GetWithdrawalId() string`

GetWithdrawalId returns the WithdrawalId field if non-nil, zero value otherwise.

### GetWithdrawalIdOk

`func (o *WithdrawTxEnvelope) GetWithdrawalIdOk() (*string, bool)`

GetWithdrawalIdOk returns a tuple with the WithdrawalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalId

`func (o *WithdrawTxEnvelope) SetWithdrawalId(v string)`

SetWithdrawalId sets WithdrawalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


