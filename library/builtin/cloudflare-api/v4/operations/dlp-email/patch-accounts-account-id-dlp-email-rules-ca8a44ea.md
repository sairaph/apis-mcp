---
title: Update email scanner rule priorities
page_id: operation-patch-accounts-account-id-dlp-email-rules-aae05ef4
path: operations/dlp-email
description: Reorders DLP email scanning rules by updating their priority values. Higher priority rules are evaluated first.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/dlp/email/rules
operation_ids:
    - dlp-email-scanner-update-rule-priorities
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update email scanner rule priorities

`PATCH /accounts/{account_id}/dlp/email/rules`

Operation ID: `dlp-email-scanner-update-rule-priorities`

Reorders DLP email scanning rules by updating their priority values. Higher priority rules are evaluated first.

## Definition

```yaml
{"operationId": "dlp-email-scanner-update-rule-priorities", "summary": "Update email scanner rule priorities", "description": "Reorders DLP email scanning rules by updating their priority values. Higher priority rules are evaluated first.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Rule priorities.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_UpdateEmailRulePriorities"}}}}, "responses": {"200": {"description": "Update Email Scanner Rule priorities response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_EmailRule"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Email Scanner Rule priorities failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Email"], "x-api-token-group": ["Zero Trust Write"]}
```
