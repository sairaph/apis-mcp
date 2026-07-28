---
title: Update an account ruleset
page_id: operation-put-accounts-account-id-rulesets-ruleset-id-2ef2c28b
path: operations/account-rulesets
description: Updates an account ruleset, creating a new version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/rulesets/{ruleset_id}
operation_ids:
    - updateAccountRuleset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an account ruleset

`PUT /accounts/{account_id}/rulesets/{ruleset_id}`

Operation ID: `updateAccountRuleset`

Updates an account ruleset, creating a new version.

## Definition

```yaml
{"operationId": "updateAccountRuleset", "summary": "Update an account ruleset", "description": "Updates an account ruleset, creating a new version.", "parameters": [{"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "requestBody": {"$ref": "#/components/requestBodies/rulesets_UpdateRuleset"}, "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "Transform Rules Write", "Select Configuration Write", "Account WAF Write", "Account Rulesets Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
