---
title: List account rulesets
page_id: operation-get-accounts-account-id-rulesets-6aa08618
path: operations/account-rulesets
description: Fetches all rulesets at the account level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rulesets
operation_ids:
    - listAccountRulesets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account rulesets

`GET /accounts/{account_id}/rulesets`

Operation ID: `listAccountRulesets`

Fetches all rulesets at the account level.

## Definition

```yaml
{"operationId": "listAccountRulesets", "summary": "List account rulesets", "description": "Fetches all rulesets at the account level.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}, {"name": "cursor", "in": "query", "schema": {"$ref": "#/components/schemas/rulesets_Cursor"}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/rulesets_PerPage"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Rulesets"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Mass URL Redirects Read", "Magic Firewall Write", "Magic Firewall Read", "L4 DDoS Managed Ruleset Write", "L4 DDoS Managed Ruleset Read", "Transform Rules Write", "Transform Rules Read", "Select Configuration Write", "Select Configuration Read", "Account WAF Write", "Account WAF Read", "Account Rulesets Read", "Account Rulesets Write", "Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
