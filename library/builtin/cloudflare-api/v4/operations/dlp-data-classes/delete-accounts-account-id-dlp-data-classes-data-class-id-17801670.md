---
title: Delete a single data class
page_id: operation-delete-accounts-account-id-dlp-data-classes-data-class-id-64c5eb89
path: operations/dlp-data-classes
description: Deletes a data class from the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/data_classes/{data_class_id}
operation_ids:
    - dlp-data-classes-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a single data class

`DELETE /accounts/{account_id}/dlp/data_classes/{data_class_id}`

Operation ID: `dlp-data-classes-delete`

Deletes a data class from the account.

## Definition

```yaml
{"operationId": "dlp-data-classes-delete", "summary": "Delete a single data class", "description": "Deletes a data class from the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "data_class_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Data class delete was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Data class delete failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Classes"], "x-api-token-group": ["Zero Trust Write"]}
```
