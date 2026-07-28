---
title: Update the attributes of a single data class
page_id: operation-put-accounts-account-id-dlp-data-classes-data-class-id-a5f1e673
path: operations/dlp-data-classes
description: Updates the configuration for a data class.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/data_classes/{data_class_id}
operation_ids:
    - dlp-data-classes-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the attributes of a single data class

`PUT /accounts/{account_id}/dlp/data_classes/{data_class_id}`

Operation ID: `dlp-data-classes-update`

Updates the configuration for a data class.

## Definition

```yaml
{"operationId": "dlp-data-classes-update", "summary": "Update the attributes of a single data class", "description": "Updates the configuration for a data class.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "data_class_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Attributes of the data class to update.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_DataClassUpdate"}}}}, "responses": {"200": {"description": "Data class update was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataClass"}}, "type": "object"}]}}}}, "4XX": {"description": "Data class update failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Classes"], "x-api-token-group": ["Zero Trust Write"]}
```
