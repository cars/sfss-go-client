# DDCENUMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportAddressFamily** | **[]map[string]interface{}** | IP address family; possible values include IPv4 and IPv6 | 
**ConfigType** | **[]map[string]interface{}** | Configuration Type of DDC; possible values are Manual and KickStart | 
**FailureReason** | **[]map[string]interface{}** | Connection Failure Reason - Possible Values are NONE, Peer Closure,Dial Failure, KATO | 
**ConnectionStatus** | **[]map[string]interface{}** | Status of the TCP connection between the DDC and the CDC instance; possible values include Online/In-Progess:Connecting/Offline | 
**TransportType** | **[]map[string]interface{}** | Protocol used for communication with the DDC; possible value is TCP | 

## Methods

### NewDDCENUMS

`func NewDDCENUMS(transportAddressFamily []map[string]interface{}, configType []map[string]interface{}, failureReason []map[string]interface{}, connectionStatus []map[string]interface{}, transportType []map[string]interface{}, ) *DDCENUMS`

NewDDCENUMS instantiates a new DDCENUMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDDCENUMSWithDefaults

`func NewDDCENUMSWithDefaults() *DDCENUMS`

NewDDCENUMSWithDefaults instantiates a new DDCENUMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportAddressFamily

`func (o *DDCENUMS) GetTransportAddressFamily() []map[string]interface{}`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *DDCENUMS) GetTransportAddressFamilyOk() (*[]map[string]interface{}, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *DDCENUMS) SetTransportAddressFamily(v []map[string]interface{})`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetConfigType

`func (o *DDCENUMS) GetConfigType() []map[string]interface{}`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *DDCENUMS) GetConfigTypeOk() (*[]map[string]interface{}, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *DDCENUMS) SetConfigType(v []map[string]interface{})`

SetConfigType sets ConfigType field to given value.


### GetFailureReason

`func (o *DDCENUMS) GetFailureReason() []map[string]interface{}`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *DDCENUMS) GetFailureReasonOk() (*[]map[string]interface{}, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *DDCENUMS) SetFailureReason(v []map[string]interface{})`

SetFailureReason sets FailureReason field to given value.


### GetConnectionStatus

`func (o *DDCENUMS) GetConnectionStatus() []map[string]interface{}`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *DDCENUMS) GetConnectionStatusOk() (*[]map[string]interface{}, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *DDCENUMS) SetConnectionStatus(v []map[string]interface{})`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetTransportType

`func (o *DDCENUMS) GetTransportType() []map[string]interface{}`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *DDCENUMS) GetTransportTypeOk() (*[]map[string]interface{}, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *DDCENUMS) SetTransportType(v []map[string]interface{})`

SetTransportType sets TransportType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


