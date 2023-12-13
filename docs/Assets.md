# Assets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lovelace** | **int32** |  | 
**Tokens** | **map[string]map[string]int32** |  | 

## Methods

### NewAssets

`func NewAssets(lovelace int32, tokens map[string]map[string]int32, ) *Assets`

NewAssets instantiates a new Assets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsWithDefaults

`func NewAssetsWithDefaults() *Assets`

NewAssetsWithDefaults instantiates a new Assets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLovelace

`func (o *Assets) GetLovelace() int32`

GetLovelace returns the Lovelace field if non-nil, zero value otherwise.

### GetLovelaceOk

`func (o *Assets) GetLovelaceOk() (*int32, bool)`

GetLovelaceOk returns a tuple with the Lovelace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLovelace

`func (o *Assets) SetLovelace(v int32)`

SetLovelace sets Lovelace field to given value.


### GetTokens

`func (o *Assets) GetTokens() map[string]map[string]int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *Assets) GetTokensOk() (*map[string]map[string]int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *Assets) SetTokens(v map[string]map[string]int32)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


