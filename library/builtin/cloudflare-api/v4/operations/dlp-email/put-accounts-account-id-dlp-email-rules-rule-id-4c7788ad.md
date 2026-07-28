---
title: Update email scanner rule
page_id: operation-put-accounts-account-id-dlp-email-rules-rule-id-f9be04ca
path: operations/dlp-email
description: Updates a DLP email scanning rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/email/rules/{rule_id}
operation_ids:
    - dlp-email-scanner-update-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update email scanner rule

`PUT /accounts/{account_id}/dlp/email/rules/{rule_id}`

Operation ID: `dlp-email-scanner-update-rule`

Updates a DLP email scanning rule.

## Definition

```yaml
{"operationId": "dlp-email-scanner-update-rule", "summary": "Update email scanner rule", "description": "Updates a DLP email scanning rule.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "rule_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Rule description.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_CreateEmailRule"}}}}, "responses": {"200": {"description": "Update Email Scanner Rule response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_EmailRule"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Email Scanner Rule failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Email"], "x-api-token-group": ["Zero Trust Write"]}
```
