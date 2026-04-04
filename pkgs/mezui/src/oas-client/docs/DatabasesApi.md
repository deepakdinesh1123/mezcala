# DatabasesApi

All URIs are relative to _localhost:7007_

| Method                                                    | HTTP request                             | Description                                      |
| --------------------------------------------------------- | ---------------------------------------- | ------------------------------------------------ |
| [**addDatabase**](#adddatabase)                           | **POST** /databases/add                  | Register an existing external database           |
| [**createDatabase**](#createdatabase)                     | **POST** /databases/create               | Create a new managed database                    |
| [**deleteDatabase**](#deletedatabase)                     | **DELETE** /databases/{dbID}/delete      | Delete a database by ID                          |
| [**getDatabaseConnectionUrl**](#getdatabaseconnectionurl) | **GET** /databases/{dbID}/connection-url | Get the connection URL for a database            |
| [**getSupportedEngines**](#getsupportedengines)           | **GET** /databases/supported             | List all supported database engines and versions |

# **addDatabase**

> DatabaseAcknowledgment addDatabase(databaseConfig)

Synchronously registers an external database connection and returns an acknowledgment with the assigned database ID.

### Example

```typescript
import { DatabasesApi, Configuration, DatabaseConfig } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabasesApi(configuration);

let databaseConfig: DatabaseConfig; //Payload for registering or creating a database connection.

const { status, data } = await apiInstance.addDatabase(databaseConfig);
```

### Parameters

| Name               | Type               | Description                                                | Notes |
| ------------------ | ------------------ | ---------------------------------------------------------- | ----- |
| **databaseConfig** | **DatabaseConfig** | Payload for registering or creating a database connection. |       |

### Return type

**DatabaseAcknowledgment**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details

| Status code | Description                                                      | Response headers |
| ----------- | ---------------------------------------------------------------- | ---------------- |
| **201**     | The database was registered or created successfully.             | -                |
| **400**     | The request is malformed or fails validation.                    | -                |
| **409**     | The resource already exists or violates a uniqueness constraint. | -                |
| **500**     | An unexpected server-side error occurred.                        | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **createDatabase**

> TaskReference createDatabase(databaseConfig)

Initiates database creation asynchronously. Returns a task ID that can be polled via `GET /databases/create/tasks/{taskID}`. On completion the task result contains an acknowledgment with the new database ID.

### Example

```typescript
import { DatabasesApi, Configuration, DatabaseConfig } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabasesApi(configuration);

let databaseConfig: DatabaseConfig; //Payload for registering or creating a database connection.

const { status, data } = await apiInstance.createDatabase(databaseConfig);
```

### Parameters

| Name               | Type               | Description                                                | Notes |
| ------------------ | ------------------ | ---------------------------------------------------------- | ----- |
| **databaseConfig** | **DatabaseConfig** | Payload for registering or creating a database connection. |       |

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
| **409**     | The resource already exists or violates a uniqueness constraint.                                                              | -                |
| **500**     | An unexpected server-side error occurred.                                                                                     | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteDatabase**

> TaskReference deleteDatabase()

Initiates database deletion asynchronously. Returns a task ID that can be polled via `GET /databases/{dbID}/delete/tasks/{taskID}`. On completion the task result contains a confirmation message.

### Example

```typescript
import { DatabasesApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabasesApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)

const { status, data } = await apiInstance.deleteDatabase(dbID);
```

### Parameters

| Name     | Type         | Description                       | Notes                 |
| -------- | ------------ | --------------------------------- | --------------------- |
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined |

### Return type

**TaskReference**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details

| Status code | Description                                                                                                                   | Response headers |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------- | ---------------- |
| **202**     | The operation has been accepted and is running asynchronously. Poll the corresponding task status endpoint to track progress. | -                |
| **404**     | The requested resource does not exist.                                                                                        | -                |
| **500**     | An unexpected server-side error occurred.                                                                                     | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getDatabaseConnectionUrl**

> GetDatabaseConnectionUrl200Response getDatabaseConnectionUrl()

Returns the fully-formed connection URL for the specified database, including credentials. Handle with care — treat the response as a secret.

### Example

```typescript
import { DatabasesApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabasesApi(configuration);

let dbID: number; //Unique identifier of the database (default to undefined)

const { status, data } = await apiInstance.getDatabaseConnectionUrl(dbID);
```

### Parameters

| Name     | Type         | Description                       | Notes                 |
| -------- | ------------ | --------------------------------- | --------------------- |
| **dbID** | [**number**] | Unique identifier of the database | defaults to undefined |

### Return type

**GetDatabaseConnectionUrl200Response**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details

| Status code | Description                                    | Response headers |
| ----------- | ---------------------------------------------- | ---------------- |
| **200**     | The connection URL for the specified database. | -                |
| **404**     | The requested resource does not exist.         | -                |
| **500**     | An unexpected server-side error occurred.      | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getSupportedEngines**

> Array<GetSupportedEngines200ResponseInner> getSupportedEngines()

### Example

```typescript
import { DatabasesApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new DatabasesApi(configuration);

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

| Status code | Description                                                       | Response headers |
| ----------- | ----------------------------------------------------------------- | ---------------- |
| **200**     | Array of supported database engines and their available versions. | -                |
| **500**     | An unexpected server-side error occurred.                         | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
