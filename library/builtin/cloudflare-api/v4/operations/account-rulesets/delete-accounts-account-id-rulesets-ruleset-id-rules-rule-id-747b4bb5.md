---
title: Delete an account ruleset rule
page_id: operation-delete-accounts-account-id-rulesets-ruleset-id-rules-rule-id-176e51bc
path: operations/account-rulesets
description: Deletes an existing rule from an account ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/rulesets/{ruleset_id}/rules/{rule_id}
operation_ids:
    - deleteAccountRulesetRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an account ruleset rule

`DELETE /accounts/{account_id}/rulesets/{ruleset_id}/rules/{rule_id}`

Operation ID: `deleteAccountRulesetRule`

Deletes an existing rule from an account ruleset.

## Definition

```yaml
{"operationId": "deleteAccountRulesetRule", "summary": "Delete an account ruleset rule", "description": "Deletes an existing rule from an account ruleset.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RuleId"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "Transform Rules Write", "Select Configuration Write", "Account WAF Write", "Account Rulesets Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rulesets.account-rulesets.rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
