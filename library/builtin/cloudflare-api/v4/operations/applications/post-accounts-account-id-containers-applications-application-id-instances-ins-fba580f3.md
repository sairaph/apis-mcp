---
title: Proxy a request to a container instance
page_id: operation-post-accounts-account-id-containers-applications-application-id-instance-46c1bd20
path: operations/applications
description: Forwards an HTTP request to a running container instance by invoking fetch on the backing Durable Object. The container must be running. The container response status, headers, and text body are returned in the API response body.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}/instances/{instance_id}/fetch
operation_ids:
    - containerInstanceFetch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Proxy a request to a container instance

`POST /accounts/{account_id}/containers/applications/{application_id}/instances/{instance_id}/fetch`

Operation ID: `containerInstanceFetch`

Forwards an HTTP request to a running container instance by invoking fetch on the backing Durable Object. The container must be running. The container response status, headers, and text body are returned in the API response body.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "containerInstanceFetch", "summary": "Proxy a request to a container instance", "description": "Forwards an HTTP request to a running container instance by invoking fetch on the backing Durable Object. The container must be running. The container response status, headers, and text body are returned in the API response body.\n", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}, {"name": "instance_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ContainerInstanceID"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_ContainerInstanceFetchRequestBody"}}}}, "responses": {"200": {"description": "Response returned by the container instance.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_ContainerInstanceFetchResult"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Container instance APIs are not enabled for this application or the request was invalid.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Container instance not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications", "Container Instances"]}
```
