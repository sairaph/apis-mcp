---
title: Creates a new data tag.
page_id: operation-post-accounts-account-id-dlp-data-tag-categories-category-id-data-tags-c5c3435c
path: operations/dlp-data-tags
description: Creates a data tag in a category.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags
operation_ids:
    - dlp-data-tags-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new data tag.

`POST /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags`

Operation ID: `dlp-data-tags-create`

Creates a data tag in a category.

## Definition

```yaml
{"operationId": "dlp-data-tags-create", "summary": "Creates a new data tag.", "description": "Creates a data tag in a category.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "category_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Attributes of the new data tag.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewDataTag"}}}}, "responses": {"200": {"description": "Data tag created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTag"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tags"], "x-api-token-group": ["Zero Trust Write"]}
```
