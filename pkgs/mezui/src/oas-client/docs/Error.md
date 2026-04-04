# ModelError

Standard error envelope returned for all 4xx and 5xx responses.

## Properties

| Name        | Type                        | Description                                 | Notes                             |
| ----------- | --------------------------- | ------------------------------------------- | --------------------------------- |
| **code**    | **string**                  | Machine-readable error code                 | [default to undefined]            |
| **message** | **string**                  | Human-readable description of the error     | [default to undefined]            |
| **details** | **{ [key: string]: any; }** | Optional additional context about the error | [optional] [default to undefined] |

## Example

```typescript
import { ModelError } from "./api";

const instance: ModelError = {
  code,
  message,
  details,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
