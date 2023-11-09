# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | [**Contract**](Contract.md) |  | 
**Input** | [**TransactionInput**](TransactionInput.md) |  | 
**Output** | [**TransactionOutput**](TransactionOutput.md) |  | 
**State** | [**MarloweState**](MarloweState.md) |  | 

## Methods

### NewTransaction

`func NewTransaction(contract Contract, input TransactionInput, output TransactionOutput, state MarloweState, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *Transaction) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *Transaction) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *Transaction) SetContract(v Contract)`

SetContract sets Contract field to given value.


### GetInput

`func (o *Transaction) GetInput() TransactionInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Transaction) GetInputOk() (*TransactionInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Transaction) SetInput(v TransactionInput)`

SetInput sets Input field to given value.


### GetOutput

`func (o *Transaction) GetOutput() TransactionOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Transaction) GetOutputOk() (*TransactionOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Transaction) SetOutput(v TransactionOutput)`

SetOutput sets Output field to given value.


### GetState

`func (o *Transaction) GetState() MarloweState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Transaction) GetStateOk() (*MarloweState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Transaction) SetState(v MarloweState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


