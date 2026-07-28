---
title: Get payload log settings
page_id: operation-get-accounts-account-id-dlp-payload-log-f934e3a9
path: operations/dlp-settings
description: Gets the current payload logging configuration for DLP, showing whether matched content is being logged.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/payload_log
operation_ids:
    - dlp-payload-log-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get payload log settings

`GET /accounts/{account_id}/dlp/payload_log`

Operation ID: `dlp-payload-log-get`

Gets the current payload logging configuration for DLP, showing whether matched content is being logged.

## Definition

```yaml
{"operationId": "dlp-payload-log-get", "summary": "Get payload log settings", "description": "Gets the current payload logging configuration for DLP, showing whether matched content is being logged.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Payload log settings.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_PayloadLogSetting"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to get payload log settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Settings"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
