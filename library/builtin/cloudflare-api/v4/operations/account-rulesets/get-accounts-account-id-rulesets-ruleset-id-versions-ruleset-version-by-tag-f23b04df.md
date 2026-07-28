---
title: List an account ruleset version's rules by tag
page_id: operation-get-accounts-account-id-rulesets-ruleset-id-versions-ruleset-version-by-0f3b9bfd
path: operations/account-rulesets
description: Fetches the rules of a managed account ruleset version for a given tag.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rulesets/{ruleset_id}/versions/{ruleset_version}/by_tag/{rule_tag}
operation_ids:
    - listAccountRulesetVersionRulesByTag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List an account ruleset version's rules by tag

`GET /accounts/{account_id}/rulesets/{ruleset_id}/versions/{ruleset_version}/by_tag/{rule_tag}`

Operation ID: `listAccountRulesetVersionRulesByTag`

Fetches the rules of a managed account ruleset version for a given tag.

## Definition

```yaml
{"operationId": "listAccountRulesetVersionRulesByTag", "summary": "List an account ruleset version's rules by tag", "description": "Fetches the rules of a managed account ruleset version for a given tag.", "parameters": [{"name": "rule_tag", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RuleCategory"}}, {"name": "ruleset_version", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetVersion"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_AccountId"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account Rulesets"], "x-api-token-group": ["Mass URL Redirects Write", "Mass URL Redirects Read", "Magic Firewall Write", "Magic Firewall Read", "L4 DDoS Managed Ruleset Write", "L4 DDoS Managed Ruleset Read", "Transform Rules Write", "Transform Rules Read", "Select Configuration Write", "Select Configuration Read", "Account WAF Write", "Account WAF Read", "Account Rulesets Read", "Account Rulesets Write", "Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
