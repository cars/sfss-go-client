# NTPServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the NTP service | [optional] 

## Methods

### NewNTPServerConfig

`func NewNTPServerConfig() *NTPServerConfig`

NewNTPServerConfig instantiates a new NTPServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNTPServerConfigWithDefaults

`func NewNTPServerConfigWithDefaults() *NTPServerConfig`

NewNTPServerConfigWithDefaults instantiates a new NTPServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NTPServerConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NTPServerConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NTPServerConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NTPServerConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


