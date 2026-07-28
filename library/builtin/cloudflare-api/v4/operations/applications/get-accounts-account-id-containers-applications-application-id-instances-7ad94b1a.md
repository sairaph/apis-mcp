---
title: List container instances
page_id: operation-get-accounts-account-id-containers-applications-application-id-instances-0be7bef8
path: operations/applications
description: Lists container instances for an application, including control-plane details and any associated deployment/placement information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}/instances
operation_ids:
    - listContainerInstances
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List container instances

`GET /accounts/{account_id}/containers/applications/{application_id}/instances`

Operation ID: `listContainerInstances`

Lists container instances for an application, including control-plane details and any associated deployment/placement information.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "listContainerInstances", "summary": "List container instances", "description": "Lists container instances for an application, including control-plane details and any associated deployment/placement information.\n", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}, {"name": "per_page", "in": "query", "description": "Maximum number of instances to return per page. Defaults to all.", "schema": {"type": "integer", "maximum": 1000, "minimum": 1}}, {"name": "page_token", "in": "query", "description": "Opaque token from a previous response to retrieve the next page.", "schema": {"type": "string"}}], "responses": {"200": {"description": "List of container instances.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BasePaginatedResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_ListContainerInstances"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Container instance APIs are not enabled for this application or the request was invalid.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Response body when an Application is not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications", "Container Instances"], "x-api-token-group": ["Workers Containers Write", "Workers Containers Read"]}
```
