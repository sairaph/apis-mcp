---
title: Delete integration entry
page_id: operation-delete-accounts-account-id-dlp-entries-integration-entry-id-81e95bde
path: operations/dlp-integration-entries
description: This is a no-op as integration entires can't be deleted but is needed for our generated terraform API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/entries/integration/{entry_id}
operation_ids:
    - dlp-entries-delete-integration-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete integration entry

`DELETE /accounts/{account_id}/dlp/entries/integration/{entry_id}`

Operation ID: `dlp-entries-delete-integration-entry`

This is a no-op as integration entires can't be deleted but is needed for our generated terraform API.

## Definition

```yaml
{"operationId": "dlp-entries-delete-integration-entry", "summary": "Delete integration entry", "description": "This is a no-op as integration entires can't be deleted but is needed for our generated terraform API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Delete integration entry response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete entry failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Integration Entries"], "x-api-token-group": ["Zero Trust Write"]}
```
