/*
SmartFabric Storage Software Application REST APIs

REST APIs used for managing SFSS application are captured in this section.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sfss

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner{}

// GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner struct for GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner
type GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner struct {
	// CDC instance identifier
	CDCInstance string `json:"CDCInstance"`
	// Response code that indicates whether the user activity is successfully completed or not
	HTTPCode string `json:"HTTPCode"`
	// Operation performed by the user
	Operation string `json:"Operation"`
	// Information provided by the user for the specific user activity
	Payload string `json:"Payload"`
	// IP address of the system from which the specific user activity was done 
	SourceIP string `json:"SourceIP"`
	// Date and time at which the user performed the operation; date in dd/mm/yyyy format and time in hh:mm:ss format
	TimeStamp string `json:"TimeStamp"`
	// RESTAPI URL that is used for the specific activity
	URL string `json:"URL"`
	// API platform from which the specific user activity was triggered
	UserAgent string `json:"UserAgent"`
	// Username of the user who performed the operation
	UserName string `json:"UserName"`
	// Role of the user
	UserRole string `json:"UserRole"`
	OdataId string `json:"@odata.id"`
	OdataType string `json:"@odata.type"`
	OdataContext string `json:"@odata.context"`
}

type _GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner

// NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner instantiates a new GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner(cDCInstance string, hTTPCode string, operation string, payload string, sourceIP string, timeStamp string, uRL string, userAgent string, userName string, userRole string, odataId string, odataType string, odataContext string) *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner {
	this := GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner{}
	this.CDCInstance = cDCInstance
	this.HTTPCode = hTTPCode
	this.Operation = operation
	this.Payload = payload
	this.SourceIP = sourceIP
	this.TimeStamp = timeStamp
	this.URL = uRL
	this.UserAgent = userAgent
	this.UserName = userName
	this.UserRole = userRole
	this.OdataId = odataId
	this.OdataType = odataType
	this.OdataContext = odataContext
	return &this
}

// NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInnerWithDefaults instantiates a new GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInnerWithDefaults() *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner {
	this := GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner{}
	return &this
}

// GetCDCInstance returns the CDCInstance field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetCDCInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CDCInstance
}

// GetCDCInstanceOk returns a tuple with the CDCInstance field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetCDCInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CDCInstance, true
}

// SetCDCInstance sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetCDCInstance(v string) {
	o.CDCInstance = v
}

// GetHTTPCode returns the HTTPCode field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetHTTPCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HTTPCode
}

// GetHTTPCodeOk returns a tuple with the HTTPCode field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetHTTPCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HTTPCode, true
}

// SetHTTPCode sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetHTTPCode(v string) {
	o.HTTPCode = v
}

// GetOperation returns the Operation field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetOperation(v string) {
	o.Operation = v
}

// GetPayload returns the Payload field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetPayload(v string) {
	o.Payload = v
}

// GetSourceIP returns the SourceIP field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetSourceIP() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceIP
}

// GetSourceIPOk returns a tuple with the SourceIP field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetSourceIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceIP, true
}

// SetSourceIP sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetSourceIP(v string) {
	o.SourceIP = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetTimeStamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetTimeStampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetTimeStamp(v string) {
	o.TimeStamp = v
}

// GetURL returns the URL field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.URL
}

// GetURLOk returns a tuple with the URL field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.URL, true
}

// SetURL sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetURL(v string) {
	o.URL = v
}

// GetUserAgent returns the UserAgent field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetUserAgent(v string) {
	o.UserAgent = v
}

// GetUserName returns the UserName field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetUserName(v string) {
	o.UserName = v
}

// GetUserRole returns the UserRole field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserRole
}

// GetUserRoleOk returns a tuple with the UserRole field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetUserRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserRole, true
}

// SetUserRole sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetUserRole(v string) {
	o.UserRole = v
}

// GetOdataId returns the OdataId field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataType returns the OdataType field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataContext returns the OdataContext field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) SetOdataContext(v string) {
	o.OdataContext = v
}

func (o GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CDCInstance"] = o.CDCInstance
	toSerialize["HTTPCode"] = o.HTTPCode
	toSerialize["Operation"] = o.Operation
	toSerialize["Payload"] = o.Payload
	toSerialize["SourceIP"] = o.SourceIP
	toSerialize["TimeStamp"] = o.TimeStamp
	toSerialize["URL"] = o.URL
	toSerialize["UserAgent"] = o.UserAgent
	toSerialize["UserName"] = o.UserName
	toSerialize["UserRole"] = o.UserRole
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.type"] = o.OdataType
	toSerialize["@odata.context"] = o.OdataContext
	return toSerialize, nil
}

func (o *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CDCInstance",
		"HTTPCode",
		"Operation",
		"Payload",
		"SourceIP",
		"TimeStamp",
		"URL",
		"UserAgent",
		"UserName",
		"UserRole",
		"@odata.id",
		"@odata.type",
		"@odata.context",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner := _GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner)

	if err != nil {
		return err
	}

	*o = GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner(varGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner)

	return err
}

type NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner struct {
	value *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner
	isSet bool
}

func (v NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) Get() *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner {
	return v.value
}

func (v *NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) Set(val *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner(val *GetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) *NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner {
	return &NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner{value: val, isSet: true}
}

func (v NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRedfishV1SFSSAppUserActivityAudit200ResponseUserActivityAuditsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


