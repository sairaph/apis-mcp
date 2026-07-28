---
title: Get an account entry point ruleset version
page_id: operation-get-accounts-account-id-rulesets-phases-ruleset-phase-entrypoint-version-1b4b7c93
path: operations/account-rulesets
description: Fetches a specific version of an account entry point ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rulesets/phases/{ruleset_phase}/entrypoint/versions/{ruleset_version}
operation_ids:
    - getAccountEntrypointRulesetVersion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an account entry point ruleset version

`GET /accounts/{account_id}/rulesets/phases/{ruleset_phase}/entrypoint/versions/{ruleset_version}`

Operation ID: `getAccountEntrypointRulesetVersion`

Fetches a specific version of an account entry point ruleset.

## Definition

```yaml
{"operationId": "getAccountEntrypointRulesetVersion", "summary": "Get an account entry point ruleset version", "description": "Fetches a specific version of an account entry point ruleset.", "parameters": [{"name": "ruleset_version", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetVersion"}}, {"name": "ruleset_phase", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetPhase"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Mass URL Redirects Read", "Magic Firewall Write", "Magic Firewall Read", "L4 DDoS Managed Ruleset Write", "L4 DDoS Managed Ruleset Read", "Transform Rules Write", "Transform Rules Read", "Select Configuration Write", "Select Configuration Read", "Account WAF Write", "Account WAF Read", "Account Rulesets Read", "Account Rulesets Write", "Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
