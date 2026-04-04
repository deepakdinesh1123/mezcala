# HealthApi

All URIs are relative to _localhost:7007_

| Method                          | HTTP request    | Description          |
| ------------------------------- | --------------- | -------------------- |
| [**checkHealth**](#checkhealth) | **GET** /health | Check Mezcala health |

# **checkHealth**

> CheckHealth200Response checkHealth()

### Example

```typescript
import { HealthApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new HealthApi(configuration);

const { status, data } = await apiInstance.checkHealth();
```

### Parameters

This endpoint does not have any parameters.

### Return type

**CheckHealth200Response**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details

| Status code | Description                                          | Response headers |
| ----------- | ---------------------------------------------------- | ---------------- |
| **200**     | Service is reachable and reports its current status. | -                |
| **500**     | An unexpected server-side error occurred.            | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
