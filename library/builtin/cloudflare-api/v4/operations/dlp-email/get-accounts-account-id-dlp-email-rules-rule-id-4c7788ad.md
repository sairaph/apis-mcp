---
title: Get an email scanner rule
page_id: operation-get-accounts-account-id-dlp-email-rules-rule-id-3c34ec61
path: operations/dlp-email
description: Gets detailed configuration for a specific DLP email scanning rule, including detection patterns and actions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/email/rules/{rule_id}
operation_ids:
    - dlp-email-scanner-get-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an email scanner rule

`GET /accounts/{account_id}/dlp/email/rules/{rule_id}`

Operation ID: `dlp-email-scanner-get-rule`

Gets detailed configuration for a specific DLP email scanning rule, including detection patterns and actions.

## Definition

```yaml
{"operationId": "dlp-email-scanner-get-rule", "summary": "Get an email scanner rule", "description": "Gets detailed configuration for a specific DLP email scanning rule, including detection patterns and actions.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "rule_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Get Email Scanner Rule response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_EmailRule"}}, "type": "object"}]}}}}, "4XX": {"description": "Get Email Scanner Rule failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Email"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
