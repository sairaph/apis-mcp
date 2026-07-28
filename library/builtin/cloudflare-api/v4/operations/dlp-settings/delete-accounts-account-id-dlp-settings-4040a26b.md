---
title: Delete (reset) DLP account-level settings to initial values.
page_id: operation-delete-accounts-account-id-dlp-settings-6e2d978a
path: operations/dlp-settings
description: Deletes account-level DLP settings and returns the initial values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/settings
operation_ids:
    - dlp-settings-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete (reset) DLP account-level settings to initial values.

`DELETE /accounts/{account_id}/dlp/settings`

Operation ID: `dlp-settings-delete`

Deletes account-level DLP settings and returns the initial values.

## Definition

```yaml
{"operationId": "dlp-settings-delete", "summary": "Delete (reset) DLP account-level settings to initial values.", "description": "Deletes account-level DLP settings and returns the initial values.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "DLP settings reset.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DlpSettings"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to delete DLP settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Settings"], "x-api-token-group": ["Zero Trust Write"]}
```
