---
title: Update an account entry point ruleset
page_id: operation-put-accounts-account-id-rulesets-phases-ruleset-phase-entrypoint-0561c36f
path: operations/account-rulesets
description: Updates an account entry point ruleset, creating a new version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/rulesets/phases/{ruleset_phase}/entrypoint
operation_ids:
    - updateAccountEntrypointRuleset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an account entry point ruleset

`PUT /accounts/{account_id}/rulesets/phases/{ruleset_phase}/entrypoint`

Operation ID: `updateAccountEntrypointRuleset`

Updates an account entry point ruleset, creating a new version.

## Definition

```yaml
{"operationId": "updateAccountEntrypointRuleset", "summary": "Update an account entry point ruleset", "description": "Updates an account entry point ruleset, creating a new version.", "parameters": [{"name": "ruleset_phase", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetPhase"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "requestBody": {"$ref": "#/components/requestBodies/rulesets_UpdateEntrypointRuleset"}, "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "Transform Rules Write", "Select Configuration Write", "Account WAF Write", "Account Rulesets Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
