---
title: Create an account ruleset rule
page_id: operation-post-accounts-account-id-rulesets-ruleset-id-rules-42dd65dd
path: operations/account-rulesets
description: Adds a new rule to an account ruleset. The rule will be added to the end of the existing list of rules in the ruleset by default.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/rulesets/{ruleset_id}/rules
operation_ids:
    - createAccountRulesetRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an account ruleset rule

`POST /accounts/{account_id}/rulesets/{ruleset_id}/rules`

Operation ID: `createAccountRulesetRule`

Adds a new rule to an account ruleset. The rule will be added to the end of the existing list of rules in the ruleset by default.

## Definition

```yaml
{"operationId": "createAccountRulesetRule", "summary": "Create an account ruleset rule", "description": "Adds a new rule to an account ruleset. The rule will be added to the end of the existing list of rules in the ruleset by default.", "parameters": [{"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "requestBody": {"$ref": "#/components/requestBodies/rulesets_Rule"}, "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "Transform Rules Write", "Select Configuration Write", "Account WAF Write", "Account Rulesets Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rulesets.account-rulesets.rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
