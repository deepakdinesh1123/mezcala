# DeleteDatabaseTaskStatus

Task status for a `deleteDatabase` operation. On completion, `result` contains a deletion confirmation. 

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**task_id** | **string** |  | [default to undefined]
**status** | **string** |  | [default to undefined]
**created_at** | **string** |  | [default to undefined]
**completed_at** | **string** |  | [optional] [default to undefined]
**message** | **string** | Human-readable status message | [optional] [default to undefined]
**error** | **Error** |  | [optional] [default to undefined]
**result** | [**DeletionAcknowledgment**](DeletionAcknowledgment.md) |  | [optional] [default to undefined]

## Example

```typescript
import { DeleteDatabaseTaskStatus } from './api';

const instance: DeleteDatabaseTaskStatus = {
    task_id,
    status,
    created_at,
    completed_at,
    message,
    error,
    result,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
