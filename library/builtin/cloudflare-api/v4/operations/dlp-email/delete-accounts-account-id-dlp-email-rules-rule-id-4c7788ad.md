---
title: Delete email scanner rule
page_id: operation-delete-accounts-account-id-dlp-email-rules-rule-id-2157a0bf
path: operations/dlp-email
description: Removes a DLP email scanning rule. The rule will no longer be applied to email messages.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/email/rules/{rule_id}
operation_ids:
    - dlp-email-scanner-delete-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete email scanner rule

`DELETE /accounts/{account_id}/dlp/email/rules/{rule_id}`

Operation ID: `dlp-email-scanner-delete-rule`

Removes a DLP email scanning rule. The rule will no longer be applied to email messages.

## Definition

```yaml
{"operationId": "dlp-email-scanner-delete-rule", "summary": "Delete email scanner rule", "description": "Removes a DLP email scanning rule. The rule will no longer be applied to email messages.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "rule_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Delete Email Scanner Rule response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_EmailRule"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Email Scanner Rule failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Email"], "x-api-token-group": ["Zero Trust Write"]}
```
