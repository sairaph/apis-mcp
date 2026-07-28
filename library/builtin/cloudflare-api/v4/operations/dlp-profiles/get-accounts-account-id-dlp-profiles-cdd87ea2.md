---
title: List all profiles
page_id: operation-get-accounts-account-id-dlp-profiles-c2c32536
path: operations/dlp-profiles
description: Lists all DLP profiles in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/profiles
operation_ids:
    - dlp-profiles-list-all-profiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all profiles

`GET /accounts/{account_id}/dlp/profiles`

Operation ID: `dlp-profiles-list-all-profiles`

Lists all DLP profiles in an account.

## Definition

```yaml
{"operationId": "dlp-profiles-list-all-profiles", "summary": "List all profiles", "description": "Lists all DLP profiles in an account.", "parameters": [{"name": "all", "in": "query", "description": "Return all profiles, including those that current account does not have access to.", "schema": {"type": "boolean"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "List all profiles response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_ProfileArray"}}, "type": "object"}]}}}}, "4XX": {"description": "List all profiles failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
