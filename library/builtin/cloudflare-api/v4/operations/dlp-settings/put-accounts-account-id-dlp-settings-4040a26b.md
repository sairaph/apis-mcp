---
title: Update DLP account-level settings (full replacement).
page_id: operation-put-accounts-account-id-dlp-settings-53df9c89
path: operations/dlp-settings
description: Missing fields are reset to initial (unconfigured) values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/settings
operation_ids:
    - dlp-settings-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update DLP account-level settings (full replacement).

`PUT /accounts/{account_id}/dlp/settings`

Operation ID: `dlp-settings-update`

Missing fields are reset to initial (unconfigured) values.

## Definition

```yaml
{"operationId": "dlp-settings-update", "summary": "Update DLP account-level settings (full replacement).", "description": "Missing fields are reset to initial (unconfigured) values.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "New DLP settings.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_DlpSettingsUpdate"}}}}, "responses": {"200": {"description": "DLP settings.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DlpSettings"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to update DLP settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Settings"], "x-api-token-group": ["Zero Trust Write"]}
```
