---
title: Creates a new data tag category.
page_id: operation-post-accounts-account-id-dlp-data-tag-categories-ba2ffbdf
path: operations/dlp-data-tag-categories
description: Creates a data tag category, optionally from a template.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories
operation_ids:
    - dlp-data-tag-categories-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new data tag category.

`POST /accounts/{account_id}/dlp/data_tag_categories`

Operation ID: `dlp-data-tag-categories-create`

Creates a data tag category, optionally from a template.

## Definition

```yaml
{"operationId": "dlp-data-tag-categories-create", "summary": "Creates a new data tag category.", "description": "Creates a data tag category, optionally from a template.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Attributes of the new data tag category.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewDataTagCategory"}}}}, "responses": {"200": {"description": "Data tag category created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTagCategory"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag category creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tag Categories"], "x-api-token-group": ["Zero Trust Write"]}
```
