---
title: Delete an account ruleset
page_id: operation-delete-accounts-account-id-rulesets-ruleset-id-a595d20c
path: operations/account-rulesets
description: Deletes all versions of an existing account ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/rulesets/{ruleset_id}
operation_ids:
    - deleteAccountRuleset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an account ruleset

`DELETE /accounts/{account_id}/rulesets/{ruleset_id}`

Operation ID: `deleteAccountRuleset`

Deletes all versions of an existing account ruleset.

## Definition

```yaml
{"operationId": "deleteAccountRuleset", "summary": "Delete an account ruleset", "description": "Deletes all versions of an existing account ruleset.", "parameters": [{"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "responses": {"204": {"$ref": "#/components/responses/rulesets_Empty"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "Transform Rules Write", "Select Configuration Write", "Account WAF Write", "Account Rulesets Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rulesets.account-rulesets", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
