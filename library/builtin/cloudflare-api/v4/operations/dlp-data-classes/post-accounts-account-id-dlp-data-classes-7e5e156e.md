---
title: Creates a new data class
page_id: operation-post-accounts-account-id-dlp-data-classes-255e01f3
path: operations/dlp-data-classes
description: Creates a data class for use in DLP profiles.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/data_classes
operation_ids:
    - dlp-data-classes-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new data class

`POST /accounts/{account_id}/dlp/data_classes`

Operation ID: `dlp-data-classes-create`

Creates a data class for use in DLP profiles.

## Definition

```yaml
{"operationId": "dlp-data-classes-create", "summary": "Creates a new data class", "description": "Creates a data class for use in DLP profiles.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Attributes of the new data class.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewDataClass"}}}}, "responses": {"200": {"description": "Data class created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataClass"}}, "type": "object"}]}}}}, "4XX": {"description": "Data class creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Classes"], "x-api-token-group": ["Zero Trust Write"]}
```
