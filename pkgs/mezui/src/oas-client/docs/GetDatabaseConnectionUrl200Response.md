# GetDatabaseConnectionUrl200Response

## Properties

| Name               | Type       | Description                                                           | Notes                  |
| ------------------ | ---------- | --------------------------------------------------------------------- | ---------------------- |
| **id**             | **number** | Database ID                                                           | [default to undefined] |
| **connection_url** | **string** | Fully-formed connection URL including credentials. Treat as a secret. | [default to undefined] |

## Example

```typescript
import { GetDatabaseConnectionUrl200Response } from "./api";

const instance: GetDatabaseConnectionUrl200Response = {
  id,
  connection_url,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
