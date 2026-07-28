---
title: Update custom entry
page_id: operation-put-accounts-account-id-dlp-entries-custom-entry-id-83c7f351
path: operations/dlp-entries
description: Updates a DLP custom entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/entries/custom/{entry_id}
operation_ids:
    - dlp-entries-update-custom-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update custom entry

`PUT /accounts/{account_id}/dlp/entries/custom/{entry_id}`

Operation ID: `dlp-entries-update-custom-entry`

Updates a DLP custom entry.

## Definition

```yaml
{"operationId": "dlp-entries-update-custom-entry", "summary": "Update custom entry", "description": "Updates a DLP custom entry.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Update to be applied to the entry.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_CustomEntryUpdate"}}}}, "responses": {"200": {"description": "Update entry response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_CustomEntry"}}, "type": "object"}]}}}}, "4XX": {"description": "Update entry failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Entries"], "x-api-token-group": ["Zero Trust Write"]}
```
