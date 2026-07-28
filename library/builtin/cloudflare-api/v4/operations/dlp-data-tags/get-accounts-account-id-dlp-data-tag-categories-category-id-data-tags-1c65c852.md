---
title: Retrieve all data tags in a data tag category
page_id: operation-get-accounts-account-id-dlp-data-tag-categories-category-id-data-tags-bdbe0fb9
path: operations/dlp-data-tags
description: Lists data tags in a category.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags
operation_ids:
    - dlp-data-tags-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all data tags in a data tag category

`GET /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags`

Operation ID: `dlp-data-tags-list`

Lists data tags in a category.

## Definition

```yaml
{"operationId": "dlp-data-tags-list", "summary": "Retrieve all data tags in a data tag category", "description": "Lists data tags in a category.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "category_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Data tags read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTagArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tags read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tags"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
