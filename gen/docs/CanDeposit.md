# CanDeposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanDeposit** | **int32** |  | 
**CaseIndex** | **int32** | Index of a \&quot;Case Action\&quot; in a \&quot;When\&quot; | 
**IntoAccount** | [**Party**](Party.md) |  | 
**IsMerkleizedContinuation** | **bool** | Indicates if a given contract continuation is merkleized | 
**OfToken** | [**Token**](Token.md) |  | 
**Party** | [**Party**](Party.md) |  | 

## Methods

### NewCanDeposit

`func NewCanDeposit(canDeposit int32, caseIndex int32, intoAccount Party, isMerkleizedContinuation bool, ofToken Token, party Party, ) *CanDeposit`

NewCanDeposit instantiates a new CanDeposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanDepositWithDefaults

`func NewCanDepositWithDefaults() *CanDeposit`

NewCanDepositWithDefaults instantiates a new CanDeposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanDeposit

`func (o *CanDeposit) GetCanDeposit() int32`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *CanDeposit) GetCanDepositOk() (*int32, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *CanDeposit) SetCanDeposit(v int32)`

SetCanDeposit sets CanDeposit field to given value.


### GetCaseIndex

`func (o *CanDeposit) GetCaseIndex() int32`

GetCaseIndex returns the CaseIndex field if non-nil, zero value otherwise.

### GetCaseIndexOk

`func (o *CanDeposit) GetCaseIndexOk() (*int32, bool)`

GetCaseIndexOk returns a tuple with the CaseIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseIndex

`func (o *CanDeposit) SetCaseIndex(v int32)`

SetCaseIndex sets CaseIndex field to given value.


### GetIntoAccount

`func (o *CanDeposit) GetIntoAccount() Party`

GetIntoAccount returns the IntoAccount field if non-nil, zero value otherwise.

### GetIntoAccountOk

`func (o *CanDeposit) GetIntoAccountOk() (*Party, bool)`

GetIntoAccountOk returns a tuple with the IntoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntoAccount

`func (o *CanDeposit) SetIntoAccount(v Party)`

SetIntoAccount sets IntoAccount field to given value.


### GetIsMerkleizedContinuation

`func (o *CanDeposit) GetIsMerkleizedContinuation() bool`

GetIsMerkleizedContinuation returns the IsMerkleizedContinuation field if non-nil, zero value otherwise.

### GetIsMerkleizedContinuationOk

`func (o *CanDeposit) GetIsMerkleizedContinuationOk() (*bool, bool)`

GetIsMerkleizedContinuationOk returns a tuple with the IsMerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerkleizedContinuation

`func (o *CanDeposit) SetIsMerkleizedContinuation(v bool)`

SetIsMerkleizedContinuation sets IsMerkleizedContinuation field to given value.


### GetOfToken

`func (o *CanDeposit) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *CanDeposit) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *CanDeposit) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetParty

`func (o *CanDeposit) GetParty() Party`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *CanDeposit) GetPartyOk() (*Party, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *CanDeposit) SetParty(v Party)`

SetParty sets Party field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


