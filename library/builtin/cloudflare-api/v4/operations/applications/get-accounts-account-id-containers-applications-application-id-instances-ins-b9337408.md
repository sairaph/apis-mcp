---
title: Get a container instance
page_id: operation-get-accounts-account-id-containers-applications-application-id-instances-bbe755aa
path: operations/applications
description: Returns control-plane details for a single container instance and any associated deployment/placement information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}/instances/{instance_id}
operation_ids:
    - getContainerInstance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a container instance

`GET /accounts/{account_id}/containers/applications/{application_id}/instances/{instance_id}`

Operation ID: `getContainerInstance`

Returns control-plane details for a single container instance and any associated deployment/placement information.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "getContainerInstance", "summary": "Get a container instance", "description": "Returns control-plane details for a single container instance and any associated deployment/placement information.\n", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}, {"name": "instance_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ContainerInstanceID"}}], "responses": {"200": {"description": "Container instance details.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_ContainerInstance"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Container instance APIs are not enabled for this application or the request was invalid.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Container instance not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications", "Container Instances"], "x-api-token-group": ["Workers Containers Write", "Workers Containers Read"]}
```
