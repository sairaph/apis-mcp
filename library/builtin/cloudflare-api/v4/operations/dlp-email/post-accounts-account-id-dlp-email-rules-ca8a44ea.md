---
title: Create email scanner rule
page_id: operation-post-accounts-account-id-dlp-email-rules-c36d9813
path: operations/dlp-email
description: Creates a new DLP email scanning rule that defines what content patterns to detect in email messages and what actions to take.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/email/rules
operation_ids:
    - dlp-email-scanner-create-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create email scanner rule

`POST /accounts/{account_id}/dlp/email/rules`

Operation ID: `dlp-email-scanner-create-rule`

Creates a new DLP email scanning rule that defines what content patterns to detect in email messages and what actions to take.

## Definition

```yaml
{"operationId": "dlp-email-scanner-create-rule", "summary": "Create email scanner rule", "description": "Creates a new DLP email scanning rule that defines what content patterns to detect in email messages and what actions to take.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Rule description.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_CreateEmailRule"}}}}, "responses": {"200": {"description": "New Email Scanner Rule response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_EmailRule"}}, "type": "object"}]}}}}, "4XX": {"description": "New Email Scanner Rule failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Email"], "x-api-token-group": ["Zero Trust Write"]}
```
