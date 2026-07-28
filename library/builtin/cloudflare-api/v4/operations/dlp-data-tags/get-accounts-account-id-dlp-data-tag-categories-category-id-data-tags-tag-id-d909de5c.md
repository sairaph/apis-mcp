---
title: Retrieve a specific data tag.
page_id: operation-get-accounts-account-id-dlp-data-tag-categories-category-id-data-tags-ta-0e1c4bc0
path: operations/dlp-data-tags
description: Gets a data tag from a category.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags/{tag_id}
operation_ids:
    - dlp-data-tags-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a specific data tag.

`GET /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags/{tag_id}`

Operation ID: `dlp-data-tags-read`

Gets a data tag from a category.

## Definition

```yaml
{"operationId": "dlp-data-tags-read", "summary": "Retrieve a specific data tag.", "description": "Gets a data tag from a category.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "category_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "tag_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Data tag read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTag"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tags"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
