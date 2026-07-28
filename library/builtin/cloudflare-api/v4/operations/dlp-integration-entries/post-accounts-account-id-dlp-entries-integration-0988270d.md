---
title: Create integration entry
page_id: operation-post-accounts-account-id-dlp-entries-integration-e7e60bdb
path: operations/dlp-integration-entries
description: |-
    Integration entries can't be created, this will update an existing integration entry.
    This is needed for our generated terraform API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/entries/integration
operation_ids:
    - dlp-entries-create-integration-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create integration entry

`POST /accounts/{account_id}/dlp/entries/integration`

Operation ID: `dlp-entries-create-integration-entry`

Integration entries can't be created, this will update an existing integration entry.
This is needed for our generated terraform API.

## Definition

```yaml
{"operationId": "dlp-entries-create-integration-entry", "summary": "Create integration entry", "description": "Integration entries can't be created, this will update an existing integration entry.\nThis is needed for our generated terraform API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "This endpoint will update an existing integration entry. It is not possible to create new integration entries.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewPredefinedEntry"}}}}, "responses": {"200": {"description": "Create integration entry response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_IntegrationEntry"}}, "type": "object"}]}}}}, "4XX": {"description": "Create entry failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Integration Entries"], "x-api-token-group": ["Zero Trust Write"]}
```
