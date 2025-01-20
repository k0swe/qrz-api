# Session

## Properties

| Name        | Type                   | Description                                                                                                                                                                                                                                        | Notes      |
| ----------- | ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Key**     | Pointer to **string**  | a valid user session key. A user session is established whenever a session key is returned. Any response from the server that does not contain the Key element indicates that no valid session exists and that a re-login is required to continue. | [optional] |
| **Count**   | Pointer to **float32** | Number of lookups performed by this user in the current 24 hour period                                                                                                                                                                             | [optional] |
| **SubExp**  | Pointer to **string**  | time and date that the users subscription will expire - or - \&quot;non-subscriber\&quot;                                                                                                                                                          | [optional] |
| **GMTime**  | Pointer to **string**  | Time stamp for this message                                                                                                                                                                                                                        | [optional] |
| **Message** | Pointer to **string**  | An informational message for the user                                                                                                                                                                                                              | [optional] |
| **Error**   | Pointer to **string**  | XML system error message                                                                                                                                                                                                                           | [optional] |

## Methods

### NewSession

`func NewSession() *Session`

NewSession instantiates a new Session object This constructor will assign
default values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set of
required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object This constructor will
only assign default values to properties that have it defined, but it doesn't
guarantee that properties required by API are set

### GetKey

`func (o *Session) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Session) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetKey

`func (o *Session) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Session) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCount

`func (o *Session) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Session) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCount

`func (o *Session) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Session) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSubExp

`func (o *Session) GetSubExp() string`

GetSubExp returns the SubExp field if non-nil, zero value otherwise.

### GetSubExpOk

`func (o *Session) GetSubExpOk() (*string, bool)`

GetSubExpOk returns a tuple with the SubExp field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetSubExp

`func (o *Session) SetSubExp(v string)`

SetSubExp sets SubExp field to given value.

### HasSubExp

`func (o *Session) HasSubExp() bool`

HasSubExp returns a boolean if a field has been set.

### GetGMTime

`func (o *Session) GetGMTime() string`

GetGMTime returns the GMTime field if non-nil, zero value otherwise.

### GetGMTimeOk

`func (o *Session) GetGMTimeOk() (*string, bool)`

GetGMTimeOk returns a tuple with the GMTime field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetGMTime

`func (o *Session) SetGMTime(v string)`

SetGMTime sets GMTime field to given value.

### HasGMTime

`func (o *Session) HasGMTime() bool`

HasGMTime returns a boolean if a field has been set.

### GetMessage

`func (o *Session) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Session) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetMessage

`func (o *Session) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Session) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *Session) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Session) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetError

`func (o *Session) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Session) HasError() bool`

HasError returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)
