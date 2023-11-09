# TokenMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]TokenMetadataFile**](TokenMetadataFile.md) |  | [optional] 
**Image** | **string** |  | 
**MediaType** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewTokenMetadata

`func NewTokenMetadata(image string, name string, ) *TokenMetadata`

NewTokenMetadata instantiates a new TokenMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenMetadataWithDefaults

`func NewTokenMetadataWithDefaults() *TokenMetadata`

NewTokenMetadataWithDefaults instantiates a new TokenMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TokenMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFiles

`func (o *TokenMetadata) GetFiles() []TokenMetadataFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TokenMetadata) GetFilesOk() (*[]TokenMetadataFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TokenMetadata) SetFiles(v []TokenMetadataFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TokenMetadata) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetImage

`func (o *TokenMetadata) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *TokenMetadata) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *TokenMetadata) SetImage(v string)`

SetImage sets Image field to given value.


### GetMediaType

`func (o *TokenMetadata) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *TokenMetadata) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *TokenMetadata) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *TokenMetadata) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetName

`func (o *TokenMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenMetadata) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


