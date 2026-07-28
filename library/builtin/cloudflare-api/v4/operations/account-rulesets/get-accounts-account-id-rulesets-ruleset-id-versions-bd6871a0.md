---
title: List an account ruleset's versions
page_id: operation-get-accounts-account-id-rulesets-ruleset-id-versions-a355d2c6
path: operations/account-rulesets
description: Fetches the versions of an account ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rulesets/{ruleset_id}/versions
operation_ids:
    - listAccountRulesetVersions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List an account ruleset's versions

`GET /accounts/{account_id}/rulesets/{ruleset_id}/versions`

Operation ID: `listAccountRulesetVersions`

Fetches the versions of an account ruleset.

## Definition

```yaml
{"operationId": "listAccountRulesetVersions", "summary": "List an account ruleset's versions", "description": "Fetches the versions of an account ruleset.", "parameters": [{"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Rulesets"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Mass URL Redirects Read", "Magic Firewall Write", "Magic Firewall Read", "L4 DDoS Managed Ruleset Write", "L4 DDoS Managed Ruleset Read", "Transform Rules Write", "Transform Rules Read", "Select Configuration Write", "Select Configuration Read", "Account WAF Write", "Account WAF Read", "Account Rulesets Read", "Account Rulesets Write", "Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
