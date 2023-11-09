# SafetyError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to [**Party**](Party.md) |  | [optional] 
**Address** | Pointer to [**PlutusAddress**](PlutusAddress.md) |  | [optional] 
**Bytes** | Pointer to **int32** |  | [optional] 
**ChoiceId** | Pointer to [**ChoiceId**](ChoiceId.md) |  | [optional] 
**Cost** | Pointer to [**ExBudget**](ExBudget.md) |  | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Error** | **string** |  | 
**Fatal** | **bool** |  | 
**Hash** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Token** | Pointer to [**Token**](Token.md) |  | [optional] 
**TokenName** | Pointer to **string** |  | [optional] 
**Transaction** | Pointer to [**Transaction**](Transaction.md) |  | [optional] 
**ValueId** | Pointer to **string** |  | [optional] 
**Warning** | Pointer to [**TransactionWarning**](TransactionWarning.md) |  | [optional] 

## Methods

### NewSafetyError

`func NewSafetyError(detail string, error_ string, fatal bool, ) *SafetyError`

NewSafetyError instantiates a new SafetyError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafetyErrorWithDefaults

`func NewSafetyErrorWithDefaults() *SafetyError`

NewSafetyErrorWithDefaults instantiates a new SafetyError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SafetyError) GetAccountId() Party`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SafetyError) GetAccountIdOk() (*Party, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SafetyError) SetAccountId(v Party)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SafetyError) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAddress

`func (o *SafetyError) GetAddress() PlutusAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SafetyError) GetAddressOk() (*PlutusAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SafetyError) SetAddress(v PlutusAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SafetyError) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBytes

`func (o *SafetyError) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *SafetyError) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *SafetyError) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *SafetyError) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetChoiceId

`func (o *SafetyError) GetChoiceId() ChoiceId`

GetChoiceId returns the ChoiceId field if non-nil, zero value otherwise.

### GetChoiceIdOk

`func (o *SafetyError) GetChoiceIdOk() (*ChoiceId, bool)`

GetChoiceIdOk returns a tuple with the ChoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceId

`func (o *SafetyError) SetChoiceId(v ChoiceId)`

SetChoiceId sets ChoiceId field to given value.

### HasChoiceId

`func (o *SafetyError) HasChoiceId() bool`

HasChoiceId returns a boolean if a field has been set.

### GetCost

`func (o *SafetyError) GetCost() ExBudget`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SafetyError) GetCostOk() (*ExBudget, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SafetyError) SetCost(v ExBudget)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SafetyError) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *SafetyError) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *SafetyError) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *SafetyError) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *SafetyError) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetDetail

`func (o *SafetyError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SafetyError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SafetyError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetError

`func (o *SafetyError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SafetyError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SafetyError) SetError(v string)`

SetError sets Error field to given value.


### GetFatal

`func (o *SafetyError) GetFatal() bool`

GetFatal returns the Fatal field if non-nil, zero value otherwise.

### GetFatalOk

`func (o *SafetyError) GetFatalOk() (*bool, bool)`

GetFatalOk returns a tuple with the Fatal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatal

`func (o *SafetyError) SetFatal(v bool)`

SetFatal sets Fatal field to given value.


### GetHash

`func (o *SafetyError) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SafetyError) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SafetyError) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SafetyError) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMessage

`func (o *SafetyError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SafetyError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SafetyError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SafetyError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRoleName

`func (o *SafetyError) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *SafetyError) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *SafetyError) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *SafetyError) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetToken

`func (o *SafetyError) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SafetyError) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SafetyError) SetToken(v Token)`

SetToken sets Token field to given value.

### HasToken

`func (o *SafetyError) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenName

`func (o *SafetyError) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *SafetyError) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *SafetyError) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *SafetyError) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetTransaction

`func (o *SafetyError) GetTransaction() Transaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *SafetyError) GetTransactionOk() (*Transaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *SafetyError) SetTransaction(v Transaction)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *SafetyError) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetValueId

`func (o *SafetyError) GetValueId() string`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *SafetyError) GetValueIdOk() (*string, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *SafetyError) SetValueId(v string)`

SetValueId sets ValueId field to given value.

### HasValueId

`func (o *SafetyError) HasValueId() bool`

HasValueId returns a boolean if a field has been set.

### GetWarning

`func (o *SafetyError) GetWarning() TransactionWarning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SafetyError) GetWarningOk() (*TransactionWarning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SafetyError) SetWarning(v TransactionWarning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *SafetyError) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


