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

// checks if the BackupsGET type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupsGET{}

// BackupsGET This model lists all the backups taken from the SFSS VM.
type BackupsGET struct {
	// A set of backups obtained from SFSS
	Backups []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner `json:"Backups"`
	// Number of backups available
	BackupsodataCount float32 `json:"Backups@odata.count"`
	OdataId string `json:"@odata.id"`
	OdataContext string `json:"@odata.context"`
	OdataType string `json:"@odata.type"`
}

type _BackupsGET BackupsGET

// NewBackupsGET instantiates a new BackupsGET object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupsGET(backups []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, backupsodataCount float32, odataId string, odataContext string, odataType string) *BackupsGET {
	this := BackupsGET{}
	this.Backups = backups
	this.BackupsodataCount = backupsodataCount
	this.OdataId = odataId
	this.OdataContext = odataContext
	this.OdataType = odataType
	return &this
}

// NewBackupsGETWithDefaults instantiates a new BackupsGET object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupsGETWithDefaults() *BackupsGET {
	this := BackupsGET{}
	return &this
}

// GetBackups returns the Backups field value
func (o *BackupsGET) GetBackups() []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner {
	if o == nil {
		var ret []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner
		return ret
	}

	return o.Backups
}

// GetBackupsOk returns a tuple with the Backups field value
// and a boolean to check if the value has been set.
func (o *BackupsGET) GetBackupsOk() ([]GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Backups, true
}

// SetBackups sets field value
func (o *BackupsGET) SetBackups(v []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner) {
	o.Backups = v
}

// GetBackupsodataCount returns the BackupsodataCount field value
func (o *BackupsGET) GetBackupsodataCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BackupsodataCount
}

// GetBackupsodataCountOk returns a tuple with the BackupsodataCount field value
// and a boolean to check if the value has been set.
func (o *BackupsGET) GetBackupsodataCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupsodataCount, true
}

// SetBackupsodataCount sets field value
func (o *BackupsGET) SetBackupsodataCount(v float32) {
	o.BackupsodataCount = v
}

// GetOdataId returns the OdataId field value
func (o *BackupsGET) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *BackupsGET) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *BackupsGET) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataContext returns the OdataContext field value
func (o *BackupsGET) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *BackupsGET) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *BackupsGET) SetOdataContext(v string) {
	o.OdataContext = v
}

// GetOdataType returns the OdataType field value
func (o *BackupsGET) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *BackupsGET) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *BackupsGET) SetOdataType(v string) {
	o.OdataType = v
}

func (o BackupsGET) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupsGET) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Backups"] = o.Backups
	toSerialize["Backups@odata.count"] = o.BackupsodataCount
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.context"] = o.OdataContext
	toSerialize["@odata.type"] = o.OdataType
	return toSerialize, nil
}

func (o *BackupsGET) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Backups",
		"Backups@odata.count",
		"@odata.id",
		"@odata.context",
		"@odata.type",
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

	varBackupsGET := _BackupsGET{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBackupsGET)

	if err != nil {
		return err
	}

	*o = BackupsGET(varBackupsGET)

	return err
}

type NullableBackupsGET struct {
	value *BackupsGET
	isSet bool
}

func (v NullableBackupsGET) Get() *BackupsGET {
	return v.value
}

func (v *NullableBackupsGET) Set(val *BackupsGET) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupsGET) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupsGET) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupsGET(val *BackupsGET) *NullableBackupsGET {
	return &NullableBackupsGET{value: val, isSet: true}
}

func (v NullableBackupsGET) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupsGET) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


