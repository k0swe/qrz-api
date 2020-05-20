/*
 * QRZ XML Logbook Data Specification
 *
 * This document describes the interface specification for access to QRZ's XML subscription data service. This service provides real-time access to information from the QRZ.COM servers and databases. Access to this service requires user authentication through the use of a valid QRZ.COM username and password. While any QRZ user may login to the service, an active QRZ Logbook Data subscription is required to access most of its features. Non-subscriber access limits the data fields that are returned and is primarily intended for testing and troubleshooting purposes only. A description of subscription plans and rates is available on the [QRZ website](http://www.qrz.com/i/subscriptions.html).
 *
 * API version: 1.33
 * Contact: flloyd@qrz.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qrz

// Session struct for Session
type Session struct {
	// a valid user session key. A user session is established whenever a session key is returned. Any response from the server that does not contain the Key element indicates that no valid session exists and that a re-login is required to continue.
	Key string `json:"Key,omitempty" xml:"Key"`
	// Number of lookups performed by this user in the current 24 hour period
	Count float32 `json:"Count,omitempty" xml:"Count"`
	// time and date that the users subscription will expire - or - \"non-subscriber\"
	SubExp string `json:"SubExp,omitempty" xml:"SubExp"`
	// Time stamp for this message
	GMTime string `json:"GMTime,omitempty" xml:"GMTime"`
	// An informational message for the user
	Message string `json:"Message,omitempty" xml:"Message"`
	// XML system error message
	Error string `json:"Error,omitempty" xml:"Error"`
}
