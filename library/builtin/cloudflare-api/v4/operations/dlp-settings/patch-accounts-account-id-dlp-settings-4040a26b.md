---
title: Partially update DLP account-level settings.
page_id: operation-patch-accounts-account-id-dlp-settings-6e5ed39a
path: operations/dlp-settings
description: Missing fields keep their existing values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/dlp/settings
operation_ids:
    - dlp-settings-edit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Partially update DLP account-level settings.

`PATCH /accounts/{account_id}/dlp/settings`

Operation ID: `dlp-settings-edit`

Missing fields keep their existing values.

## Definition

```yaml
{"operationId": "dlp-settings-edit", "summary": "Partially update DLP account-level settings.", "description": "Missing fields keep their existing values.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "DLP settings fields to update.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_DlpSettingsUpdate"}}}}, "responses": {"200": {"description": "DLP settings.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DlpSettings"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to update DLP settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Settings"], "x-api-token-group": ["Zero Trust Write"]}
```
