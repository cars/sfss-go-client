# SFSSImagesGETENUMs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **[]interface{}** | Status of the add image operation; possible values inclulde Failure, InProgress, Success, and NotStarted  | 
**TransportType** | **[]interface{}** | Transport type used to copy the SFSS image from the remote repository; possible values include SCP, SFTP, HTTP, and HTTPs | 

## Methods

### NewSFSSImagesGETENUMs

`func NewSFSSImagesGETENUMs(status []interface{}, transportType []interface{}, ) *SFSSImagesGETENUMs`

NewSFSSImagesGETENUMs instantiates a new SFSSImagesGETENUMs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSFSSImagesGETENUMsWithDefaults

`func NewSFSSImagesGETENUMsWithDefaults() *SFSSImagesGETENUMs`

NewSFSSImagesGETENUMsWithDefaults instantiates a new SFSSImagesGETENUMs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SFSSImagesGETENUMs) GetStatus() []interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SFSSImagesGETENUMs) GetStatusOk() (*[]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SFSSImagesGETENUMs) SetStatus(v []interface{})`

SetStatus sets Status field to given value.


### GetTransportType

`func (o *SFSSImagesGETENUMs) GetTransportType() []interface{}`

GetTransportType returns the TransportType field if non-nil, zero value otherwise.

### GetTransportTypeOk

`func (o *SFSSImagesGETENUMs) GetTransportTypeOk() (*[]interface{}, bool)`

GetTransportTypeOk returns a tuple with the TransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportType

`func (o *SFSSImagesGETENUMs) SetTransportType(v []interface{})`

SetTransportType sets TransportType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


