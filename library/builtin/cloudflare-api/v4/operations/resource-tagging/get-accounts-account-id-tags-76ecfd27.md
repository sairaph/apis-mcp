---
title: Get tags for an account-level resource
page_id: operation-get-accounts-account-id-tags-44af4e84
path: operations/resource-tagging
description: Retrieves tags for a specific account-level resource.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tags
operation_ids:
    - tags-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get tags for an account-level resource

`GET /accounts/{account_id}/tags`

Operation ID: `tags-get`

Retrieves tags for a specific account-level resource.

## Definition

```yaml
{"operationId": "tags-get", "summary": "Get tags for an account-level resource", "description": "Retrieves tags for a specific account-level resource.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_account_id"}}, {"name": "resource_id", "in": "query", "description": "The ID of the resource to retrieve tags for.", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_resource_id"}}, {"name": "resource_type", "in": "query", "description": "The type of the resource.", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_account_resource_type"}}, {"name": "worker_id", "in": "query", "description": "Worker identifier. Required for worker_version resources.", "schema": {"$ref": "#/components/schemas/resource-tagging_worker_id"}}], "responses": {"200": {"description": "Get tags for single resource response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_tagged_resource_response_single"}}}}, "4XX": {"description": "Get tags for single resource response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "Get tags for single resource response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "account-tags", "x-fern-sdk-method-name": "get", "x-stability": "beta"}
```
