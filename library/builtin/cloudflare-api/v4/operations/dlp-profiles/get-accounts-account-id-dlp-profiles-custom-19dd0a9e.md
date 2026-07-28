---
title: List all custom profiles
page_id: operation-get-accounts-account-id-dlp-profiles-custom-06516d44
path: operations/dlp-profiles
description: Lists all DLP custom profiles in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/profiles/custom
operation_ids:
    - dlp-profiles-list-all-custom-profiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all custom profiles

`GET /accounts/{account_id}/dlp/profiles/custom`

Operation ID: `dlp-profiles-list-all-custom-profiles`

Lists all DLP custom profiles in an account.

## Definition

```yaml
{"operationId": "dlp-profiles-list-all-custom-profiles", "summary": "List all custom profiles", "description": "Lists all DLP custom profiles in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "List all custom profiles response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_CustomProfileArray"}}, "type": "object"}]}}}}, "4XX": {"description": "List all profiles failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Profiles"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
