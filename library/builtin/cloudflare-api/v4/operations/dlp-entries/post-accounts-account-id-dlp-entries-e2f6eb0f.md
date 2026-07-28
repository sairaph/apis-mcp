---
title: Create custom entry
page_id: operation-post-accounts-account-id-dlp-entries-0e43b7aa
path: operations/dlp-entries
description: Creates a DLP custom entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/entries
operation_ids:
    - dlp-entries-create-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create custom entry

`POST /accounts/{account_id}/dlp/entries`

Operation ID: `dlp-entries-create-entry`

Creates a DLP custom entry.

## Definition

```yaml
{"operationId": "dlp-entries-create-entry", "summary": "Create custom entry", "description": "Creates a DLP custom entry.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "A new entry to create.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewEntry"}}}}, "responses": {"200": {"description": "Create new custom entry response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_CustomEntry"}}, "type": "object"}]}}}}, "4XX": {"description": "Create new custom entry failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Entries"], "x-api-token-group": ["Zero Trust Write"]}
```
