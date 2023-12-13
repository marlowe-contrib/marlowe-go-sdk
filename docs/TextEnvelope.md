# TextEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CborHex** | **string** |  | 
**Description** | **string** |  | 
**Type** | **string** | What type of data is encoded in the CBOR Hex. Valid values include \&quot;Tx &lt;era&gt;\&quot;, \&quot;TxBody &lt;era&gt;\&quot;, and \&quot;ShelleyTxWitness &lt;era&gt;\&quot; where &lt;era&gt; is one of \&quot;BabbageEra\&quot;, \&quot;ConwayEra\&quot;. | 

## Methods

### NewTextEnvelope

`func NewTextEnvelope(cborHex string, description string, type_ string, ) *TextEnvelope`

NewTextEnvelope instantiates a new TextEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextEnvelopeWithDefaults

`func NewTextEnvelopeWithDefaults() *TextEnvelope`

NewTextEnvelopeWithDefaults instantiates a new TextEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCborHex

`func (o *TextEnvelope) GetCborHex() string`

GetCborHex returns the CborHex field if non-nil, zero value otherwise.

### GetCborHexOk

`func (o *TextEnvelope) GetCborHexOk() (*string, bool)`

GetCborHexOk returns a tuple with the CborHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCborHex

`func (o *TextEnvelope) SetCborHex(v string)`

SetCborHex sets CborHex field to given value.


### GetDescription

`func (o *TextEnvelope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TextEnvelope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TextEnvelope) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *TextEnvelope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextEnvelope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextEnvelope) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


