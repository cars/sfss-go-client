# NTPServerGET

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NtpServer** | Pointer to **string** | IP address of the NTP server | [optional] 
**State** | Pointer to **string** | NTP connection state; possible values include Connected and Not Connected  | [optional] 
**Status** | Pointer to **string** | Status of the NTP service; possible values include Enable and Disable | [optional] 
**OdataId** | Pointer to **string** |  | [optional] 
**OdataType** | Pointer to **string** |  | [optional] 
**OdataContext** | Pointer to **string** |  | [optional] 

## Methods

### NewNTPServerGET

`func NewNTPServerGET() *NTPServerGET`

NewNTPServerGET instantiates a new NTPServerGET object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNTPServerGETWithDefaults

`func NewNTPServerGETWithDefaults() *NTPServerGET`

NewNTPServerGETWithDefaults instantiates a new NTPServerGET object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNtpServer

`func (o *NTPServerGET) GetNtpServer() string`

GetNtpServer returns the NtpServer field if non-nil, zero value otherwise.

### GetNtpServerOk

`func (o *NTPServerGET) GetNtpServerOk() (*string, bool)`

GetNtpServerOk returns a tuple with the NtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServer

`func (o *NTPServerGET) SetNtpServer(v string)`

SetNtpServer sets NtpServer field to given value.

### HasNtpServer

`func (o *NTPServerGET) HasNtpServer() bool`

HasNtpServer returns a boolean if a field has been set.

### GetState

`func (o *NTPServerGET) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NTPServerGET) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NTPServerGET) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NTPServerGET) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *NTPServerGET) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NTPServerGET) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NTPServerGET) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NTPServerGET) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOdataId

`func (o *NTPServerGET) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *NTPServerGET) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *NTPServerGET) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.

### HasOdataId

`func (o *NTPServerGET) HasOdataId() bool`

HasOdataId returns a boolean if a field has been set.

### GetOdataType

`func (o *NTPServerGET) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *NTPServerGET) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *NTPServerGET) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.

### HasOdataType

`func (o *NTPServerGET) HasOdataType() bool`

HasOdataType returns a boolean if a field has been set.

### GetOdataContext

`func (o *NTPServerGET) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *NTPServerGET) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *NTPServerGET) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.

### HasOdataContext

`func (o *NTPServerGET) HasOdataContext() bool`

HasOdataContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


