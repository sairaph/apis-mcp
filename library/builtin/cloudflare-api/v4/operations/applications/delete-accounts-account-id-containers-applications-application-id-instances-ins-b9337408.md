---
title: Delete a container instance
page_id: operation-delete-accounts-account-id-containers-applications-application-id-instan-652c945f
path: operations/applications
description: Stops the backing Durable Object container by sending SIGKILL. The instance remains visible until normal runtime lifecycle processing marks it asleep and eventually prunes it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}/instances/{instance_id}
operation_ids:
    - deleteContainerInstance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a container instance

`DELETE /accounts/{account_id}/containers/applications/{application_id}/instances/{instance_id}`

Operation ID: `deleteContainerInstance`

Stops the backing Durable Object container by sending SIGKILL. The instance remains visible until normal runtime lifecycle processing marks it asleep and eventually prunes it.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "deleteContainerInstance", "summary": "Delete a container instance", "description": "Stops the backing Durable Object container by sending SIGKILL. The instance remains visible until normal runtime lifecycle processing marks it asleep and eventually prunes it.\n", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}, {"name": "instance_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ContainerInstanceID"}}], "responses": {"200": {"description": "Container instance successfully stopped, or was already asleep.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_EmptyResponse"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Container instance APIs are not enabled for this application or the request was invalid.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Container instance not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications", "Container Instances"], "x-api-token-group": ["Workers Containers Write"]}
```
