# DefaultApi

All URIs are relative to *https://mezcala.net/api*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**addDatabase**](#adddatabase) | **POST** /databases/add | Register an existing external database|
|[**checkHealth**](#checkhealth) | **GET** /health | Check Mezcala health|
|[**createDatabase**](#createdatabase) | **POST** /databases/create | Create a new managed database|
|[**createSyncRule**](#createsyncrule) | **POST** /databases/{dbID}/sync/init | Create a new sync rule for a database|
|[**deleteDatabase**](#deletedatabase) | **DELETE** /databases/{dbID}/delete | Delete a database by ID|
|[**deleteSyncRule**](#deletesyncrule) | **DELETE** /databases/{dbID}/sync/{syncID}/delete | Delete a sync rule|
|[**executeSyncRule**](#executesyncrule) | **PUT** /databases/{dbID}/sync | Trigger execution of a sync rule|
|[**getCreateDatabaseTaskStatus**](#getcreatedatabasetaskstatus) | **GET** /databases/create/tasks/{taskID} | Get the status of a createDatabase task|
|[**getDatabaseConnectionUrl**](#getdatabaseconnectionurl) | **GET** /databases/{dbID}/connection-url | Get the connection URL for a database|
|[**getDeleteDatabaseTaskStatus**](#getdeletedatabasetaskstatus) | **GET** /databases/{dbID}/delete/tasks/{taskID} | Get the status of a deleteDatabase task|
|[**getExecuteSyncTaskStatus**](#getexecutesynctaskstatus) | **GET** /databases/{dbID}/sync/{syncID}/tasks/{taskID} | Get the status of an executeSyncRule task|
|[**getSupportedEngines**](#getsupportedengines) | **GET** /databases/supported | List all supported database engines and versions|
|[**updateSyncRule**](#updatesyncrule) | **PUT** /databases/{dbID}/sync/{syncID}/update | Update an existing sync rule|

# **addDatabase**
> DatabaseAcknowledgment addDatabase(databaseConfig)

Synchronously registers an external database connection and returns an acknowledgment with the assigned database ID. 

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    DatabaseConfig
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let databaseConfig: DatabaseConfig; //Payload for registering or creating a database connection.

const { status, data } = await apiInstance.addDatabase(
    databaseConfig
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **databaseConfig** | **DatabaseConfig**| Payload for registering or creating a database connection. | |


### Return type

**DatabaseAcknowledgment**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | The database was registered or created successfully. |  -  |
|**400** | The request is malformed or fails validation. |  -  |
|**409** | The resource already exists or violates a uniqueness constraint. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **checkHealth**
> CheckHealth200Response checkHealth()


### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

const { status, data } = await apiInstance.checkHealth();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**CheckHealth200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Service is reachable and reports its current status. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **createDatabase**
> TaskReference createDatabase(databaseConfig)

Initiates database creation asynchronously. Returns a task ID that can be polled via `GET /databases/create/tasks/{taskID}`. On completion the task result contains an acknowledgment with the new database ID. 

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    DatabaseConfig
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let databaseConfig: DatabaseConfig; //Payload for registering or creating a database connection.

const { status, data } = await apiInstance.createDatabase(
    databaseConfig
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **databaseConfig** | **DatabaseConfig**| Payload for registering or creating a database connection. | |


### Return type

**TaskReference**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**202** | The operation has been accepted and is running asynchronously. Poll the corresponding task status endpoint to track progress.  |  -  |
|**400** | The request is malformed or fails validation. |  -  |
|**409** | The resource already exists or violates a uniqueness constraint. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **createSyncRule**
> CreateSyncRule201Response createSyncRule(syncRuleConfig)

Synchronously creates a sync rule and returns it with its assigned ID. 

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    SyncRuleConfig
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let syncRuleConfig: SyncRuleConfig; //Payload for creating or updating a sync rule.

const { status, data } = await apiInstance.createSyncRule(
    dbID,
    syncRuleConfig
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **syncRuleConfig** | **SyncRuleConfig**| Payload for creating or updating a sync rule. | |
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined|


### Return type

**CreateSyncRule201Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | The created or updated sync rule. |  -  |
|**400** | The request is malformed or fails validation. |  -  |
|**404** | The requested resource does not exist. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteDatabase**
> TaskReference deleteDatabase()

Initiates database deletion asynchronously. Returns a task ID that can be polled via `GET /databases/{dbID}/delete/tasks/{taskID}`. On completion the task result contains a confirmation message. 

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)

const { status, data } = await apiInstance.deleteDatabase(
    dbID
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined|


### Return type

**TaskReference**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**202** | The operation has been accepted and is running asynchronously. Poll the corresponding task status endpoint to track progress.  |  -  |
|**404** | The requested resource does not exist. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteSyncRule**
> deleteSyncRule()

Synchronously deletes a sync rule.

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let syncID: number; //Unique identifier of the sync rule (default to undefined)

const { status, data } = await apiInstance.deleteSyncRule(
    dbID,
    syncID
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined|
| **syncID** | [**number**] | Unique identifier of the sync rule | defaults to undefined|


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**204** | Sync rule deleted successfully. |  -  |
|**404** | The requested resource does not exist. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **executeSyncRule**
> TaskReference executeSyncRule(executeSyncRuleRequest)

Triggers a sync rule run asynchronously. Returns a task ID that can be polled via `GET /databases/{dbID}/sync/{syncID}/tasks/{taskID}`. On completion the task result contains the sync execution summary. 

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    ExecuteSyncRuleRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let executeSyncRuleRequest: ExecuteSyncRuleRequest; //Identifies which sync rule to execute.

const { status, data } = await apiInstance.executeSyncRule(
    dbID,
    executeSyncRuleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **executeSyncRuleRequest** | **ExecuteSyncRuleRequest**| Identifies which sync rule to execute. | |
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined|


### Return type

**TaskReference**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**202** | The operation has been accepted and is running asynchronously. Poll the corresponding task status endpoint to track progress.  |  -  |
|**400** | The request is malformed or fails validation. |  -  |
|**404** | The requested resource does not exist. |  -  |
|**422** | The request is well-formed but contains semantic errors. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCreateDatabaseTaskStatus**
> CreateDatabaseTaskStatus getCreateDatabaseTaskStatus()

Polls the status of a task initiated by `POST /databases/create`. When `status` is `completed` the `result` field contains a `DatabaseAcknowledgment` with the new database ID. 

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let taskID: string; //Unique identifier of the async task (default to undefined)

const { status, data } = await apiInstance.getCreateDatabaseTaskStatus(
    taskID
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **taskID** | [**string**] | Unique identifier of the async task | defaults to undefined|


### Return type

**CreateDatabaseTaskStatus**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Task status for a createDatabase operation. |  -  |
|**404** | The requested resource does not exist. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getDatabaseConnectionUrl**
> GetDatabaseConnectionUrl200Response getDatabaseConnectionUrl()

Returns the fully-formed connection URL for the specified database, including credentials. Handle with care — treat the response as a secret. 

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)

const { status, data } = await apiInstance.getDatabaseConnectionUrl(
    dbID
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined|


### Return type

**GetDatabaseConnectionUrl200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | The connection URL for the specified database. |  -  |
|**404** | The requested resource does not exist. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getDeleteDatabaseTaskStatus**
> DeleteDatabaseTaskStatus getDeleteDatabaseTaskStatus()

Polls the status of a task initiated by `DELETE /databases/{dbID}/delete`. When `status` is `completed` the `result` field contains a `DeletionAcknowledgment`. 

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let taskID: string; //Unique identifier of the async task (default to undefined)

const { status, data } = await apiInstance.getDeleteDatabaseTaskStatus(
    dbID,
    taskID
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined|
| **taskID** | [**string**] | Unique identifier of the async task | defaults to undefined|


### Return type

**DeleteDatabaseTaskStatus**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Task status for a deleteDatabase operation. |  -  |
|**404** | The requested resource does not exist. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getExecuteSyncTaskStatus**
> ExecuteSyncTaskStatus getExecuteSyncTaskStatus()

Polls the status of a task initiated by `PUT /databases/{dbID}/sync`. When `status` is `completed` the `result` field contains a `SyncExecutionResult` with timing and outcome details. 

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let syncID: number; //Unique identifier of the sync rule (default to undefined)
let taskID: string; //Unique identifier of the async task (default to undefined)

const { status, data } = await apiInstance.getExecuteSyncTaskStatus(
    dbID,
    syncID,
    taskID
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined|
| **syncID** | [**number**] | Unique identifier of the sync rule | defaults to undefined|
| **taskID** | [**string**] | Unique identifier of the async task | defaults to undefined|


### Return type

**ExecuteSyncTaskStatus**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Task status for an executeSyncRule operation. |  -  |
|**404** | The requested resource does not exist. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getSupportedEngines**
> Array<GetSupportedEngines200ResponseInner> getSupportedEngines()


### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

const { status, data } = await apiInstance.getSupportedEngines();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**Array<GetSupportedEngines200ResponseInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Array of supported database engines and their available versions. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateSyncRule**
> CreateSyncRule201Response updateSyncRule(syncRuleConfig)

Synchronously updates a sync rule and returns the updated record. 

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    SyncRuleConfig
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)
let syncID: number; //Unique identifier of the sync rule (default to undefined)
let syncRuleConfig: SyncRuleConfig; //Payload for creating or updating a sync rule.

const { status, data } = await apiInstance.updateSyncRule(
    dbID,
    syncID,
    syncRuleConfig
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **syncRuleConfig** | **SyncRuleConfig**| Payload for creating or updating a sync rule. | |
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined|
| **syncID** | [**number**] | Unique identifier of the sync rule | defaults to undefined|


### Return type

**CreateSyncRule201Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | The created or updated sync rule. |  -  |
|**400** | The request is malformed or fails validation. |  -  |
|**404** | The requested resource does not exist. |  -  |
|**500** | An unexpected server-side error occurred. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

