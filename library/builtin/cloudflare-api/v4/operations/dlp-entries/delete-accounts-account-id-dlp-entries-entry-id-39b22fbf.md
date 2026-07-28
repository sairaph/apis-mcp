---
title: Delete custom entry
page_id: operation-delete-accounts-account-id-dlp-entries-entry-id-2d4387af
path: operations/dlp-entries
description: Deletes a DLP custom entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/entries/{entry_id}
operation_ids:
    - dlp-entries-delete-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete custom entry

`DELETE /accounts/{account_id}/dlp/entries/{entry_id}`

Operation ID: `dlp-entries-delete-entry`

Deletes a DLP custom entry.

## Definition

```yaml
{"operationId": "dlp-entries-delete-entry", "summary": "Delete custom entry", "description": "Deletes a DLP custom entry.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Delete custom entry response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete custom entry failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Entries"], "x-api-token-group": ["Zero Trust Write"]}
```
