---
title: Set tags for an account-level resource
page_id: operation-put-accounts-account-id-tags-578020b0
path: operations/resource-tagging
description: Creates or updates tags for a specific account-level resource.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/tags
operation_ids:
    - tags-set
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set tags for an account-level resource

`PUT /accounts/{account_id}/tags`

Operation ID: `tags-set`

Creates or updates tags for a specific account-level resource.

## Definition

```yaml
{"operationId": "tags-set", "summary": "Set tags for an account-level resource", "description": "Creates or updates tags for a specific account-level resource.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_account_id"}}, {"$ref": "#/components/parameters/resource-tagging_if_match"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_set_tags_request_account_level"}}}}, "responses": {"200": {"description": "Set tags response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_tagged_resource_response_single"}}}}, "412": {"description": "Precondition failed. The resource has been modified since the provided ETag was obtained.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "4XX": {"description": "Set tags response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "Set tags response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "account-tags", "x-fern-sdk-method-name": "update", "x-stability": "beta"}
```
