---
title: Create an account ruleset
page_id: operation-post-accounts-account-id-rulesets-572446d2
path: operations/account-rulesets
description: Creates a ruleset at the account level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/rulesets
operation_ids:
    - createAccountRuleset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an account ruleset

`POST /accounts/{account_id}/rulesets`

Operation ID: `createAccountRuleset`

Creates a ruleset at the account level.

## Definition

```yaml
{"operationId": "createAccountRuleset", "summary": "Create an account ruleset", "description": "Creates a ruleset at the account level.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "requestBody": {"$ref": "#/components/requestBodies/rulesets_CreateRuleset"}, "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "Transform Rules Write", "Select Configuration Write", "Account WAF Write", "Account Rulesets Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
