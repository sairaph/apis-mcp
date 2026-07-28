---
title: Update an account ruleset rule
page_id: operation-patch-accounts-account-id-rulesets-ruleset-id-rules-rule-id-e5acb960
path: operations/account-rulesets
description: Updates an existing rule in an account ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/rulesets/{ruleset_id}/rules/{rule_id}
operation_ids:
    - updateAccountRulesetRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an account ruleset rule

`PATCH /accounts/{account_id}/rulesets/{ruleset_id}/rules/{rule_id}`

Operation ID: `updateAccountRulesetRule`

Updates an existing rule in an account ruleset.

## Definition

```yaml
{"operationId": "updateAccountRulesetRule", "summary": "Update an account ruleset rule", "description": "Updates an existing rule in an account ruleset.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RuleId"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "requestBody": {"$ref": "#/components/requestBodies/rulesets_Rule"}, "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "Transform Rules Write", "Select Configuration Write", "Account WAF Write", "Account Rulesets Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rulesets.account-rulesets.rules", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
