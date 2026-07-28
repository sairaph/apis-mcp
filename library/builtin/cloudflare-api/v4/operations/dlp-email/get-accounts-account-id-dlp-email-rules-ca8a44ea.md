---
title: List all email scanner rules
page_id: operation-get-accounts-account-id-dlp-email-rules-c992e2a8
path: operations/dlp-email
description: Lists all email scanner rules for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/email/rules
operation_ids:
    - dlp-email-scanner-list-all-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all email scanner rules

`GET /accounts/{account_id}/dlp/email/rules`

Operation ID: `dlp-email-scanner-list-all-rules`

Lists all email scanner rules for an account.

## Definition

```yaml
{"operationId": "dlp-email-scanner-list-all-rules", "summary": "List all email scanner rules", "description": "Lists all email scanner rules for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "List all email scanner rules response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_EmailRuleArray"}}, "type": "object"}]}}}}, "4XX": {"description": "List all email scanner rules failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Email"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
