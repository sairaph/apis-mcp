---
title: Retrieve all data tag category templates in an account
page_id: operation-get-accounts-account-id-dlp-data-tag-category-templates-76aa9bc5
path: operations/dlp-data-tag-category-templates
description: Lists available data tag category templates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_category_templates
operation_ids:
    - dlp-data-tag-category-templates-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all data tag category templates in an account

`GET /accounts/{account_id}/dlp/data_tag_category_templates`

Operation ID: `dlp-data-tag-category-templates-list`

Lists available data tag category templates.

## Definition

```yaml
{"operationId": "dlp-data-tag-category-templates-list", "summary": "Retrieve all data tag category templates in an account", "description": "Lists available data tag category templates.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Data tag category template read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTagCategoryTemplateArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag category template read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tag Category Templates"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-stainless-skip": ["terraform"]}
```
