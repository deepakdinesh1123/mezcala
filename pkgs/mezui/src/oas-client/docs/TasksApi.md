# TasksApi

All URIs are relative to _localhost:7007_

| Method                                                          | HTTP request                                           | Description                               |
| --------------------------------------------------------------- | ------------------------------------------------------ | ----------------------------------------- |
| [**getCreateDatabaseTaskStatus**](#getcreatedatabasetaskstatus) | **GET** /databases/create/tasks/{taskID}               | Get the status of a createDatabase task   |
| [**getDeleteDatabaseTaskStatus**](#getdeletedatabasetaskstatus) | **GET** /databases/{dbID}/delete/tasks/{taskID}        | Get the status of a deleteDatabase task   |
| [**getExecuteSyncTaskStatus**](#getexecutesynctaskstatus)       | **GET** /databases/{dbID}/sync/{syncID}/tasks/{taskID} | Get the status of an executeSyncRule task |

# **getCreateDatabaseTaskStatus**

> CreateDatabaseTaskStatus getCreateDatabaseTaskStatus()

Polls the status of a task initiated by `POST /databases/create`. When `status` is `completed` the `result` field contains a `DatabaseAcknowledgment` with the new database ID.

### Example

```typescript
import { TasksApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new TasksApi(configuration);

let taskID: string; //Unique identifier of the async task (default to undefined)

const { status, data } = await apiInstance.getCreateDatabaseTaskStatus(taskID);
```

### Parameters

| Name       | Type         | Description                         | Notes                 |
| ---------- | ------------ | ----------------------------------- | --------------------- |
| **taskID** | [**string**] | Unique identifier of the async task | defaults to undefined |

### Return type

**CreateDatabaseTaskStatus**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details

| Status code | Description                                 | Response headers |
| ----------- | ------------------------------------------- | ---------------- |
| **200**     | Task status for a createDatabase operation. | -                |
| **404**     | The requested resource does not exist.      | -                |
| **500**     | An unexpected server-side error occurred.   | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getDeleteDatabaseTaskStatus**

> DeleteDatabaseTaskStatus getDeleteDatabaseTaskStatus()

Polls the status of a task initiated by `DELETE /databases/{dbID}/delete`. When `status` is `completed` the `result` field contains a `DeletionAcknowledgment`.

### Example

```typescript
import { TasksApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new TasksApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let taskID: string; //Unique identifier of the async task (default to undefined)

const { status, data } = await apiInstance.getDeleteDatabaseTaskStatus(
  dbID,
  taskID,
);
```

### Parameters

| Name       | Type         | Description                         | Notes                 |
| ---------- | ------------ | ----------------------------------- | --------------------- |
| **dbID**   | [**number**] | Unique identifier of the database   | defaults to undefined |
| **taskID** | [**string**] | Unique identifier of the async task | defaults to undefined |

### Return type

**DeleteDatabaseTaskStatus**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details

| Status code | Description                                 | Response headers |
| ----------- | ------------------------------------------- | ---------------- |
| **200**     | Task status for a deleteDatabase operation. | -                |
| **404**     | The requested resource does not exist.      | -                |
| **500**     | An unexpected server-side error occurred.   | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getExecuteSyncTaskStatus**

> ExecuteSyncTaskStatus getExecuteSyncTaskStatus()

Polls the status of a task initiated by `PUT /databases/{dbID}/sync`. When `status` is `completed` the `result` field contains a `SyncExecutionResult` with timing and outcome details.

### Example

```typescript
import { TasksApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new TasksApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let syncID: number; //Unique identifier of the sync rule (default to undefined)
let taskID: string; //Unique identifier of the async task (default to undefined)

const { status, data } = await apiInstance.getExecuteSyncTaskStatus(
  dbID,
  syncID,
  taskID,
);
```

### Parameters

| Name       | Type         | Description                         | Notes                 |
| ---------- | ------------ | ----------------------------------- | --------------------- |
| **dbID**   | [**number**] | Unique identifier of the database   | defaults to undefined |
| **syncID** | [**number**] | Unique identifier of the sync rule  | defaults to undefined |
| **taskID** | [**string**] | Unique identifier of the async task | defaults to undefined |

### Return type

**ExecuteSyncTaskStatus**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details

| Status code | Description                                   | Response headers |
| ----------- | --------------------------------------------- | ---------------- |
| **200**     | Task status for an executeSyncRule operation. | -                |
| **404**     | The requested resource does not exist.        | -                |
| **500**     | An unexpected server-side error occurred.     | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
