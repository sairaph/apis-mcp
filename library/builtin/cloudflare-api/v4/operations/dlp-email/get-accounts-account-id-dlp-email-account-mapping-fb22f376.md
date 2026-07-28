---
title: Get mapping
page_id: operation-get-accounts-account-id-dlp-email-account-mapping-ad722288
path: operations/dlp-email
description: Retrieves the email provider mapping configuration for DLP email scanning.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/email/account_mapping
operation_ids:
    - dlp-email-scanner-get-account-mapping
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get mapping

`GET /accounts/{account_id}/dlp/email/account_mapping`

Operation ID: `dlp-email-scanner-get-account-mapping`

Retrieves the email provider mapping configuration for DLP email scanning.

## Definition

```yaml
{"operationId": "dlp-email-scanner-get-account-mapping", "summary": "Get mapping", "description": "Retrieves the email provider mapping configuration for DLP email scanning.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Get Email Scanner Account Mapping response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_AddinAccountMapping"}}, "type": "object"}]}}}}, "4XX": {"description": "Get Email Scanner Account Mapping failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Email"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
