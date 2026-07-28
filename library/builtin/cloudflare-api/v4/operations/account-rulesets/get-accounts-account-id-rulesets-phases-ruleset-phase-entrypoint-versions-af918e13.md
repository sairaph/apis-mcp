---
title: List an account entry point ruleset's versions
page_id: operation-get-accounts-account-id-rulesets-phases-ruleset-phase-entrypoint-version-98d2f604
path: operations/account-rulesets
description: Fetches the versions of an account entry point ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rulesets/phases/{ruleset_phase}/entrypoint/versions
operation_ids:
    - listAccountEntrypointRulesetVersions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List an account entry point ruleset's versions

`GET /accounts/{account_id}/rulesets/phases/{ruleset_phase}/entrypoint/versions`

Operation ID: `listAccountEntrypointRulesetVersions`

Fetches the versions of an account entry point ruleset.

## Definition

```yaml
{"operationId": "listAccountEntrypointRulesetVersions", "summary": "List an account entry point ruleset's versions", "description": "Fetches the versions of an account entry point ruleset.", "parameters": [{"name": "ruleset_phase", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetPhase"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Rulesets"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Mass URL Redirects Read", "Magic Firewall Write", "Magic Firewall Read", "L4 DDoS Managed Ruleset Write", "L4 DDoS Managed Ruleset Read", "Transform Rules Write", "Transform Rules Read", "Select Configuration Write", "Select Configuration Read", "Account WAF Write", "Account WAF Read", "Account Rulesets Read", "Account Rulesets Write", "Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rulesets.account-rulesets.phases.versions", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
