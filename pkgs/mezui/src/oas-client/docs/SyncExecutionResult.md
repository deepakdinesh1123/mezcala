# SyncExecutionResult

Outcome details of a sync rule execution.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sync_id** | **number** |  | [default to undefined]
**status** | **string** |  | [default to undefined]
**started_at** | **string** |  | [optional] [default to undefined]
**finished_at** | **string** |  | [optional] [default to undefined]
**tables_synced** | **number** | Number of tables processed during the sync | [optional] [default to undefined]
**rows_synced** | **number** | Total number of rows written to the destination | [optional] [default to undefined]
**message** | **string** |  | [optional] [default to undefined]

## Example

```typescript
import { SyncExecutionResult } from './api';

const instance: SyncExecutionResult = {
    sync_id,
    status,
    started_at,
    finished_at,
    tables_synced,
    rows_synced,
    message,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
