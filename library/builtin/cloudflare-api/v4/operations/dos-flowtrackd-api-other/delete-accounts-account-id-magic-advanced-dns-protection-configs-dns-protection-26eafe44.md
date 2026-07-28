---
title: Delete all DNS Protection rules.
page_id: operation-delete-accounts-account-id-magic-advanced-dns-protection-configs-dns-pro-2ef2ed13
path: operations/dos-flowtrackd-api-other
description: Delete all DNS Protection rules for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules
operation_ids:
    - deleteDnsProtectionRulesForAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete all DNS Protection rules.

`DELETE /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules`

Operation ID: `deleteDnsProtectionRulesForAccount`

Delete all DNS Protection rules for an account.

## Definition

```yaml
{"operationId": "deleteDnsProtectionRulesForAccount", "summary": "Delete all DNS Protection rules.", "description": "Delete all DNS Protection rules for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "responses": {"200": {"description": "Delete all DNS Protection rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
