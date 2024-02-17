# \TaskAPI

All URIs are relative to *http://localhost:8080/todoapp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksCreate**](TaskAPI.md#TasksCreate) | **Post** /task | Create a new task
[**TasksDelete**](TaskAPI.md#TasksDelete) | **Delete** /task/{taskId} | Delete an existing task
[**TasksGetAll**](TaskAPI.md#TasksGetAll) | **Get** /task | Get the list of all tasks
[**TasksRead**](TaskAPI.md#TasksRead) | **Get** /task/{taskId} | Get a single task based on its id
[**TasksUpdate**](TaskAPI.md#TasksUpdate) | **Put** /task/{taskId} | Update an existing task



## TasksCreate

> Task TasksCreate(ctx).Task(task).Execute()

Create a new task

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    task := *openapiclient.NewTask("My important task") // Task |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.TasksCreate(context.Background()).Task(task).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksCreate`: Task
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **task** | [**Task**](Task.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksDelete

> TasksDelete(ctx, taskId).Execute()

Delete an existing task

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    taskId := "e1cb23d0-6cbe-4a29-b586-bfa424bc93fd" // string | The id of the task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaskAPI.TasksDelete(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The id of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksGetAll

> []Task TasksGetAll(ctx).Execute()

Get the list of all tasks

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.TasksGetAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksGetAll`: []Task
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksGetAllRequest struct via the builder pattern


### Return type

[**[]Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRead

> Task TasksRead(ctx, taskId).Execute()

Get a single task based on its id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    taskId := "e1cb23d0-6cbe-4a29-b586-bfa424bc93fd" // string | The id of the task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.TasksRead(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksRead`: Task
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The id of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksUpdate

> Task TasksUpdate(ctx, taskId).Task(task).Execute()

Update an existing task

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    taskId := "e1cb23d0-6cbe-4a29-b586-bfa424bc93fd" // string | The id of the task
    task := *openapiclient.NewTask("My important task") // Task |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.TasksUpdate(context.Background(), taskId).Task(task).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.TasksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksUpdate`: Task
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.TasksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The id of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | [**Task**](Task.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

