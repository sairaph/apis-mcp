---
title: Retrieve a specific data tag category.
page_id: operation-get-accounts-account-id-dlp-data-tag-categories-category-id-eddc686e
path: operations/dlp-data-tag-categories
description: Gets a data tag category and its tags.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories/{category_id}
operation_ids:
    - dlp-data-tag-categories-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a specific data tag category.

`GET /accounts/{account_id}/dlp/data_tag_categories/{category_id}`

Operation ID: `dlp-data-tag-categories-read`

Gets a data tag category and its tags.

## Definition

```yaml
{"operationId": "dlp-data-tag-categories-read", "summary": "Retrieve a specific data tag category.", "description": "Gets a data tag category and its tags.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "category_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Data tag category read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTagCategory"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag category read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tag Categories"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
