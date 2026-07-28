---
title: Get DLP Entry
page_id: operation-get-accounts-account-id-dlp-entries-entry-id-22c35279
path: operations/dlp-entries
description: Fetches a DLP entry by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/entries/{entry_id}
operation_ids:
    - dlp-entries-get-dlp-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get DLP Entry

`GET /accounts/{account_id}/dlp/entries/{entry_id}`

Operation ID: `dlp-entries-get-dlp-entry`

Fetches a DLP entry by ID.

## Definition

```yaml
{"operationId": "dlp-entries-get-dlp-entry", "summary": "Get DLP Entry", "description": "Fetches a DLP entry by ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "entry_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Get entry response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_EntryWithSharedProfiles"}}, "type": "object"}]}}}}, "4XX": {"description": "Get entry failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Entries"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
