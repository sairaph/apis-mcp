---
title: Create predefined entry
page_id: operation-post-accounts-account-id-dlp-entries-predefined-793874c0
path: operations/dlp-predefined-entries
description: |-
    Predefined entries can't be created, this will update an existing predefined entry.
    This is needed for our generated terraform API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/entries/predefined
operation_ids:
    - dlp-entries-create-predefined-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create predefined entry

`POST /accounts/{account_id}/dlp/entries/predefined`

Operation ID: `dlp-entries-create-predefined-entry`

Predefined entries can't be created, this will update an existing predefined entry.
This is needed for our generated terraform API.

## Definition

```yaml
{"operationId": "dlp-entries-create-predefined-entry", "summary": "Create predefined entry", "description": "Predefined entries can't be created, this will update an existing predefined entry.\nThis is needed for our generated terraform API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "This endpoint will update an existing predefined entry. It is not possible to create new predefined entries.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewPredefinedEntry"}}}}, "responses": {"200": {"description": "Create predefined entry response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_PredefinedEntry"}}, "type": "object"}]}}}}, "4XX": {"description": "Create entry failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Predefined Entries"], "x-api-token-group": ["Zero Trust Write"]}
```
