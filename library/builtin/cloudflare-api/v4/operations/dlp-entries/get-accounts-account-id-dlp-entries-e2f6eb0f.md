---
title: List all entries
page_id: operation-get-accounts-account-id-dlp-entries-77d04534
path: operations/dlp-entries
description: Lists all DLP entries in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/entries
operation_ids:
    - dlp-entries-list-all-entries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all entries

`GET /accounts/{account_id}/dlp/entries`

Operation ID: `dlp-entries-list-all-entries`

Lists all DLP entries in an account.

## Definition

```yaml
{"operationId": "dlp-entries-list-all-entries", "summary": "List all entries", "description": "Lists all DLP entries in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "List all entries response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_EntryWithUploadStatus"}}}, "type": "object"}]}}}}, "4XX": {"description": "List all entries failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Entries"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
