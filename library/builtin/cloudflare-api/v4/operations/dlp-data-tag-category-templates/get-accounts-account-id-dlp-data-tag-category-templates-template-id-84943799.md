---
title: Retrieve a specific data tag category template.
page_id: operation-get-accounts-account-id-dlp-data-tag-category-templates-template-id-6ecfe874
path: operations/dlp-data-tag-category-templates
description: Gets an available data tag category template.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_category_templates/{template_id}
operation_ids:
    - dlp-data-tag-category-template-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a specific data tag category template.

`GET /accounts/{account_id}/dlp/data_tag_category_templates/{template_id}`

Operation ID: `dlp-data-tag-category-template-read`

Gets an available data tag category template.

## Definition

```yaml
{"operationId": "dlp-data-tag-category-template-read", "summary": "Retrieve a specific data tag category template.", "description": "Gets an available data tag category template.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "template_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Data tag category template read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTagCategoryTemplate"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag category template read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tag Category Templates"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-stainless-skip": ["terraform"]}
```
