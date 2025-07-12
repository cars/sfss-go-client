# GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CDCInstance** | **string** | CDC instance identifier | 
**HTTPCode** | **string** | Response code that indicates whether the user activity is successfully completed or not | 
**Operation** | **string** | Operation performed by the user | 
**Payload** | **string** | Information provided by the user for the specific user activity | 
**SourceIP** | **string** | IP address of the system from which the specific user activity was done  | 
**TimeStamp** | **string** | Date and time at which the user performed the operation; date in dd/mm/yyyy format and time in hh:mm:ss format | 
**URL** | **string** | RESTAPI URL that is used for the specific activity | 
**UserAgent** | **string** | API platform from which the specific user activity was triggered | 
**UserName** | **string** | Username of the user who performed the operation | 
**UserRole** | **string** | Role of the user | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner

`func NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner(cDCInstance string, hTTPCode string, operation string, payload string, sourceIP string, timeStamp string, uRL string, userAgent string, userName string, userRole string, odataId string, odataType string, odataContext string, ) *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner`

NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner instantiates a new GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInnerWithDefaults

`func NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInnerWithDefaults() *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner`

NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInnerWithDefaults instantiates a new GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCDCInstance

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetCDCInstance() string`

GetCDCInstance returns the CDCInstance field if non-nil, zero value otherwise.

### GetCDCInstanceOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetCDCInstanceOk() (*string, bool)`

GetCDCInstanceOk returns a tuple with the CDCInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCInstance

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetCDCInstance(v string)`

SetCDCInstance sets CDCInstance field to given value.


### GetHTTPCode

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetHTTPCode() string`

GetHTTPCode returns the HTTPCode field if non-nil, zero value otherwise.

### GetHTTPCodeOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetHTTPCodeOk() (*string, bool)`

GetHTTPCodeOk returns a tuple with the HTTPCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPCode

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetHTTPCode(v string)`

SetHTTPCode sets HTTPCode field to given value.


### GetOperation

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetPayload

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetSourceIP

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetSourceIP() string`

GetSourceIP returns the SourceIP field if non-nil, zero value otherwise.

### GetSourceIPOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetSourceIPOk() (*string, bool)`

GetSourceIPOk returns a tuple with the SourceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIP

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetSourceIP(v string)`

SetSourceIP sets SourceIP field to given value.


### GetTimeStamp

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetURL

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetURL(v string)`

SetURL sets URL field to given value.


### GetUserAgent

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetUserName

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserRole

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.


### GetOdataId

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


