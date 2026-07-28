---
title: Create a container instance
page_id: operation-post-accounts-account-id-containers-applications-application-id-instance-14ec79ff
path: operations/applications
description: Creates a new container instance within an application. The instance's container is started immediately by invoking startAndWaitForPorts on the backing Durable Object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}/instances
operation_ids:
    - createContainerInstance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a container instance

`POST /accounts/{account_id}/containers/applications/{application_id}/instances`

Operation ID: `createContainerInstance`

Creates a new container instance within an application. The instance's container is started immediately by invoking startAndWaitForPorts on the backing Durable Object.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "createContainerInstance", "summary": "Create a container instance", "description": "Creates a new container instance within an application. The instance's container is started immediately by invoking startAndWaitForPorts on the backing Durable Object.\n", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_CreateContainerInstanceRequestBody"}}}}, "responses": {"201": {"description": "Successfully created the container instance.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_ContainerInstance"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Container instance APIs are not enabled for this application or the request was invalid.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Response body when an Application is not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications", "Container Instances"]}
```
