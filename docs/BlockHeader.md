# BlockHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHeaderHash** | **string** |  | 
**BlockNo** | **int64** |  | 
**SlotNo** | **int64** |  | 

## Methods

### NewBlockHeader

`func NewBlockHeader(blockHeaderHash string, blockNo int64, slotNo int64, ) *BlockHeader`

NewBlockHeader instantiates a new BlockHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockHeaderWithDefaults

`func NewBlockHeaderWithDefaults() *BlockHeader`

NewBlockHeaderWithDefaults instantiates a new BlockHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHeaderHash

`func (o *BlockHeader) GetBlockHeaderHash() string`

GetBlockHeaderHash returns the BlockHeaderHash field if non-nil, zero value otherwise.

### GetBlockHeaderHashOk

`func (o *BlockHeader) GetBlockHeaderHashOk() (*string, bool)`

GetBlockHeaderHashOk returns a tuple with the BlockHeaderHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeaderHash

`func (o *BlockHeader) SetBlockHeaderHash(v string)`

SetBlockHeaderHash sets BlockHeaderHash field to given value.


### GetBlockNo

`func (o *BlockHeader) GetBlockNo() int64`

GetBlockNo returns the BlockNo field if non-nil, zero value otherwise.

### GetBlockNoOk

`func (o *BlockHeader) GetBlockNoOk() (*int64, bool)`

GetBlockNoOk returns a tuple with the BlockNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNo

`func (o *BlockHeader) SetBlockNo(v int64)`

SetBlockNo sets BlockNo field to given value.


### GetSlotNo

`func (o *BlockHeader) GetSlotNo() int64`

GetSlotNo returns the SlotNo field if non-nil, zero value otherwise.

### GetSlotNoOk

`func (o *BlockHeader) GetSlotNoOk() (*int64, bool)`

GetSlotNoOk returns a tuple with the SlotNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNo

`func (o *BlockHeader) SetSlotNo(v int64)`

SetSlotNo sets SlotNo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


