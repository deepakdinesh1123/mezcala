# SyncRuleConfig

Configuration for syncing data between databases with optional masking rules.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**from** | **number** | Source database ID registered in Mezcala | [default to undefined]
**to** | **number** | Destination database ID registered in Mezcala | [default to undefined]
**exclude** | **Array&lt;string&gt;** | Table names to exclude from the sync | [optional] [default to undefined]
**schemas** | **Array&lt;string&gt;** | Schemas to include in the sync | [optional] [default to undefined]
**data_rules** | **{ [key: string]: string; }** | Column-level anonymization rules. Keys are column names or glob-style patterns (e.g. encrypted_*). Use null to drop a column entirely.  | [optional] [default to undefined]

## Example

```typescript
import { SyncRuleConfig } from './api';

const instance: SyncRuleConfig = {
    from,
    to,
    exclude,
    schemas,
    data_rules,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
