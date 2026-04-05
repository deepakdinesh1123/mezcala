# DatabaseAcknowledgment

Confirmation returned after a database is successfully created or registered. Contains only the system-assigned ID. 

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **number** | System-assigned unique identifier for the new database | [default to undefined]
**message** | **string** |  | [default to undefined]

## Example

```typescript
import { DatabaseAcknowledgment } from './api';

const instance: DatabaseAcknowledgment = {
    id,
    message,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
