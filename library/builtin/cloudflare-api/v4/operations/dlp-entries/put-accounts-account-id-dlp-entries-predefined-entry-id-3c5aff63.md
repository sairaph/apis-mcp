---
title: Update predefined entry
page_id: operation-put-accounts-account-id-dlp-entries-predefined-entry-id-191f50a0
path: operations/dlp-entries
description: Updates a DLP entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/entries/predefined/{entry_id}
operation_ids:
    - dlp-entries-update-predefined-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update predefined entry

`PUT /accounts/{account_id}/dlp/entries/predefined/{entry_id}`

Operation ID: `dlp-entries-update-predefined-entry`

Updates a DLP entry.

## Definition

```yaml
{"operationId": "dlp-entries-update-predefined-entry", "summary": "Update predefined entry", "description": "Updates a DLP entry.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Settings to enable or disable predefined entry in owning profile.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_PredefinedEntryUpdate"}}}}, "responses": {"200": {"description": "Update predefined entry response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_PredefinedEntry"}}, "type": "object"}]}}}}, "4XX": {"description": "Update entry failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Entries"], "x-api-token-group": ["Zero Trust Write"]}
```
