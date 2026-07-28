---
title: Get DLP account-level settings.
page_id: operation-get-accounts-account-id-dlp-settings-8e9b7326
path: operations/dlp-settings
description: Gets the account-level DLP settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/settings
operation_ids:
    - dlp-settings-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get DLP account-level settings.

`GET /accounts/{account_id}/dlp/settings`

Operation ID: `dlp-settings-get`

Gets the account-level DLP settings.

## Definition

```yaml
{"operationId": "dlp-settings-get", "summary": "Get DLP account-level settings.", "description": "Gets the account-level DLP settings.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "DLP settings.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DlpSettings"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to get DLP settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Settings"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
