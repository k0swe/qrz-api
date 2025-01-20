# QRZDatabase

## Properties

| Name         | Type                                   | Description | Notes      |
| ------------ | -------------------------------------- | ----------- | ---------- |
| **Session**  | Pointer to [**Session**](Session.md)   |             | [optional] |
| **Callsign** | Pointer to [**Callsign**](Callsign.md) |             | [optional] |
| **DXCC**     | Pointer to [**DXCC**](DXCC.md)         |             | [optional] |

## Methods

### NewQRZDatabase

`func NewQRZDatabase() *QRZDatabase`

NewQRZDatabase instantiates a new QRZDatabase object This constructor will
assign default values to properties that have it defined, and makes sure
properties required by API are set, but the set of arguments will change when
the set of required properties is changed

### NewQRZDatabaseWithDefaults

`func NewQRZDatabaseWithDefaults() *QRZDatabase`

NewQRZDatabaseWithDefaults instantiates a new QRZDatabase object This
constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *QRZDatabase) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *QRZDatabase) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetSession

`func (o *QRZDatabase) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *QRZDatabase) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetCallsign

`func (o *QRZDatabase) GetCallsign() Callsign`

GetCallsign returns the Callsign field if non-nil, zero value otherwise.

### GetCallsignOk

`func (o *QRZDatabase) GetCallsignOk() (*Callsign, bool)`

GetCallsignOk returns a tuple with the Callsign field if it's non-nil, zero
value otherwise and a boolean to check if the value has been set.

### SetCallsign

`func (o *QRZDatabase) SetCallsign(v Callsign)`

SetCallsign sets Callsign field to given value.

### HasCallsign

`func (o *QRZDatabase) HasCallsign() bool`

HasCallsign returns a boolean if a field has been set.

### GetDXCC

`func (o *QRZDatabase) GetDXCC() DXCC`

GetDXCC returns the DXCC field if non-nil, zero value otherwise.

### GetDXCCOk

`func (o *QRZDatabase) GetDXCCOk() (*DXCC, bool)`

GetDXCCOk returns a tuple with the DXCC field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetDXCC

`func (o *QRZDatabase) SetDXCC(v DXCC)`

SetDXCC sets DXCC field to given value.

### HasDXCC

`func (o *QRZDatabase) HasDXCC() bool`

HasDXCC returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)
