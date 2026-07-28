---
title: Update the attributes of a single data tag.
page_id: operation-put-accounts-account-id-dlp-data-tag-categories-category-id-data-tags-ta-738bf6ab
path: operations/dlp-data-tags
description: Updates a data tag in a category.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags/{tag_id}
operation_ids:
    - dlp-data-tags-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the attributes of a single data tag.

`PUT /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags/{tag_id}`

Operation ID: `dlp-data-tags-update`

Updates a data tag in a category.

## Definition

```yaml
{"operationId": "dlp-data-tags-update", "summary": "Update the attributes of a single data tag.", "description": "Updates a data tag in a category.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "category_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "tag_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Attributes of the data tag to update.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_DataTagUpdate"}}}}, "responses": {"200": {"description": "Data tag update was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataTag"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag update failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tags"], "x-api-token-group": ["Zero Trust Write"]}
```
