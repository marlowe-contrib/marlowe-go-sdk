# PlutusCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PubKeyCredential** | **string** |  | 
**ScriptCredential** | **string** |  | 

## Methods

### NewPlutusCredential

`func NewPlutusCredential(pubKeyCredential string, scriptCredential string, ) *PlutusCredential`

NewPlutusCredential instantiates a new PlutusCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlutusCredentialWithDefaults

`func NewPlutusCredentialWithDefaults() *PlutusCredential`

NewPlutusCredentialWithDefaults instantiates a new PlutusCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubKeyCredential

`func (o *PlutusCredential) GetPubKeyCredential() string`

GetPubKeyCredential returns the PubKeyCredential field if non-nil, zero value otherwise.

### GetPubKeyCredentialOk

`func (o *PlutusCredential) GetPubKeyCredentialOk() (*string, bool)`

GetPubKeyCredentialOk returns a tuple with the PubKeyCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKeyCredential

`func (o *PlutusCredential) SetPubKeyCredential(v string)`

SetPubKeyCredential sets PubKeyCredential field to given value.


### GetScriptCredential

`func (o *PlutusCredential) GetScriptCredential() string`

GetScriptCredential returns the ScriptCredential field if non-nil, zero value otherwise.

### GetScriptCredentialOk

`func (o *PlutusCredential) GetScriptCredentialOk() (*string, bool)`

GetScriptCredentialOk returns a tuple with the ScriptCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptCredential

`func (o *PlutusCredential) SetScriptCredential(v string)`

SetScriptCredential sets ScriptCredential field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


