---
title: Create mapping
page_id: operation-post-accounts-account-id-dlp-email-account-mapping-b76b979b
path: operations/dlp-email
description: Creates a mapping between a Cloudflare account and an email provider for DLP email scanning integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/email/account_mapping
operation_ids:
    - dlp-email-scanner-create-account-mapping
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create mapping

`POST /accounts/{account_id}/dlp/email/account_mapping`

Operation ID: `dlp-email-scanner-create-account-mapping`

Creates a mapping between a Cloudflare account and an email provider for DLP email scanning integration.

## Definition

```yaml
{"operationId": "dlp-email-scanner-create-account-mapping", "summary": "Create mapping", "description": "Creates a mapping between a Cloudflare account and an email provider for DLP email scanning integration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Account mapping.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_UpdateAddinAccountMapping"}}}}, "responses": {"200": {"description": "New Email Scanner Account Mapping response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_AddinAccountMapping"}}, "type": "object"}]}}}}, "4XX": {"description": "New Email Scanner Account Mapping failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Email"], "x-api-token-group": ["Zero Trust Write"]}
```
