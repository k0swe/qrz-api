# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | a valid user session key. A user session is established whenever a session key is returned. Any response from the server that does not contain the Key element indicates that no valid session exists and that a re-login is required to continue. | [optional] 
**Count** | **float32** | Number of lookups performed by this user in the current 24 hour period | [optional] 
**SubExp** | **string** | time and date that the users subscription will expire - or - \&quot;non-subscriber\&quot; | [optional] 
**GMTime** | **string** | Time stamp for this message | [optional] 
**Message** | **string** | An informational message for the user | [optional] 
**Error** | **string** | XML system error message | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


