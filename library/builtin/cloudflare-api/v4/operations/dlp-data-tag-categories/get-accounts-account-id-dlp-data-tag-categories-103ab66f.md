---
title: Retrieve all data tag categories in an account
page_id: operation-get-accounts-account-id-dlp-data-tag-categories-8328598b
path: operations/dlp-data-tag-categories
description: Lists data tag categories configured for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories
operation_ids:
    - dlp-data-tag-categories-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all data tag categories in an account

`GET /accounts/{account_id}/dlp/data_tag_categories`

Operation ID: `dlp-data-tag-categories-list`

Lists data tag categories configured for the account.

## Definition

```yaml
{"operationId": "dlp-data-tag-categories-list", "summary": "Retrieve all data tag categories in an account", "description": "Lists data tag categories configured for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Data tag category read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTagCategoryArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag category read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tag Categories"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
