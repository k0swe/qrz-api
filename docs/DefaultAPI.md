# \DefaultAPI

All URIs are relative to *https://xmldata.qrz.com/xml/1.34*

| Method                               | HTTP request | Description                |
| ------------------------------------ | ------------ | -------------------------- |
| [**RootGet**](DefaultAPI.md#RootGet) | **Get** /    | The do-everything endpoint |

## RootGet

> QRZDatabase
> RootGet(ctx).Username(username).Password(password).Agent(agent).S(s).Callsign(callsign).Dxcc(dxcc).Execute()

The do-everything endpoint

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/k0swe/qrz-api"
)

func main() {
	username := "username_example" // string |  (optional)
	password := "password_example" // string |  (optional)
	agent := "agent_example" // string |  (optional)
	s := "s_example" // string | session token (optional)
	callsign := "callsign_example" // string | perform a callsign info lookup (optional)
	dxcc := openapiclient.__get_dxcc_parameter{Float32: new(float32)} // GetDxccParameter | perform a DXCC info lookup (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RootGet(context.Background()).Username(username).Password(password).Agent(agent).S(s).Callsign(callsign).Dxcc(dxcc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RootGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RootGet`: QRZDatabase
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RootGet`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via
the builder pattern

| Name         | Type                                        | Description                    | Notes |
| ------------ | ------------------------------------------- | ------------------------------ | ----- |
| **username** | **string**                                  |                                |
| **password** | **string**                                  |                                |
| **agent**    | **string**                                  |                                |
| **s**        | **string**                                  | session token                  |
| **callsign** | **string**                                  | perform a callsign info lookup |
| **dxcc**     | [**GetDxccParameter**](GetDxccParameter.md) | perform a DXCC info lookup     |

### Return type

[**QRZDatabase**](QRZDatabase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#)
[[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
