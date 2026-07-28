---
title: Delete tags from an account-level resource
page_id: operation-delete-accounts-account-id-tags-fff87f20
path: operations/resource-tagging
description: Removes all tags from a specific account-level resource.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/tags
operation_ids:
    - tags-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete tags from an account-level resource

`DELETE /accounts/{account_id}/tags`

Operation ID: `tags-delete`

Removes all tags from a specific account-level resource.

## Definition

```yaml
{"operationId": "tags-delete", "summary": "Delete tags from an account-level resource", "description": "Removes all tags from a specific account-level resource.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_account_id"}}, {"$ref": "#/components/parameters/resource-tagging_if_match"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_delete_tags_request_account_level"}}}}, "responses": {"204": {"description": "Tags successfully deleted (no content)."}, "412": {"description": "Precondition failed. The resource has been modified since the provided ETag was obtained.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "4XX": {"description": "Delete tags response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "Delete tags response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "account-tags", "x-fern-sdk-method-name": "delete", "x-stability": "beta"}
```
