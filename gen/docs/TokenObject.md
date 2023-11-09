# TokenObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencySymbol** | **string** |  | 
**TokenName** | **string** |  | 
**Ref** | **string** | An arbitrary text identifier for an object in a Marlowe object bundle. | 

## Methods

### NewTokenObject

`func NewTokenObject(currencySymbol string, tokenName string, ref string, ) *TokenObject`

NewTokenObject instantiates a new TokenObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenObjectWithDefaults

`func NewTokenObjectWithDefaults() *TokenObject`

NewTokenObjectWithDefaults instantiates a new TokenObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencySymbol

`func (o *TokenObject) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *TokenObject) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *TokenObject) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.


### GetTokenName

`func (o *TokenObject) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TokenObject) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TokenObject) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetRef

`func (o *TokenObject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *TokenObject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *TokenObject) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


