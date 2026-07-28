---
title: Update the attributes of a single data tag category.
page_id: operation-put-accounts-account-id-dlp-data-tag-categories-category-id-524c9167
path: operations/dlp-data-tag-categories
description: Updates a data tag category and its tags.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories/{category_id}
operation_ids:
    - dlp-data-tag-categories-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the attributes of a single data tag category.

`PUT /accounts/{account_id}/dlp/data_tag_categories/{category_id}`

Operation ID: `dlp-data-tag-categories-update`

Updates a data tag category and its tags.

## Definition

```yaml
{"operationId": "dlp-data-tag-categories-update", "summary": "Update the attributes of a single data tag category.", "description": "Updates a data tag category and its tags.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "category_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Attributes of the data tag category to update.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_DataTagCategoryUpdate"}}}}, "responses": {"200": {"description": "Data tag category update was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTagCategory"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag category update failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tag Categories"], "x-api-token-group": ["Zero Trust Write"]}
```
