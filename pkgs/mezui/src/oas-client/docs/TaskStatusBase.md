# TaskStatusBase

Common fields shared by all task status responses. When `status` is `completed` the operation-specific `result` field is populated. When `status` is `failed` the `error` field is populated.

## Properties

| Name             | Type       | Description                   | Notes                             |
| ---------------- | ---------- | ----------------------------- | --------------------------------- |
| **task_id**      | **string** |                               | [default to undefined]            |
| **status**       | **string** |                               | [default to undefined]            |
| **created_at**   | **string** |                               | [default to undefined]            |
| **completed_at** | **string** |                               | [optional] [default to undefined] |
| **message**      | **string** | Human-readable status message | [optional] [default to undefined] |
| **error**        | **Error**  |                               | [optional] [default to undefined] |

## Example

```typescript
import { TaskStatusBase } from "./api";

const instance: TaskStatusBase = {
  task_id,
  status,
  created_at,
  completed_at,
  message,
  error,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
