---
title: Update a tag
page_id: operation-put-accounts-account-id-access-tags-tag-name-5d35eaf0
path: operations/access-tags
description: Update a tag
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/tags/{tag_name}
operation_ids:
    - access-tags-update-a-tag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a tag

`PUT /accounts/{account_id}/access/tags/{tag_name}`

Operation ID: `access-tags-update-a-tag`

Update a tag

## Definition

```yaml
{"operationId": "access-tags-update-a-tag", "summary": "Update a tag", "description": "Update a tag", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "tag_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_name-13"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_tag_without_app_count"}}}}, "responses": {"200": {"description": "Update a tag response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-14"}}}}, "4XX": {"description": "Update a tag response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access tags"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.tags", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
