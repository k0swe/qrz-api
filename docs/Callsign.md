# Callsign

## Properties

| Name          | Type                   | Description                                                                                   | Notes      |
| ------------- | ---------------------- | --------------------------------------------------------------------------------------------- | ---------- |
| **Call**      | Pointer to **string**  | callsign                                                                                      | [optional] |
| **Xref**      | Pointer to **string**  | Cross reference: the query callsign that returned this record                                 | [optional] |
| **Aliases**   | Pointer to **string**  | Other callsigns that resolve to this record                                                   | [optional] |
| **Dxcc**      | Pointer to **string**  | DXCC entity ID (country code) for the callsign                                                | [optional] |
| **Fname**     | Pointer to **string**  | first name                                                                                    | [optional] |
| **Name**      | Pointer to **string**  | last name                                                                                     | [optional] |
| **Addr1**     | Pointer to **string**  | address line 1 (i.e. house # and street)                                                      | [optional] |
| **Addr2**     | Pointer to **string**  | address line 2 (i.e, city name)                                                               | [optional] |
| **State**     | Pointer to **string**  | state (USA Only)                                                                              | [optional] |
| **Zip**       | Pointer to **string**  | Zip/postal code                                                                               | [optional] |
| **Country**   | Pointer to **string**  | country name for the QSL mailing address                                                      | [optional] |
| **Ccode**     | Pointer to **string**  | dxcc entity code for the mailing address country                                              | [optional] |
| **Lat**       | Pointer to **float64** | lattitude of address (signed decimal) S &lt; 0 &gt; N                                         | [optional] |
| **Lon**       | Pointer to **float64** | longitude of address (signed decimal) W &lt; 0 &gt; E                                         | [optional] |
| **Grid**      | Pointer to **string**  | Maidenhead grid locator                                                                       | [optional] |
| **County**    | Pointer to **string**  | county name (USA)                                                                             | [optional] |
| **Fips**      | Pointer to **string**  | FIPS county identifier (USA)                                                                  | [optional] |
| **Land**      | Pointer to **string**  | DXCC country name of the callsign                                                             | [optional] |
| **Efdate**    | Pointer to **string**  | license effective date (USA)                                                                  | [optional] |
| **Expdate**   | Pointer to **string**  | license expiration date (USA)                                                                 | [optional] |
| **PCall**     | Pointer to **string**  | previous callsign                                                                             | [optional] |
| **Class**     | Pointer to **string**  | license class                                                                                 | [optional] |
| **Codes**     | Pointer to **string**  | license type codes (USA)                                                                      | [optional] |
| **Qslmgr**    | Pointer to **string**  | QSL manager info                                                                              | [optional] |
| **Email**     | Pointer to **string**  | email address                                                                                 | [optional] |
| **Url**       | Pointer to **string**  | web page address                                                                              | [optional] |
| **UViews**    | Pointer to **float32** | QRZ web page views                                                                            | [optional] |
| **Bio**       | Pointer to **string**  | approximate length of the bio HTML in bytes                                                   | [optional] |
| **Biodate**   | Pointer to **string**  | date of the last bio update                                                                   | [optional] |
| **Image**     | Pointer to **string**  | full URL of the callsign&#39;s primary image                                                  | [optional] |
| **Imageinfo** | Pointer to **string**  | height:width:size in bytes, of the image file                                                 | [optional] |
| **Serial**    | Pointer to **string**  | QRZ db serial number                                                                          | [optional] |
| **Moddate**   | Pointer to **string**  | QRZ callsign last modified date                                                               | [optional] |
| **MSA**       | Pointer to **string**  | Metro Service Area (USPS)                                                                     | [optional] |
| **AreaCode**  | Pointer to **string**  | Telephone Area Code (USA)                                                                     | [optional] |
| **TimeZone**  | Pointer to **string**  | Time Zone (USA)                                                                               | [optional] |
| **GMTOffset** | Pointer to **string**  | GMT Time Offset                                                                               | [optional] |
| **DST**       | Pointer to **string**  | Daylight Saving Time Observed                                                                 | [optional] |
| **Eqsl**      | Pointer to **string**  | Will accept e-qsl (0/1 or blank if unknown)                                                   | [optional] |
| **Mqsl**      | Pointer to **string**  | Will return paper QSL (0/1 or blank if unknown)                                               | [optional] |
| **Cqzone**    | Pointer to **string**  | CQ Zone identifier                                                                            | [optional] |
| **Ituzone**   | Pointer to **string**  | ITU Zone identifier                                                                           | [optional] |
| **Born**      | Pointer to **float32** | operator&#39;s year of birth                                                                  | [optional] |
| **User**      | Pointer to **string**  | User who manages this callsign on QRZ                                                         | [optional] |
| **Lotw**      | Pointer to **string**  | Will accept LOTW (0/1 or blank if unknown)                                                    | [optional] |
| **Iota**      | Pointer to **string**  | IOTA Designator (blank if unknown)                                                            | [optional] |
| **Geoloc**    | Pointer to **string**  | Describes source of lat/long data                                                             | [optional] |
| **Attn**      | Pointer to **string**  | Attention address line, this line should be prepended to the address                          | [optional] |
| **Nickname**  | Pointer to **string**  | A different or shortened name used on the air                                                 | [optional] |
| **NameFmt**   | Pointer to **string**  | Combined full name and nickname in the format used by QRZ. This fortmat is subject to change. | [optional] |

## Methods

### NewCallsign

`func NewCallsign() *Callsign`

NewCallsign instantiates a new Callsign object This constructor will assign
default values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set of
required properties is changed

### NewCallsignWithDefaults

`func NewCallsignWithDefaults() *Callsign`

NewCallsignWithDefaults instantiates a new Callsign object This constructor will
only assign default values to properties that have it defined, but it doesn't
guarantee that properties required by API are set

### GetCall

`func (o *Callsign) GetCall() string`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *Callsign) GetCallOk() (*string, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCall

`func (o *Callsign) SetCall(v string)`

SetCall sets Call field to given value.

### HasCall

`func (o *Callsign) HasCall() bool`

HasCall returns a boolean if a field has been set.

### GetXref

`func (o *Callsign) GetXref() string`

GetXref returns the Xref field if non-nil, zero value otherwise.

### GetXrefOk

`func (o *Callsign) GetXrefOk() (*string, bool)`

GetXrefOk returns a tuple with the Xref field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetXref

`func (o *Callsign) SetXref(v string)`

SetXref sets Xref field to given value.

### HasXref

`func (o *Callsign) HasXref() bool`

HasXref returns a boolean if a field has been set.

### GetAliases

`func (o *Callsign) GetAliases() string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Callsign) GetAliasesOk() (*string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetAliases

`func (o *Callsign) SetAliases(v string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Callsign) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetDxcc

`func (o *Callsign) GetDxcc() string`

GetDxcc returns the Dxcc field if non-nil, zero value otherwise.

### GetDxccOk

`func (o *Callsign) GetDxccOk() (*string, bool)`

GetDxccOk returns a tuple with the Dxcc field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetDxcc

`func (o *Callsign) SetDxcc(v string)`

SetDxcc sets Dxcc field to given value.

### HasDxcc

`func (o *Callsign) HasDxcc() bool`

HasDxcc returns a boolean if a field has been set.

### GetFname

`func (o *Callsign) GetFname() string`

GetFname returns the Fname field if non-nil, zero value otherwise.

### GetFnameOk

`func (o *Callsign) GetFnameOk() (*string, bool)`

GetFnameOk returns a tuple with the Fname field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetFname

`func (o *Callsign) SetFname(v string)`

SetFname sets Fname field to given value.

### HasFname

`func (o *Callsign) HasFname() bool`

HasFname returns a boolean if a field has been set.

### GetName

`func (o *Callsign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Callsign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetName

`func (o *Callsign) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Callsign) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddr1

`func (o *Callsign) GetAddr1() string`

GetAddr1 returns the Addr1 field if non-nil, zero value otherwise.

### GetAddr1Ok

`func (o *Callsign) GetAddr1Ok() (*string, bool)`

GetAddr1Ok returns a tuple with the Addr1 field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetAddr1

`func (o *Callsign) SetAddr1(v string)`

SetAddr1 sets Addr1 field to given value.

### HasAddr1

`func (o *Callsign) HasAddr1() bool`

HasAddr1 returns a boolean if a field has been set.

### GetAddr2

`func (o *Callsign) GetAddr2() string`

GetAddr2 returns the Addr2 field if non-nil, zero value otherwise.

### GetAddr2Ok

`func (o *Callsign) GetAddr2Ok() (*string, bool)`

GetAddr2Ok returns a tuple with the Addr2 field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetAddr2

`func (o *Callsign) SetAddr2(v string)`

SetAddr2 sets Addr2 field to given value.

### HasAddr2

`func (o *Callsign) HasAddr2() bool`

HasAddr2 returns a boolean if a field has been set.

### GetState

`func (o *Callsign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Callsign) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetState

`func (o *Callsign) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Callsign) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *Callsign) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *Callsign) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetZip

`func (o *Callsign) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *Callsign) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCountry

`func (o *Callsign) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Callsign) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCountry

`func (o *Callsign) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Callsign) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCcode

`func (o *Callsign) GetCcode() string`

GetCcode returns the Ccode field if non-nil, zero value otherwise.

### GetCcodeOk

`func (o *Callsign) GetCcodeOk() (*string, bool)`

GetCcodeOk returns a tuple with the Ccode field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCcode

`func (o *Callsign) SetCcode(v string)`

SetCcode sets Ccode field to given value.

### HasCcode

`func (o *Callsign) HasCcode() bool`

HasCcode returns a boolean if a field has been set.

### GetLat

`func (o *Callsign) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *Callsign) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLat

`func (o *Callsign) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *Callsign) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *Callsign) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *Callsign) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLon

`func (o *Callsign) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *Callsign) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetGrid

`func (o *Callsign) GetGrid() string`

GetGrid returns the Grid field if non-nil, zero value otherwise.

### GetGridOk

`func (o *Callsign) GetGridOk() (*string, bool)`

GetGridOk returns a tuple with the Grid field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetGrid

`func (o *Callsign) SetGrid(v string)`

SetGrid sets Grid field to given value.

### HasGrid

`func (o *Callsign) HasGrid() bool`

HasGrid returns a boolean if a field has been set.

### GetCounty

`func (o *Callsign) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *Callsign) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCounty

`func (o *Callsign) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *Callsign) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetFips

`func (o *Callsign) GetFips() string`

GetFips returns the Fips field if non-nil, zero value otherwise.

### GetFipsOk

`func (o *Callsign) GetFipsOk() (*string, bool)`

GetFipsOk returns a tuple with the Fips field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetFips

`func (o *Callsign) SetFips(v string)`

SetFips sets Fips field to given value.

### HasFips

`func (o *Callsign) HasFips() bool`

HasFips returns a boolean if a field has been set.

### GetLand

`func (o *Callsign) GetLand() string`

GetLand returns the Land field if non-nil, zero value otherwise.

### GetLandOk

`func (o *Callsign) GetLandOk() (*string, bool)`

GetLandOk returns a tuple with the Land field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLand

`func (o *Callsign) SetLand(v string)`

SetLand sets Land field to given value.

### HasLand

`func (o *Callsign) HasLand() bool`

HasLand returns a boolean if a field has been set.

### GetEfdate

`func (o *Callsign) GetEfdate() string`

GetEfdate returns the Efdate field if non-nil, zero value otherwise.

### GetEfdateOk

`func (o *Callsign) GetEfdateOk() (*string, bool)`

GetEfdateOk returns a tuple with the Efdate field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetEfdate

`func (o *Callsign) SetEfdate(v string)`

SetEfdate sets Efdate field to given value.

### HasEfdate

`func (o *Callsign) HasEfdate() bool`

HasEfdate returns a boolean if a field has been set.

### GetExpdate

`func (o *Callsign) GetExpdate() string`

GetExpdate returns the Expdate field if non-nil, zero value otherwise.

### GetExpdateOk

`func (o *Callsign) GetExpdateOk() (*string, bool)`

GetExpdateOk returns a tuple with the Expdate field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetExpdate

`func (o *Callsign) SetExpdate(v string)`

SetExpdate sets Expdate field to given value.

### HasExpdate

`func (o *Callsign) HasExpdate() bool`

HasExpdate returns a boolean if a field has been set.

### GetPCall

`func (o *Callsign) GetPCall() string`

GetPCall returns the PCall field if non-nil, zero value otherwise.

### GetPCallOk

`func (o *Callsign) GetPCallOk() (*string, bool)`

GetPCallOk returns a tuple with the PCall field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetPCall

`func (o *Callsign) SetPCall(v string)`

SetPCall sets PCall field to given value.

### HasPCall

`func (o *Callsign) HasPCall() bool`

HasPCall returns a boolean if a field has been set.

### GetClass

`func (o *Callsign) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Callsign) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetClass

`func (o *Callsign) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Callsign) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCodes

`func (o *Callsign) GetCodes() string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *Callsign) GetCodesOk() (*string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCodes

`func (o *Callsign) SetCodes(v string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *Callsign) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetQslmgr

`func (o *Callsign) GetQslmgr() string`

GetQslmgr returns the Qslmgr field if non-nil, zero value otherwise.

### GetQslmgrOk

`func (o *Callsign) GetQslmgrOk() (*string, bool)`

GetQslmgrOk returns a tuple with the Qslmgr field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetQslmgr

`func (o *Callsign) SetQslmgr(v string)`

SetQslmgr sets Qslmgr field to given value.

### HasQslmgr

`func (o *Callsign) HasQslmgr() bool`

HasQslmgr returns a boolean if a field has been set.

### GetEmail

`func (o *Callsign) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Callsign) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetEmail

`func (o *Callsign) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Callsign) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUrl

`func (o *Callsign) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Callsign) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetUrl

`func (o *Callsign) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Callsign) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUViews

`func (o *Callsign) GetUViews() float32`

GetUViews returns the UViews field if non-nil, zero value otherwise.

### GetUViewsOk

`func (o *Callsign) GetUViewsOk() (*float32, bool)`

GetUViewsOk returns a tuple with the UViews field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetUViews

`func (o *Callsign) SetUViews(v float32)`

SetUViews sets UViews field to given value.

### HasUViews

`func (o *Callsign) HasUViews() bool`

HasUViews returns a boolean if a field has been set.

### GetBio

`func (o *Callsign) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *Callsign) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetBio

`func (o *Callsign) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *Callsign) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBiodate

`func (o *Callsign) GetBiodate() string`

GetBiodate returns the Biodate field if non-nil, zero value otherwise.

### GetBiodateOk

`func (o *Callsign) GetBiodateOk() (*string, bool)`

GetBiodateOk returns a tuple with the Biodate field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetBiodate

`func (o *Callsign) SetBiodate(v string)`

SetBiodate sets Biodate field to given value.

### HasBiodate

`func (o *Callsign) HasBiodate() bool`

HasBiodate returns a boolean if a field has been set.

### GetImage

`func (o *Callsign) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Callsign) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetImage

`func (o *Callsign) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Callsign) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageinfo

`func (o *Callsign) GetImageinfo() string`

GetImageinfo returns the Imageinfo field if non-nil, zero value otherwise.

### GetImageinfoOk

`func (o *Callsign) GetImageinfoOk() (*string, bool)`

GetImageinfoOk returns a tuple with the Imageinfo field if it's non-nil, zero
value otherwise and a boolean to check if the value has been set.

### SetImageinfo

`func (o *Callsign) SetImageinfo(v string)`

SetImageinfo sets Imageinfo field to given value.

### HasImageinfo

`func (o *Callsign) HasImageinfo() bool`

HasImageinfo returns a boolean if a field has been set.

### GetSerial

`func (o *Callsign) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Callsign) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetSerial

`func (o *Callsign) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Callsign) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModdate

`func (o *Callsign) GetModdate() string`

GetModdate returns the Moddate field if non-nil, zero value otherwise.

### GetModdateOk

`func (o *Callsign) GetModdateOk() (*string, bool)`

GetModdateOk returns a tuple with the Moddate field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetModdate

`func (o *Callsign) SetModdate(v string)`

SetModdate sets Moddate field to given value.

### HasModdate

`func (o *Callsign) HasModdate() bool`

HasModdate returns a boolean if a field has been set.

### GetMSA

`func (o *Callsign) GetMSA() string`

GetMSA returns the MSA field if non-nil, zero value otherwise.

### GetMSAOk

`func (o *Callsign) GetMSAOk() (*string, bool)`

GetMSAOk returns a tuple with the MSA field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetMSA

`func (o *Callsign) SetMSA(v string)`

SetMSA sets MSA field to given value.

### HasMSA

`func (o *Callsign) HasMSA() bool`

HasMSA returns a boolean if a field has been set.

### GetAreaCode

`func (o *Callsign) GetAreaCode() string`

GetAreaCode returns the AreaCode field if non-nil, zero value otherwise.

### GetAreaCodeOk

`func (o *Callsign) GetAreaCodeOk() (*string, bool)`

GetAreaCodeOk returns a tuple with the AreaCode field if it's non-nil, zero
value otherwise and a boolean to check if the value has been set.

### SetAreaCode

`func (o *Callsign) SetAreaCode(v string)`

SetAreaCode sets AreaCode field to given value.

### HasAreaCode

`func (o *Callsign) HasAreaCode() bool`

HasAreaCode returns a boolean if a field has been set.

### GetTimeZone

`func (o *Callsign) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Callsign) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero
value otherwise and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Callsign) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Callsign) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetGMTOffset

`func (o *Callsign) GetGMTOffset() string`

GetGMTOffset returns the GMTOffset field if non-nil, zero value otherwise.

### GetGMTOffsetOk

`func (o *Callsign) GetGMTOffsetOk() (*string, bool)`

GetGMTOffsetOk returns a tuple with the GMTOffset field if it's non-nil, zero
value otherwise and a boolean to check if the value has been set.

### SetGMTOffset

`func (o *Callsign) SetGMTOffset(v string)`

SetGMTOffset sets GMTOffset field to given value.

### HasGMTOffset

`func (o *Callsign) HasGMTOffset() bool`

HasGMTOffset returns a boolean if a field has been set.

### GetDST

`func (o *Callsign) GetDST() string`

GetDST returns the DST field if non-nil, zero value otherwise.

### GetDSTOk

`func (o *Callsign) GetDSTOk() (*string, bool)`

GetDSTOk returns a tuple with the DST field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetDST

`func (o *Callsign) SetDST(v string)`

SetDST sets DST field to given value.

### HasDST

`func (o *Callsign) HasDST() bool`

HasDST returns a boolean if a field has been set.

### GetEqsl

`func (o *Callsign) GetEqsl() string`

GetEqsl returns the Eqsl field if non-nil, zero value otherwise.

### GetEqslOk

`func (o *Callsign) GetEqslOk() (*string, bool)`

GetEqslOk returns a tuple with the Eqsl field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetEqsl

`func (o *Callsign) SetEqsl(v string)`

SetEqsl sets Eqsl field to given value.

### HasEqsl

`func (o *Callsign) HasEqsl() bool`

HasEqsl returns a boolean if a field has been set.

### GetMqsl

`func (o *Callsign) GetMqsl() string`

GetMqsl returns the Mqsl field if non-nil, zero value otherwise.

### GetMqslOk

`func (o *Callsign) GetMqslOk() (*string, bool)`

GetMqslOk returns a tuple with the Mqsl field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetMqsl

`func (o *Callsign) SetMqsl(v string)`

SetMqsl sets Mqsl field to given value.

### HasMqsl

`func (o *Callsign) HasMqsl() bool`

HasMqsl returns a boolean if a field has been set.

### GetCqzone

`func (o *Callsign) GetCqzone() string`

GetCqzone returns the Cqzone field if non-nil, zero value otherwise.

### GetCqzoneOk

`func (o *Callsign) GetCqzoneOk() (*string, bool)`

GetCqzoneOk returns a tuple with the Cqzone field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetCqzone

`func (o *Callsign) SetCqzone(v string)`

SetCqzone sets Cqzone field to given value.

### HasCqzone

`func (o *Callsign) HasCqzone() bool`

HasCqzone returns a boolean if a field has been set.

### GetItuzone

`func (o *Callsign) GetItuzone() string`

GetItuzone returns the Ituzone field if non-nil, zero value otherwise.

### GetItuzoneOk

`func (o *Callsign) GetItuzoneOk() (*string, bool)`

GetItuzoneOk returns a tuple with the Ituzone field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetItuzone

`func (o *Callsign) SetItuzone(v string)`

SetItuzone sets Ituzone field to given value.

### HasItuzone

`func (o *Callsign) HasItuzone() bool`

HasItuzone returns a boolean if a field has been set.

### GetBorn

`func (o *Callsign) GetBorn() float32`

GetBorn returns the Born field if non-nil, zero value otherwise.

### GetBornOk

`func (o *Callsign) GetBornOk() (*float32, bool)`

GetBornOk returns a tuple with the Born field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetBorn

`func (o *Callsign) SetBorn(v float32)`

SetBorn sets Born field to given value.

### HasBorn

`func (o *Callsign) HasBorn() bool`

HasBorn returns a boolean if a field has been set.

### GetUser

`func (o *Callsign) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Callsign) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetUser

`func (o *Callsign) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Callsign) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLotw

`func (o *Callsign) GetLotw() string`

GetLotw returns the Lotw field if non-nil, zero value otherwise.

### GetLotwOk

`func (o *Callsign) GetLotwOk() (*string, bool)`

GetLotwOk returns a tuple with the Lotw field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetLotw

`func (o *Callsign) SetLotw(v string)`

SetLotw sets Lotw field to given value.

### HasLotw

`func (o *Callsign) HasLotw() bool`

HasLotw returns a boolean if a field has been set.

### GetIota

`func (o *Callsign) GetIota() string`

GetIota returns the Iota field if non-nil, zero value otherwise.

### GetIotaOk

`func (o *Callsign) GetIotaOk() (*string, bool)`

GetIotaOk returns a tuple with the Iota field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetIota

`func (o *Callsign) SetIota(v string)`

SetIota sets Iota field to given value.

### HasIota

`func (o *Callsign) HasIota() bool`

HasIota returns a boolean if a field has been set.

### GetGeoloc

`func (o *Callsign) GetGeoloc() string`

GetGeoloc returns the Geoloc field if non-nil, zero value otherwise.

### GetGeolocOk

`func (o *Callsign) GetGeolocOk() (*string, bool)`

GetGeolocOk returns a tuple with the Geoloc field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetGeoloc

`func (o *Callsign) SetGeoloc(v string)`

SetGeoloc sets Geoloc field to given value.

### HasGeoloc

`func (o *Callsign) HasGeoloc() bool`

HasGeoloc returns a boolean if a field has been set.

### GetAttn

`func (o *Callsign) GetAttn() string`

GetAttn returns the Attn field if non-nil, zero value otherwise.

### GetAttnOk

`func (o *Callsign) GetAttnOk() (*string, bool)`

GetAttnOk returns a tuple with the Attn field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetAttn

`func (o *Callsign) SetAttn(v string)`

SetAttn sets Attn field to given value.

### HasAttn

`func (o *Callsign) HasAttn() bool`

HasAttn returns a boolean if a field has been set.

### GetNickname

`func (o *Callsign) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Callsign) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero
value otherwise and a boolean to check if the value has been set.

### SetNickname

`func (o *Callsign) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Callsign) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetNameFmt

`func (o *Callsign) GetNameFmt() string`

GetNameFmt returns the NameFmt field if non-nil, zero value otherwise.

### GetNameFmtOk

`func (o *Callsign) GetNameFmtOk() (*string, bool)`

GetNameFmtOk returns a tuple with the NameFmt field if it's non-nil, zero value
otherwise and a boolean to check if the value has been set.

### SetNameFmt

`func (o *Callsign) SetNameFmt(v string)`

SetNameFmt sets NameFmt field to given value.

### HasNameFmt

`func (o *Callsign) HasNameFmt() bool`

HasNameFmt returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)
