/*
QRZ XML Logbook Data Specification

This document describes the interface specification for access to QRZ's XML subscription data service. This service provides real-time access to information from the QRZ.COM servers and databases. Access to this service requires user authentication through the use of a valid QRZ.COM username and password. While any QRZ user may login to the service, an active QRZ Logbook Data subscription is required to access most of its features. Non-subscriber access limits the data fields that are returned and is primarily intended for testing and troubleshooting purposes only. A description of subscription plans and rates is available on the [QRZ website](http://www.qrz.com/i/subscriptions.html).

API version: 1.34
Contact: va7stv@qrz.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qrz

import (
	"encoding/json"
)

// checks if the DXCC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DXCC{}

// DXCC struct for DXCC
type DXCC struct {
	// DXCC entity number for this record
	Dxcc *string `json:"dxcc,omitempty" xml:"dxcc"`
	// 2-letter country code (ISO-3166)
	Cc *string `json:"cc,omitempty" xml:"cc"`
	// 3-letter country code (ISO-3166)
	Ccc *string `json:"ccc,omitempty" xml:"ccc"`
	// long name
	Name *string `json:"name,omitempty" xml:"name"`
	// 2-letter continent designator
	Continent *string `json:"continent,omitempty" xml:"continent"`
	// ITU Zone
	Ituzone *string `json:"ituzone,omitempty" xml:"ituzone"`
	// CQ Zone
	Cqzone *string `json:"cqzone,omitempty" xml:"cqzone"`
	// timezone offset +/- UTC in hours. Odd timezones, such as 0545 mean \"5 hours, 45 minutes\". The plus (+) sign is implied.
	Timezone *int32 `json:"timezone,omitempty" xml:"timezone"`
	// Latitude (approx.)
	Lat *float64 `json:"lat,omitempty" xml:"lat"`
	// Longitude (approx.)
	Lon *float64 `json:"lon,omitempty" xml:"lon"`
	// Special notes and/or exceptions
	Notes *string `json:"notes,omitempty" xml:"notes"`
}

// NewDXCC instantiates a new DXCC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDXCC() *DXCC {
	this := DXCC{}
	return &this
}

// NewDXCCWithDefaults instantiates a new DXCC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDXCCWithDefaults() *DXCC {
	this := DXCC{}
	return &this
}

// GetDxcc returns the Dxcc field value if set, zero value otherwise.
func (o *DXCC) GetDxcc() string {
	if o == nil || IsNil(o.Dxcc) {
		var ret string
		return ret
	}
	return *o.Dxcc
}

// GetDxccOk returns a tuple with the Dxcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetDxccOk() (*string, bool) {
	if o == nil || IsNil(o.Dxcc) {
		return nil, false
	}
	return o.Dxcc, true
}

// HasDxcc returns a boolean if a field has been set.
func (o *DXCC) HasDxcc() bool {
	if o != nil && !IsNil(o.Dxcc) {
		return true
	}

	return false
}

// SetDxcc gets a reference to the given string and assigns it to the Dxcc field.
func (o *DXCC) SetDxcc(v string) {
	o.Dxcc = &v
}

// GetCc returns the Cc field value if set, zero value otherwise.
func (o *DXCC) GetCc() string {
	if o == nil || IsNil(o.Cc) {
		var ret string
		return ret
	}
	return *o.Cc
}

// GetCcOk returns a tuple with the Cc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetCcOk() (*string, bool) {
	if o == nil || IsNil(o.Cc) {
		return nil, false
	}
	return o.Cc, true
}

// HasCc returns a boolean if a field has been set.
func (o *DXCC) HasCc() bool {
	if o != nil && !IsNil(o.Cc) {
		return true
	}

	return false
}

// SetCc gets a reference to the given string and assigns it to the Cc field.
func (o *DXCC) SetCc(v string) {
	o.Cc = &v
}

// GetCcc returns the Ccc field value if set, zero value otherwise.
func (o *DXCC) GetCcc() string {
	if o == nil || IsNil(o.Ccc) {
		var ret string
		return ret
	}
	return *o.Ccc
}

// GetCccOk returns a tuple with the Ccc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetCccOk() (*string, bool) {
	if o == nil || IsNil(o.Ccc) {
		return nil, false
	}
	return o.Ccc, true
}

// HasCcc returns a boolean if a field has been set.
func (o *DXCC) HasCcc() bool {
	if o != nil && !IsNil(o.Ccc) {
		return true
	}

	return false
}

// SetCcc gets a reference to the given string and assigns it to the Ccc field.
func (o *DXCC) SetCcc(v string) {
	o.Ccc = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DXCC) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DXCC) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DXCC) SetName(v string) {
	o.Name = &v
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *DXCC) GetContinent() string {
	if o == nil || IsNil(o.Continent) {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetContinentOk() (*string, bool) {
	if o == nil || IsNil(o.Continent) {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *DXCC) HasContinent() bool {
	if o != nil && !IsNil(o.Continent) {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *DXCC) SetContinent(v string) {
	o.Continent = &v
}

// GetItuzone returns the Ituzone field value if set, zero value otherwise.
func (o *DXCC) GetItuzone() string {
	if o == nil || IsNil(o.Ituzone) {
		var ret string
		return ret
	}
	return *o.Ituzone
}

// GetItuzoneOk returns a tuple with the Ituzone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetItuzoneOk() (*string, bool) {
	if o == nil || IsNil(o.Ituzone) {
		return nil, false
	}
	return o.Ituzone, true
}

// HasItuzone returns a boolean if a field has been set.
func (o *DXCC) HasItuzone() bool {
	if o != nil && !IsNil(o.Ituzone) {
		return true
	}

	return false
}

// SetItuzone gets a reference to the given string and assigns it to the Ituzone field.
func (o *DXCC) SetItuzone(v string) {
	o.Ituzone = &v
}

// GetCqzone returns the Cqzone field value if set, zero value otherwise.
func (o *DXCC) GetCqzone() string {
	if o == nil || IsNil(o.Cqzone) {
		var ret string
		return ret
	}
	return *o.Cqzone
}

// GetCqzoneOk returns a tuple with the Cqzone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetCqzoneOk() (*string, bool) {
	if o == nil || IsNil(o.Cqzone) {
		return nil, false
	}
	return o.Cqzone, true
}

// HasCqzone returns a boolean if a field has been set.
func (o *DXCC) HasCqzone() bool {
	if o != nil && !IsNil(o.Cqzone) {
		return true
	}

	return false
}

// SetCqzone gets a reference to the given string and assigns it to the Cqzone field.
func (o *DXCC) SetCqzone(v string) {
	o.Cqzone = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *DXCC) GetTimezone() int32 {
	if o == nil || IsNil(o.Timezone) {
		var ret int32
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetTimezoneOk() (*int32, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *DXCC) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given int32 and assigns it to the Timezone field.
func (o *DXCC) SetTimezone(v int32) {
	o.Timezone = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *DXCC) GetLat() float64 {
	if o == nil || IsNil(o.Lat) {
		var ret float64
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetLatOk() (*float64, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *DXCC) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float64 and assigns it to the Lat field.
func (o *DXCC) SetLat(v float64) {
	o.Lat = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *DXCC) GetLon() float64 {
	if o == nil || IsNil(o.Lon) {
		var ret float64
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetLonOk() (*float64, bool) {
	if o == nil || IsNil(o.Lon) {
		return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *DXCC) HasLon() bool {
	if o != nil && !IsNil(o.Lon) {
		return true
	}

	return false
}

// SetLon gets a reference to the given float64 and assigns it to the Lon field.
func (o *DXCC) SetLon(v float64) {
	o.Lon = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *DXCC) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DXCC) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *DXCC) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *DXCC) SetNotes(v string) {
	o.Notes = &v
}

func (o DXCC) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DXCC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dxcc) {
		toSerialize["dxcc"] = o.Dxcc
	}
	if !IsNil(o.Cc) {
		toSerialize["cc"] = o.Cc
	}
	if !IsNil(o.Ccc) {
		toSerialize["ccc"] = o.Ccc
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Continent) {
		toSerialize["continent"] = o.Continent
	}
	if !IsNil(o.Ituzone) {
		toSerialize["ituzone"] = o.Ituzone
	}
	if !IsNil(o.Cqzone) {
		toSerialize["cqzone"] = o.Cqzone
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lon) {
		toSerialize["lon"] = o.Lon
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableDXCC struct {
	value *DXCC
	isSet bool
}

func (v NullableDXCC) Get() *DXCC {
	return v.value
}

func (v *NullableDXCC) Set(val *DXCC) {
	v.value = val
	v.isSet = true
}

func (v NullableDXCC) IsSet() bool {
	return v.isSet
}

func (v *NullableDXCC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDXCC(val *DXCC) *NullableDXCC {
	return &NullableDXCC{value: val, isSet: true}
}

func (v NullableDXCC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDXCC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
