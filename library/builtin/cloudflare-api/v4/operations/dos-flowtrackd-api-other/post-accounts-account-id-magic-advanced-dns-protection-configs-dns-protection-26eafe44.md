---
title: Create DNS Protection rule.
page_id: operation-post-accounts-account-id-magic-advanced-dns-protection-configs-dns-prote-af36a958
path: operations/dos-flowtrackd-api-other
description: Create a DNS Protection rule for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules
operation_ids:
    - createDnsProtectionRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create DNS Protection rule.

`POST /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules`

Operation ID: `createDnsProtectionRule`

Create a DNS Protection rule for an account.

## Definition

```yaml
{"operationId": "createDnsProtectionRule", "summary": "Create DNS Protection rule.", "description": "Create a DNS Protection rule for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "requestBody": {"description": "The new DNS Protection rule to add.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_NewDnsProtectionRule"}}}}, "responses": {"200": {"description": "Create DNS Protection rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_dns-protection-rule-response"}}}}, "4XX": {"description": "Create DNS Protection rule failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
