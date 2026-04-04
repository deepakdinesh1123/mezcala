# DatabaseSyncApi

All URIs are relative to _localhost:7007_

| Method                                  | HTTP request                                      | Description                           |
| --------------------------------------- | ------------------------------------------------- | ------------------------------------- |
| [**createSyncRule**](#createsyncrule)   | **POST** /databases/{dbID}/sync/init              | Create a new sync rule for a database |
| [**deleteSyncRule**](#deletesyncrule)   | **DELETE** /databases/{dbID}/sync/{syncID}/delete | Delete a sync rule                    |
| [**executeSyncRule**](#executesyncrule) | **PUT** /databases/{dbID}/sync                    | Trigger execution of a sync rule      |
| [**updateSyncRule**](#updatesyncrule)   | **PUT** /databases/{dbID}/sync/{syncID}/update    | Update an existing sync rule          |

# **createSyncRule**

> CreateSyncRule201Response createSyncRule(syncRuleConfig)

Synchronously creates a sync rule and returns it with its assigned ID.

### Example

```typescript
import { DatabaseSyncApi, Configuration, SyncRuleConfig } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabaseSyncApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let syncRuleConfig: SyncRuleConfig; //Payload for creating or updating a sync rule.

const { status, data } = await apiInstance.createSyncRule(dbID, syncRuleConfig);
```

### Parameters

| Name               | Type               | Description                                   | Notes                 |
| ------------------ | ------------------ | --------------------------------------------- | --------------------- |
| **syncRuleConfig** | **SyncRuleConfig** | Payload for creating or updating a sync rule. |                       |
| **dbID**           | [**number**]       | Unique identifier of the database             | defaults to undefined |

### Return type

**CreateSyncRule201Response**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details

| Status code | Description                                   | Response headers |
| ----------- | --------------------------------------------- | ---------------- |
| **201**     | The created or updated sync rule.             | -                |
| **400**     | The request is malformed or fails validation. | -                |
| **404**     | The requested resource does not exist.        | -                |
| **500**     | An unexpected server-side error occurred.     | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteSyncRule**

> deleteSyncRule()

Synchronously deletes a sync rule.

### Example

```typescript
import { DatabaseSyncApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabaseSyncApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let syncID: number; //Unique identifier of the sync rule (default to undefined)

const { status, data } = await apiInstance.deleteSyncRule(dbID, syncID);
```

### Parameters

| Name       | Type         | Description                        | Notes                 |
| ---------- | ------------ | ---------------------------------- | --------------------- |
| **dbID**   | [**number**] | Unique identifier of the database  | defaults to undefined |
| **syncID** | [**number**] | Unique identifier of the sync rule | defaults to undefined |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details

| Status code | Description                               | Response headers |
| ----------- | ----------------------------------------- | ---------------- |
| **204**     | Sync rule deleted successfully.           | -                |
| **404**     | The requested resource does not exist.    | -                |
| **500**     | An unexpected server-side error occurred. | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **executeSyncRule**

> TaskReference executeSyncRule(executeSyncRuleRequest)

Triggers a sync rule run asynchronously. Returns a task ID that can be polled via `GET /databases/{dbID}/sync/{syncID}/tasks/{taskID}`. On completion the task result contains the sync execution summary.

### Example

```typescript
import { DatabaseSyncApi, Configuration, ExecuteSyncRuleRequest } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabaseSyncApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let executeSyncRuleRequest: ExecuteSyncRuleRequest; //Identifies which sync rule to execute.

const { status, data } = await apiInstance.executeSyncRule(
  dbID,
  executeSyncRuleRequest,
);
```

### Parameters

| Name                       | Type                       | Description                            | Notes                 |
| -------------------------- | -------------------------- | -------------------------------------- | --------------------- |
| **executeSyncRuleRequest** | **ExecuteSyncRuleRequest** | Identifies which sync rule to execute. |                       |
| **dbID**                   | [**number**]               | Unique identifier of the database      | defaults to undefined |

### Return type

**TaskReference**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details

| Status code | Description                                                                                                                   | Response headers |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------- | ---------------- |
| **202**     | The operation has been accepted and is running asynchronously. Poll the corresponding task status endpoint to track progress. | -                |
| **400**     | The request is malformed or fails validation.                                                                                 | -                |
| **404**     | The requested resource does not exist.                                                                                        | -                |
| **422**     | The request is well-formed but contains semantic errors.                                                                      | -                |
| **500**     | An unexpected server-side error occurred.                                                                                     | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateSyncRule**

> CreateSyncRule201Response updateSyncRule(syncRuleConfig)

Synchronously updates a sync rule and returns the updated record.

### Example

```typescript
import { DatabaseSyncApi, Configuration, SyncRuleConfig } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabaseSyncApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let syncID: number; //Unique identifier of the sync rule (default to undefined)
let syncRuleConfig: SyncRuleConfig; //Payload for creating or updating a sync rule.

const { status, data } = await apiInstance.updateSyncRule(
  dbID,
  syncID,
  syncRuleConfig,
);
```

### Parameters

| Name               | Type               | Description                                   | Notes                 |
| ------------------ | ------------------ | --------------------------------------------- | --------------------- |
| **syncRuleConfig** | **SyncRuleConfig** | Payload for creating or updating a sync rule. |                       |
| **dbID**           | [**number**]       | Unique identifier of the database             | defaults to undefined |
| **syncID**         | [**number**]       | Unique identifier of the sync rule            | defaults to undefined |

### Return type

**CreateSyncRule201Response**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details

| Status code | Description                                   | Response headers |
| ----------- | --------------------------------------------- | ---------------- |
| **200**     | The created or updated sync rule.             | -                |
| **400**     | The request is malformed or fails validation. | -                |
| **404**     | The requested resource does not exist.        | -                |
| **500**     | An unexpected server-side error occurred.     | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
