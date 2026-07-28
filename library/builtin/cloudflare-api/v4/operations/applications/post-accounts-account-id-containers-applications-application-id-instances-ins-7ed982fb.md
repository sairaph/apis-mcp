---
title: Execute a command in a container instance
page_id: operation-post-accounts-account-id-containers-applications-application-id-instance-627b9405
path: operations/applications
description: Executes a command in a running container instance and returns its buffered standard output, standard error, and exit code.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}/instances/{instance_id}/exec
operation_ids:
    - containerInstanceExec
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Execute a command in a container instance

`POST /accounts/{account_id}/containers/applications/{application_id}/instances/{instance_id}/exec`

Operation ID: `containerInstanceExec`

Executes a command in a running container instance and returns its buffered standard output, standard error, and exit code.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "containerInstanceExec", "summary": "Execute a command in a container instance", "description": "Executes a command in a running container instance and returns its buffered standard output, standard error, and exit code.\n", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}, {"name": "instance_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ContainerInstanceID"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_ContainerInstanceExecRequestBody"}}}}, "responses": {"200": {"description": "Command execution completed.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_ContainerInstanceExecResult"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Container instance APIs are not enabled for this application or the request was invalid.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Container instance not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "503": {"description": "The container instance is not running.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications", "Container Instances"]}
```
