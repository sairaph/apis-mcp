---
title: List all DNS Protection rules.
page_id: operation-get-accounts-account-id-magic-advanced-dns-protection-configs-dns-protec-d7a393f7
path: operations/dos-flowtrackd-api-other
description: List all DNS Protection rules for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules
operation_ids:
    - listDnsProtectionRulesForAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all DNS Protection rules.

`GET /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules`

Operation ID: `listDnsProtectionRulesForAccount`

List all DNS Protection rules for an account.

## Definition

```yaml
{"operationId": "listDnsProtectionRulesForAccount", "summary": "List all DNS Protection rules.", "description": "List all DNS Protection rules for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "page", "in": "query", "description": "The page number for pagination. Defaults to 1.", "schema": {"type": "integer", "format": "int64", "x-auditable": true}, "explode": false}, {"name": "per_page", "in": "query", "description": "The number of items per page. Must be between 10 and 1000. Defaults to 25.", "schema": {"type": "integer", "format": "int64", "x-auditable": true}, "explode": false}, {"name": "order", "in": "query", "description": "The field to order by. Defaults to 'prefix'.", "schema": {"type": "string", "x-auditable": true}, "explode": false}, {"name": "direction", "in": "query", "description": "The direction of ordering (ASC or DESC). Defaults to 'ASC'.", "schema": {"type": "string", "x-auditable": true}, "explode": false}], "responses": {"200": {"description": "List all DNS Protection rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_dns-protection-rule-list-response"}}}}, "4XX": {"description": "List all DNS Protection rules failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write", "DDoS Protection Read"]}
```
