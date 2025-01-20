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

// checks if the QRZDatabase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QRZDatabase{}

// QRZDatabase struct for QRZDatabase
type QRZDatabase struct {
	Session  *Session  `json:"Session,omitempty" xml:"Session"`
	Callsign *Callsign `json:"Callsign,omitempty" xml:"Callsign"`
	DXCC     *DXCC     `json:"DXCC,omitempty" xml:"DXCC"`
}

// NewQRZDatabase instantiates a new QRZDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQRZDatabase() *QRZDatabase {
	this := QRZDatabase{}
	return &this
}

// NewQRZDatabaseWithDefaults instantiates a new QRZDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQRZDatabaseWithDefaults() *QRZDatabase {
	this := QRZDatabase{}
	return &this
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *QRZDatabase) GetSession() Session {
	if o == nil || IsNil(o.Session) {
		var ret Session
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QRZDatabase) GetSessionOk() (*Session, bool) {
	if o == nil || IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *QRZDatabase) HasSession() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given Session and assigns it to the Session field.
func (o *QRZDatabase) SetSession(v Session) {
	o.Session = &v
}

// GetCallsign returns the Callsign field value if set, zero value otherwise.
func (o *QRZDatabase) GetCallsign() Callsign {
	if o == nil || IsNil(o.Callsign) {
		var ret Callsign
		return ret
	}
	return *o.Callsign
}

// GetCallsignOk returns a tuple with the Callsign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QRZDatabase) GetCallsignOk() (*Callsign, bool) {
	if o == nil || IsNil(o.Callsign) {
		return nil, false
	}
	return o.Callsign, true
}

// HasCallsign returns a boolean if a field has been set.
func (o *QRZDatabase) HasCallsign() bool {
	if o != nil && !IsNil(o.Callsign) {
		return true
	}

	return false
}

// SetCallsign gets a reference to the given Callsign and assigns it to the Callsign field.
func (o *QRZDatabase) SetCallsign(v Callsign) {
	o.Callsign = &v
}

// GetDXCC returns the DXCC field value if set, zero value otherwise.
func (o *QRZDatabase) GetDXCC() DXCC {
	if o == nil || IsNil(o.DXCC) {
		var ret DXCC
		return ret
	}
	return *o.DXCC
}

// GetDXCCOk returns a tuple with the DXCC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QRZDatabase) GetDXCCOk() (*DXCC, bool) {
	if o == nil || IsNil(o.DXCC) {
		return nil, false
	}
	return o.DXCC, true
}

// HasDXCC returns a boolean if a field has been set.
func (o *QRZDatabase) HasDXCC() bool {
	if o != nil && !IsNil(o.DXCC) {
		return true
	}

	return false
}

// SetDXCC gets a reference to the given DXCC and assigns it to the DXCC field.
func (o *QRZDatabase) SetDXCC(v DXCC) {
	o.DXCC = &v
}

func (o QRZDatabase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QRZDatabase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Session) {
		toSerialize["Session"] = o.Session
	}
	if !IsNil(o.Callsign) {
		toSerialize["Callsign"] = o.Callsign
	}
	if !IsNil(o.DXCC) {
		toSerialize["DXCC"] = o.DXCC
	}
	return toSerialize, nil
}

type NullableQRZDatabase struct {
	value *QRZDatabase
	isSet bool
}

func (v NullableQRZDatabase) Get() *QRZDatabase {
	return v.value
}

func (v *NullableQRZDatabase) Set(val *QRZDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableQRZDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableQRZDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQRZDatabase(val *QRZDatabase) *NullableQRZDatabase {
	return &NullableQRZDatabase{value: val, isSet: true}
}

func (v NullableQRZDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQRZDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
