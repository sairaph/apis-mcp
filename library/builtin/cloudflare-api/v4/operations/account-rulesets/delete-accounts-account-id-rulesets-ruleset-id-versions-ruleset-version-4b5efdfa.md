---
title: Delete an account ruleset version
page_id: operation-delete-accounts-account-id-rulesets-ruleset-id-versions-ruleset-version-b91b7c3d
path: operations/account-rulesets
description: Deletes an existing version of an account ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/rulesets/{ruleset_id}/versions/{ruleset_version}
operation_ids:
    - deleteAccountRulesetVersion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an account ruleset version

`DELETE /accounts/{account_id}/rulesets/{ruleset_id}/versions/{ruleset_version}`

Operation ID: `deleteAccountRulesetVersion`

Deletes an existing version of an account ruleset.

## Definition

```yaml
{"operationId": "deleteAccountRulesetVersion", "summary": "Delete an account ruleset version", "description": "Deletes an existing version of an account ruleset.", "parameters": [{"name": "ruleset_version", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetVersion"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "responses": {"204": {"$ref": "#/components/responses/rulesets_Empty"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "Transform Rules Write", "Select Configuration Write", "Account WAF Write", "Account Rulesets Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
