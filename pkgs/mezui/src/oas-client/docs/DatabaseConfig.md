# DatabaseConfig

Connection configuration for a database.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** |  | [optional] [default to undefined]
**username** | **string** |  | [default to undefined]
**password** | **string** |  | [default to undefined]
**engine** | **string** |  | [default to undefined]
**version** | **string** |  | [optional] [default to undefined]
**alias** | **string** | Optional human-friendly label | [optional] [default to undefined]
**host** | **string** |  | [default to undefined]
**port** | **string** |  | [default to undefined]
**ssl_mode** | **string** |  | [optional] [default to undefined]
**path** | **string** | Optional filesystem path for local/embedded engines | [optional] [default to undefined]
**role** | **string** | Database role for connection | [optional] [default to undefined]
**image** | **string** | The docker image to use | [optional] [default to undefined]

## Example

```typescript
import { DatabaseConfig } from './api';

const instance: DatabaseConfig = {
    name,
    username,
    password,
    engine,
    version,
    alias,
    host,
    port,
    ssl_mode,
    path,
    role,
    image,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
