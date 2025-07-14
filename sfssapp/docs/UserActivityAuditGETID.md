# UserActivityAuditGETID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CDCInstance** | **string** | CDC instance identifier | 
**HTTPCode** | **string** | Response code that indicates whether the user activity is successfully completed or not | 
**Operation** | **string** | Operation performed by the user | 
**Payload** | **string** | Information provided by the user for the specific user activity | 
**SourceIP** | **string** | IP address of the system from which the specific user activity was done  | 
**TimeStamp** | **string** | Date and time at which the operation was performed; the date is of the mm/dd/yyyy format and time is of hh:mm:ss format | 
**URL** | **string** | RESTAPI URL that is used for the specific user activity | 
**UserAgent** | **string** |  | 
**UserName** | **string** | Username of the user who performed the operation  | 
**UserRole** | **string** | Role of the user | 
**OdataId** | **string** |  | 
**OdataType** | **string** |  | 
**OdataContext** | **string** |  | 

## Methods

### NewUserActivityAuditGETID

`func NewUserActivityAuditGETID(cDCInstance string, hTTPCode string, operation string, payload string, sourceIP string, timeStamp string, uRL string, userAgent string, userName string, userRole string, odataId string, odataType string, odataContext string, ) *UserActivityAuditGETID`

NewUserActivityAuditGETID instantiates a new UserActivityAuditGETID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityAuditGETIDWithDefaults

`func NewUserActivityAuditGETIDWithDefaults() *UserActivityAuditGETID`

NewUserActivityAuditGETIDWithDefaults instantiates a new UserActivityAuditGETID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCDCInstance

`func (o *UserActivityAuditGETID) GetCDCInstance() string`

GetCDCInstance returns the CDCInstance field if non-nil, zero value otherwise.

### GetCDCInstanceOk

`func (o *UserActivityAuditGETID) GetCDCInstanceOk() (*string, bool)`

GetCDCInstanceOk returns a tuple with the CDCInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCDCInstance

`func (o *UserActivityAuditGETID) SetCDCInstance(v string)`

SetCDCInstance sets CDCInstance field to given value.


### GetHTTPCode

`func (o *UserActivityAuditGETID) GetHTTPCode() string`

GetHTTPCode returns the HTTPCode field if non-nil, zero value otherwise.

### GetHTTPCodeOk

`func (o *UserActivityAuditGETID) GetHTTPCodeOk() (*string, bool)`

GetHTTPCodeOk returns a tuple with the HTTPCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPCode

`func (o *UserActivityAuditGETID) SetHTTPCode(v string)`

SetHTTPCode sets HTTPCode field to given value.


### GetOperation

`func (o *UserActivityAuditGETID) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *UserActivityAuditGETID) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *UserActivityAuditGETID) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetPayload

`func (o *UserActivityAuditGETID) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UserActivityAuditGETID) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UserActivityAuditGETID) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetSourceIP

`func (o *UserActivityAuditGETID) GetSourceIP() string`

GetSourceIP returns the SourceIP field if non-nil, zero value otherwise.

### GetSourceIPOk

`func (o *UserActivityAuditGETID) GetSourceIPOk() (*string, bool)`

GetSourceIPOk returns a tuple with the SourceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIP

`func (o *UserActivityAuditGETID) SetSourceIP(v string)`

SetSourceIP sets SourceIP field to given value.


### GetTimeStamp

`func (o *UserActivityAuditGETID) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *UserActivityAuditGETID) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *UserActivityAuditGETID) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.


### GetURL

`func (o *UserActivityAuditGETID) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *UserActivityAuditGETID) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *UserActivityAuditGETID) SetURL(v string)`

SetURL sets URL field to given value.


### GetUserAgent

`func (o *UserActivityAuditGETID) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *UserActivityAuditGETID) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *UserActivityAuditGETID) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetUserName

`func (o *UserActivityAuditGETID) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserActivityAuditGETID) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserActivityAuditGETID) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserRole

`func (o *UserActivityAuditGETID) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *UserActivityAuditGETID) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *UserActivityAuditGETID) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.


### GetOdataId

`func (o *UserActivityAuditGETID) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *UserActivityAuditGETID) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *UserActivityAuditGETID) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *UserActivityAuditGETID) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *UserActivityAuditGETID) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *UserActivityAuditGETID) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataContext

`func (o *UserActivityAuditGETID) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *UserActivityAuditGETID) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *UserActivityAuditGETID) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


