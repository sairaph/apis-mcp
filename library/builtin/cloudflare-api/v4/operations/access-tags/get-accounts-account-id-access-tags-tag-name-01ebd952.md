---
title: Get a tag
page_id: operation-get-accounts-account-id-access-tags-tag-name-c4830973
path: operations/access-tags
description: Get a tag
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/tags/{tag_name}
operation_ids:
    - access-tags-get-a-tag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a tag

`GET /accounts/{account_id}/access/tags/{tag_name}`

Operation ID: `access-tags-get-a-tag`

Get a tag

## Definition

```yaml
{"operationId": "access-tags-get-a-tag", "summary": "Get a tag", "description": "Get a tag", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "tag_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_name-13"}}], "responses": {"200": {"description": "Get a tag response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-14"}}}}, "4XX": {"description": "Get a tag response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access tags"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.tags", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
