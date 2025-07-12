# GetRedfishV1SFSSAppEnums200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **[]string** | Status of the add image operation; possible values include Failure, InProgress, Success, and NotStarted | 
**TransportType** | **[]string** | Transport type that is used to copy the SFSS image from the remote server; possible values include SCP, HTTP, HTTPS, and SFTP | 

## Methods

### NewGetRedfishV1SFSSAppEnums200Response

`func NewGetRedfishV1SFSSAppEnums200Response(status []string, transportType []string, ) *GetRedfishV1SFSSAppEnums200Response`

NewGetRedfishV1SFSSAppEnums200Response instantiates a new GetRedfishV1SFSSAppEnums200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppEnums200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppEnums200ResponseWithDefaults() *GetRedfishV1SFSSAppEnums200Response`

NewGetRedfishV1SFSSAppEnums200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppEnums200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetRedfishV1SFSSAppEnums200Response) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRedfishV1SFSSAppEnums200Response) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRedfishV1SFSSAppEnums200Response) SetStatus(v []string)`

SetStatus sets Status field to given value.


### GetTransportType

`func (o *GetRedfishV1SFSSAppEnums200Response) GetTransportType() []string`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *GetRedfishV1SFSSAppEnums200Response) GetTransportTypeOk() (*[]string, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *GetRedfishV1SFSSAppEnums200Response) SetTransportType(v []string)`

SetTransportType sets TransportType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


