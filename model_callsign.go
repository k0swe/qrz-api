/*
 * QRZ XML Logbook Data Specification
 *
 * This document describes the interface specification for access to QRZ's XML subscription data service. This service provides real-time access to information from the QRZ.COM servers and databases. Access to this service requires user authentication through the use of a valid QRZ.COM username and password. While any QRZ user may login to the service, an active QRZ Logbook Data subscription is required to access most of its features. Non-subscriber access limits the data fields that are returned and is primarily intended for testing and troubleshooting purposes only. A description of subscription plans and rates is available on the [QRZ website](http://www.qrz.com/i/subscriptions.html).
 *
 * API version: 1.34
 * Contact: va7stv@qrz.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qrz

// Callsign struct for Callsign
type Callsign struct {
	// callsign
	Call string `json:"call,omitempty" xml:"call"`
	// Cross reference: the query callsign that returned this record
	Xref string `json:"xref,omitempty" xml:"xref"`
	// Other callsigns that resolve to this record
	Aliases string `json:"aliases,omitempty" xml:"aliases"`
	// DXCC entity ID (country code) for the callsign
	Dxcc string `json:"dxcc,omitempty" xml:"dxcc"`
	// first name
	Fname string `json:"fname,omitempty" xml:"fname"`
	// last name
	Name string `json:"name,omitempty" xml:"name"`
	// address line 1 (i.e. house # and street)
	Addr1 string `json:"addr1,omitempty" xml:"addr1"`
	// address line 2 (i.e, city name)
	Addr2 string `json:"addr2,omitempty" xml:"addr2"`
	// state (USA Only)
	State string `json:"state,omitempty" xml:"state"`
	// Zip/postal code
	Zip string `json:"zip,omitempty" xml:"zip"`
	// country name for the QSL mailing address
	Country string `json:"country,omitempty" xml:"country"`
	// dxcc entity code for the mailing address country
	Ccode string `json:"ccode,omitempty" xml:"ccode"`
	// lattitude of address (signed decimal) S < 0 > N
	Lat float64 `json:"lat,omitempty" xml:"lat"`
	// longitude of address (signed decimal) W < 0 > E
	Lon float64 `json:"lon,omitempty" xml:"lon"`
	// Maidenhead grid locator
	Grid string `json:"grid,omitempty" xml:"grid"`
	// county name (USA)
	County string `json:"county,omitempty" xml:"county"`
	// FIPS county identifier (USA)
	Fips string `json:"fips,omitempty" xml:"fips"`
	// DXCC country name of the callsign
	Land string `json:"land,omitempty" xml:"land"`
	// license effective date (USA)
	Efdate string `json:"efdate,omitempty" xml:"efdate"`
	// license expiration date (USA)
	Expdate string `json:"expdate,omitempty" xml:"expdate"`
	// previous callsign
	PCall string `json:"p_call,omitempty" xml:"p_call"`
	// license class
	Class string `json:"class,omitempty" xml:"class"`
	// license type codes (USA)
	Codes string `json:"codes,omitempty" xml:"codes"`
	// QSL manager info
	Qslmgr string `json:"qslmgr,omitempty" xml:"qslmgr"`
	// email address
	Email string `json:"email,omitempty" xml:"email"`
	// web page address
	Url string `json:"url,omitempty" xml:"url"`
	// QRZ web page views
	UViews float32 `json:"u_views,omitempty" xml:"u_views"`
	// approximate length of the bio HTML in bytes
	Bio string `json:"bio,omitempty" xml:"bio"`
	// date of the last bio update
	Biodate string `json:"biodate,omitempty" xml:"biodate"`
	// full URL of the callsign's primary image
	Image string `json:"image,omitempty" xml:"image"`
	// height:width:size in bytes, of the image file
	Imageinfo string `json:"imageinfo,omitempty" xml:"imageinfo"`
	// QRZ db serial number
	Serial string `json:"serial,omitempty" xml:"serial"`
	// QRZ callsign last modified date
	Moddate string `json:"moddate,omitempty" xml:"moddate"`
	// Metro Service Area (USPS)
	MSA string `json:"MSA,omitempty" xml:"MSA"`
	// Telephone Area Code (USA)
	AreaCode string `json:"AreaCode,omitempty" xml:"AreaCode"`
	// Time Zone (USA)
	TimeZone string `json:"TimeZone,omitempty" xml:"TimeZone"`
	// GMT Time Offset
	GMTOffset string `json:"GMTOffset,omitempty" xml:"GMTOffset"`
	// Daylight Saving Time Observed
	DST string `json:"DST,omitempty" xml:"DST"`
	// Will accept e-qsl (0/1 or blank if unknown)
	Eqsl string `json:"eqsl,omitempty" xml:"eqsl"`
	// Will return paper QSL (0/1 or blank if unknown)
	Mqsl string `json:"mqsl,omitempty" xml:"mqsl"`
	// CQ Zone identifier
	Cqzone string `json:"cqzone,omitempty" xml:"cqzone"`
	// ITU Zone identifier
	Ituzone string `json:"ituzone,omitempty" xml:"ituzone"`
	// operator's year of birth
	Born float32 `json:"born,omitempty" xml:"born"`
	// User who manages this callsign on QRZ
	User string `json:"user,omitempty" xml:"user"`
	// Will accept LOTW (0/1 or blank if unknown)
	Lotw string `json:"lotw,omitempty" xml:"lotw"`
	// IOTA Designator (blank if unknown)
	Iota string `json:"iota,omitempty" xml:"iota"`
	// Describes source of lat/long data
	Geoloc string `json:"geoloc,omitempty" xml:"geoloc"`
	// Attention address line, this line should be prepended to the address
	Attn string `json:"attn,omitempty" xml:"attn"`
	// A different or shortened name used on the air
	Nickname string `json:"nickname,omitempty" xml:"nickname"`
	// Combined full name and nickname in the format used by QRZ. This fortmat is subject to change.
	NameFmt string `json:"name_fmt,omitempty" xml:"name_fmt"`
}
