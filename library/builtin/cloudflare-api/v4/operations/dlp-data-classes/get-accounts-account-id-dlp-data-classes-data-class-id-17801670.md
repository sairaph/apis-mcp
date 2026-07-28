---
title: Retrieve a specific data class
page_id: operation-get-accounts-account-id-dlp-data-classes-data-class-id-43f26534
path: operations/dlp-data-classes
description: Gets the configuration for a data class.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/data_classes/{data_class_id}
operation_ids:
    - dlp-data-classes-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a specific data class

`GET /accounts/{account_id}/dlp/data_classes/{data_class_id}`

Operation ID: `dlp-data-classes-read`

Gets the configuration for a data class.

## Definition

```yaml
{"operationId": "dlp-data-classes-read", "summary": "Retrieve a specific data class", "description": "Gets the configuration for a data class.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "data_class_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Data class read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataClass"}}, "type": "object"}]}}}}, "4XX": {"description": "Data class read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Classes"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
