---
title: Delete a single data tag.
page_id: operation-delete-accounts-account-id-dlp-data-tag-categories-category-id-data-tags-5b10dfcc
path: operations/dlp-data-tags
description: Deletes a data tag from a category.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags/{tag_id}
operation_ids:
    - dlp-data-tags-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a single data tag.

`DELETE /accounts/{account_id}/dlp/data_tag_categories/{category_id}/data_tags/{tag_id}`

Operation ID: `dlp-data-tags-delete`

Deletes a data tag from a category.

## Definition

```yaml
{"operationId": "dlp-data-tags-delete", "summary": "Delete a single data tag.", "description": "Deletes a data tag from a category.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "category_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "tag_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Data tag delete was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Data tag delete failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Tags"], "x-api-token-group": ["Zero Trust Write"]}
```
