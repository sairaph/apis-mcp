---
title: Set payload log settings
page_id: operation-put-accounts-account-id-dlp-payload-log-4164b85c
path: operations/dlp-settings
description: Enables or disables payload logging for DLP matches. When enabled, matched content is stored for review.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/payload_log
operation_ids:
    - dlp-payload-log-put
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set payload log settings

`PUT /accounts/{account_id}/dlp/payload_log`

Operation ID: `dlp-payload-log-put`

Enables or disables payload logging for DLP matches. When enabled, matched content is stored for review.

## Definition

```yaml
{"operationId": "dlp-payload-log-put", "summary": "Set payload log settings", "description": "Enables or disables payload logging for DLP matches. When enabled, matched content is stored for review.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "New payload log settings.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_PayloadLogSettingUpdateLegacy"}}}}, "responses": {"200": {"description": "Payload log settings.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_PayloadLogSetting"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to set payload log settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Settings"], "x-api-token-group": ["Zero Trust Write"]}
```
