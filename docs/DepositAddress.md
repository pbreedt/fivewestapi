# DepositAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Memo** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDepositAddress

`func NewDepositAddress(address string, ) *DepositAddress`

NewDepositAddress instantiates a new DepositAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositAddressWithDefaults

`func NewDepositAddressWithDefaults() *DepositAddress`

NewDepositAddressWithDefaults instantiates a new DepositAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DepositAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DepositAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DepositAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMemo

`func (o *DepositAddress) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DepositAddress) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DepositAddress) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *DepositAddress) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *DepositAddress) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *DepositAddress) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


