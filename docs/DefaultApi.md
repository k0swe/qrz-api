# \DefaultApi

All URIs are relative to *https://xmldata.qrz.com/xml/1.34*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RootGet**](DefaultApi.md#RootGet) | **Get** / | The do-everything endpoint



## RootGet

> QrzDatabase RootGet(ctx, optional)

The do-everything endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RootGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RootGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 
 **password** | **optional.String**|  | 
 **agent** | **optional.String**|  | 
 **s** | **optional.String**| session token | 
 **callsign** | **optional.String**| perform a callsign info lookup | 
 **dxcc** | [**optional.Interface of OneOfnumberstring**](.md)| perform a DXCC info lookup | 

### Return type

[**QrzDatabase**](QRZDatabase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

