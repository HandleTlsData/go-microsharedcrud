# \EntityApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntityDelete**](EntityApi.md#EntityDelete) | **Delete** /entity/delete/{entName} | Delete a single Entity by given Name
[**EntityHandler**](EntityApi.md#EntityHandler) | **Get** /entity/{entityName} | Returns a single Entity by given Name
[**EntityStore**](EntityApi.md#EntityStore) | **Post** /entity | Add a new Entity if such Name is not exists. Update otherwise


# **EntityDelete**
> []Entity EntityDelete(ctx, entName)
Delete a single Entity by given Name

Delete full entity by it's Name field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entName** | **string**| Name of entity that need to be returned | 

### Return type

[**[]Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntityHandler**
> []Entity EntityHandler(ctx, entityName)
Returns a single Entity by given Name

Finds full entity by it's Name field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityName** | **string**| Name of entity that need to be returned | 

### Return type

[**[]Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntityStore**
> EntityStore(ctx, body)
Add a new Entity if such Name is not exists. Update otherwise



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Entity**](Entity.md)| Entity object that needs to be stored | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

