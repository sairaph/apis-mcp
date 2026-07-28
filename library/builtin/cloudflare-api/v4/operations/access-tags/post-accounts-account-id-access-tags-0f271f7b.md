---
title: Create a tag
page_id: operation-post-accounts-account-id-access-tags-d8d4b975
path: operations/access-tags
description: Create a tag
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/tags
operation_ids:
    - access-tags-create-tag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a tag

`POST /accounts/{account_id}/access/tags`

Operation ID: `access-tags-create-tag`

Create a tag

## Definition

```yaml
{"operationId": "access-tags-create-tag", "summary": "Create a tag", "description": "Create a tag", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"properties": {"name": {"$ref": "#/components/schemas/access_name-13"}}}}}}, "responses": {"201": {"description": "Create a tag response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-14"}}}}, "4XX": {"description": "Create a tag response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access tags"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.tags", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
