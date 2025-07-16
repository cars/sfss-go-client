# DDCGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Identity of the Entity - Fully Qualified Name (FQN) of the Direct Discovery Controller(DDC) | 
**TransportAddress** | **string** | IP address of the DDC | 
**TransportAddressFamily** | **string** | IP address family; possible values include IPv4 and IPv6 | 
**PortId** | **float32** | Port on which the DDC listens for CDC to connect to | 
**TransportType** | **string** | Protocol used for communication with the DDC; possible value is TCP | 
**Activate** | **bool** | Activation of DDC that initiates a pull registration from CDC | 
**ConfigType** | **string** | Configuration Type of DDC; possible values are Manual and KickStart | 
**ConnectionStatus** | **string** | Status of the TCP connection between the DDC and the CDC instance. Possible values include Online/In-Progess:Connecting/Offline | 
**FailureReason** | **string** | Reason for the DDC being offline | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewDDCGETID

`func NewDDCGETID(id string, transportAddress string, transportAddressFamily string, portId float32, transportType string, activate bool, configType string, connectionStatus string, failureReason string, odataId string, odataType string, odataContext string, ) *DDCGETID`

NewDDCGETID instantiates a new DDCGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDDCGETIDWithDefaults

`func NewDDCGETIDWithDefaults() *DDCGETID`

NewDDCGETIDWithDefaults instantiates a new DDCGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DDCGETID) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DDCGETID) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DDCGETID) SetId(v string)`

SetId sets Id field to given value.


### GetTransportAddress

`func (o *DDCGETID) GetTransportAddress() string`

GetTransportAddress returns the TransportAddress field if non-nil, zero value otherwise.

### GetTransportAddressOk

`func (o *DDCGETID) GetTransportAddressOk() (*string, bool)`

GetTransportAddressOk returns a tuple with the TransportAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddress

`func (o *DDCGETID) SetTransportAddress(v string)`

SetTransportAddress sets TransportAddress field to given value.


### GetTransportAddressFamily

`func (o *DDCGETID) GetTransportAddressFamily() string`

GetTransportAddressFamily returns the TransportAddressFamily field if non-nil, zero value otherwise.

### GetTransportAddressFamilyOk

`func (o *DDCGETID) GetTransportAddressFamilyOk() (*string, bool)`

GetTransportAddressFamilyOk returns a tuple with the TransportAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportAddressFamily

`func (o *DDCGETID) SetTransportAddressFamily(v string)`

SetTransportAddressFamily sets TransportAddressFamily field to given value.


### GetPortId

`func (o *DDCGETID) GetPortId() float32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *DDCGETID) GetPortIdOk() (*float32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *DDCGETID) SetPortId(v float32)`

SetPortId sets PortId field to given value.


### GetTransportType

`func (o *DDCGETID) GetTransportType() string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *DDCGETID) GetTransportTypeOk() (*string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *DDCGETID) SetTransportType(v string)`

SetTransportType sets TransportType field to given value.


### GetActivate

`func (o *DDCGETID) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *DDCGETID) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *DDCGETID) SetActivate(v bool)`

SetActivate sets Activate field to given value.


### GetConfigType

`func (o *DDCGETID) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *DDCGETID) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *DDCGETID) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetConnectionStatus

`func (o *DDCGETID) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *DDCGETID) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *DDCGETID) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.


### GetFailureReason

`func (o *DDCGETID) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *DDCGETID) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *DDCGETID) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### GetOdataId

`func (o *DDCGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *DDCGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *DDCGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *DDCGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *DDCGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *DDCGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *DDCGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *DDCGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *DDCGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


