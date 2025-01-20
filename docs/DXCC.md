# DXCC

## Properties

| Name          | Type                   | Description                                                                                                                         | Notes      |
| ------------- | ---------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Dxcc**      | Pointer to **string**  | DXCC entity number for this record                                                                                                  | [optional] |
| **Cc**        | Pointer to **string**  | 2-letter country code (ISO-3166)                                                                                                    | [optional] |
| **Ccc**       | Pointer to **string**  | 3-letter country code (ISO-3166)                                                                                                    | [optional] |
| **Name**      | Pointer to **string**  | long name                                                                                                                           | [optional] |
| **Continent** | Pointer to **string**  | 2-letter continent designator                                                                                                       | [optional] |
| **Ituzone**   | Pointer to **string**  | ITU Zone                                                                                                                            | [optional] |
| **Cqzone**    | Pointer to **string**  | CQ Zone                                                                                                                             | [optional] |
| **Timezone**  | Pointer to **int32**   | timezone offset +/- UTC in hours. Odd timezones, such as 0545 mean \&quot;5 hours, 45 minutes\&quot;. The plus (+) sign is implied. | [optional] |
| **Lat**       | Pointer to **float64** | Latitude (approx.)                                                                                                                  | [optional] |
| **Lon**       | Pointer to **float64** | Longitude (approx.)                                                                                                                 | [optional] |
| **Notes**     | Pointer to **string**  | Special notes and/or exceptions                                                                                                     | [optional] |

## Methods

### NewDXCC

`func NewDXCC() *DXCC`

NewDXCC instantiates a new DXCC object This constructor will assign default
values to properties that have it defined, and makes sure properties required by
API are set, but the set of arguments will change when the set of required
properties is changed

### NewDXCCWithDefaults

`func NewDXCCWithDefaults() *DXCC`

NewDXCCWithDefaults instantiates a new DXCC object This constructor will only
assign default values to properties that have it defined, but it doesn't
guarantee that properties required by API are set

### GetDxcc

`func (o *DXCC) GetDxcc() string`

GetDxcc returns the Dxcc field if non-nil, zero value otherwise.

### GetDxccOk

`func (o *DXCC) GetDxccOk() (*string, bool)`

GetDxccOk returns a tuple with the Dxcc field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetDxcc

`func (o *DXCC) SetDxcc(v string)`

SetDxcc sets Dxcc field to given value.

### HasDxcc

`func (o *DXCC) HasDxcc() bool`

HasDxcc returns a boolean if a field has been set.

### GetCc

`func (o *DXCC) GetCc() string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *DXCC) GetCcOk() (*string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *DXCC) SetCc(v string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *DXCC) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetCcc

`func (o *DXCC) GetCcc() string`

GetCcc returns the Ccc field if non-nil, zero value otherwise.

### GetCccOk

`func (o *DXCC) GetCccOk() (*string, bool)`

GetCccOk returns a tuple with the Ccc field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCcc

`func (o *DXCC) SetCcc(v string)`

SetCcc sets Ccc field to given value.

### HasCcc

`func (o *DXCC) HasCcc() bool`

HasCcc returns a boolean if a field has been set.

### GetName

`func (o *DXCC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DXCC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetName

`func (o *DXCC) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DXCC) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContinent

`func (o *DXCC) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *DXCC) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero
value otherwise and a boolean to check if the value has been set.

### SetContinent

`func (o *DXCC) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *DXCC) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetItuzone

`func (o *DXCC) GetItuzone() string`

GetItuzone returns the Ituzone field if non-nil, zero value otherwise.

### GetItuzoneOk

`func (o *DXCC) GetItuzoneOk() (*string, bool)`

GetItuzoneOk returns a tuple with the Ituzone field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetItuzone

`func (o *DXCC) SetItuzone(v string)`

SetItuzone sets Ituzone field to given value.

### HasItuzone

`func (o *DXCC) HasItuzone() bool`

HasItuzone returns a boolean if a field has been set.

### GetCqzone

`func (o *DXCC) GetCqzone() string`

GetCqzone returns the Cqzone field if non-nil, zero value otherwise.

### GetCqzoneOk

`func (o *DXCC) GetCqzoneOk() (*string, bool)`

GetCqzoneOk returns a tuple with the Cqzone field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCqzone

`func (o *DXCC) SetCqzone(v string)`

SetCqzone sets Cqzone field to given value.

### HasCqzone

`func (o *DXCC) HasCqzone() bool`

HasCqzone returns a boolean if a field has been set.

### GetTimezone

`func (o *DXCC) GetTimezone() int32`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DXCC) GetTimezoneOk() (*int32, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero
value otherwise and a boolean to check if the value has been set.

### SetTimezone

`func (o *DXCC) SetTimezone(v int32)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DXCC) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetLat

`func (o *DXCC) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *DXCC) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLat

`func (o *DXCC) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *DXCC) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *DXCC) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *DXCC) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLon

`func (o *DXCC) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *DXCC) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetNotes

`func (o *DXCC) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DXCC) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetNotes

`func (o *DXCC) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DXCC) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)
