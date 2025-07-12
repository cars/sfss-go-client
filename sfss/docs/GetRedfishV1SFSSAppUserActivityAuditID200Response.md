# GetRedfishV1SFSSAppUserActivityAuditID200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CDCInstance** | **string** | CDC instance identifier | 
**HTTPCode** | **string** | Response code that indicates whether the user activity is successfully completed or not  | 
**Operation** | **string** | Operation performed by the user  | 
**Payload** | **string** | Information provided by the user for the specific user activity | 
**SourceIP** | **string** | IP address of the system from which the specific user activity is performed | 
**TimeStamp** | **string** | Date and time at which the operation was performed; date in dd/mm/yyyy format and time in hh:mm:ss format | 
**URL** | **string** | RESTAPI URL that is used for the specific user activity | 
**UserAgent** | **string** |  | 
**UserName** | **string** | Username of the user who performed the operation | 
**UserRole** | **string** | Role of the user | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppUserActivityAuditID200Response

`func NewGetRedfishV1SFSSAppUserActivityAuditID200Response(cDCInstance string, hTTPCode string, operation string, payload string, sourceIP string, timeStamp string, uRL string, userAgent string, userName string, userRole string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppUserActivityAuditID200Response`

NewGetRedfishV1SFSSAppUserActivityAuditID200Response instantiates a new GetRedfishV1SFSSAppUserActivityAuditID200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppUserActivityAuditID200ResponseWithDefaults

`func NewGetRedfishV1SFSSAppUserActivityAuditID200ResponseWithDefaults() *GetRedfishV1SFSSAppUserActivityAuditID200Response`

NewGetRedfishV1SFSSAppUserActivityAuditID200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppUserActivityAuditID200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCDCInstance

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetCDCInstance() string`

GetCDCInstance returns the CDCInstance field if non-nil, zero value otherwise.

### GetCDCInstanceOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetCDCInstanceOk() (*string, bool)`

GetCDCInstanceOk returns a tuple with the CDCInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCInstance

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetCDCInstance(v string)`

SetCDCInstance sets CDCInstance field to given value.


### GetHTTPCode

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetHTTPCode() string`

GetHTTPCode returns the HTTPCode field if non-nil, zero value otherwise.

### GetHTTPCodeOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetHTTPCodeOk() (*string, bool)`

GetHTTPCodeOk returns a tuple with the HTTPCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPCode

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetHTTPCode(v string)`

SetHTTPCode sets HTTPCode field to given value.


### GetOperation

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetPayload

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetSourceIP

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetSourceIP() string`

GetSourceIP returns the SourceIP field if non-nil, zero value otherwise.

### GetSourceIPOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetSourceIPOk() (*string, bool)`

GetSourceIPOk returns a tuple with the SourceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIP

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetSourceIP(v string)`

SetSourceIP sets SourceIP field to given value.


### GetTimeStamp

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetURL

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetURL(v string)`

SetURL sets URL field to given value.


### GetUserAgent

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetUserName

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserRole

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppUserActivityAuditID200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


