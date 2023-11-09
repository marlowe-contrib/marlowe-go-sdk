# WithdrawalHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | Pointer to [**BlockHeader**](BlockHeader.md) |  | [optional] 
**Status** | [**TxStatus**](TxStatus.md) |  | 
**WithdrawalId** | **string** | The hex-encoded identifier of a Cardano transaction | 

## Methods

### NewWithdrawalHeader

`func NewWithdrawalHeader(status TxStatus, withdrawalId string, ) *WithdrawalHeader`

NewWithdrawalHeader instantiates a new WithdrawalHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalHeaderWithDefaults

`func NewWithdrawalHeaderWithDefaults() *WithdrawalHeader`

NewWithdrawalHeaderWithDefaults instantiates a new WithdrawalHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *WithdrawalHeader) GetBlock() BlockHeader`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *WithdrawalHeader) GetBlockOk() (*BlockHeader, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *WithdrawalHeader) SetBlock(v BlockHeader)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *WithdrawalHeader) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetStatus

`func (o *WithdrawalHeader) GetStatus() TxStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WithdrawalHeader) GetStatusOk() (*TxStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WithdrawalHeader) SetStatus(v TxStatus)`

SetStatus sets Status field to given value.


### GetWithdrawalId

`func (o *WithdrawalHeader) GetWithdrawalId() string`

GetWithdrawalId returns the WithdrawalId field if non-nil, zero value otherwise.

### GetWithdrawalIdOk

`func (o *WithdrawalHeader) GetWithdrawalIdOk() (*string, bool)`

GetWithdrawalIdOk returns a tuple with the WithdrawalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalId

`func (o *WithdrawalHeader) SetWithdrawalId(v string)`

SetWithdrawalId sets WithdrawalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


