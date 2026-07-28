---
title: List a zone ruleset's versions
page_id: operation-get-zones-zone-id-rulesets-ruleset-id-versions-32cb8767
path: operations/zone-rulesets
description: Fetches the versions of a zone ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/rulesets/{ruleset_id}/versions
operation_ids:
    - listZoneRulesetVersions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List a zone ruleset's versions

`GET /zones/{zone_id}/rulesets/{ruleset_id}/versions`

Operation ID: `listZoneRulesetVersions`

Fetches the versions of a zone ruleset.

## Definition

```yaml
{"operationId": "listZoneRulesetVersions", "summary": "List a zone ruleset's versions", "description": "Fetches the versions of a zone ruleset.", "parameters": [{"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Rulesets"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Rulesets"], "x-api-token-group": ["Response Compression Write", "Response Compression Read", "Config Settings Write", "Config Settings Read", "Dynamic URL Redirects Write", "Dynamic URL Redirects Read", "Cache Settings Write", "Cache Settings Read", "Custom Errors Write", "Custom Errors Read", "Origin Write", "Origin Read", "Managed headers Write", "Managed headers Read", "Zone Transform Rules Write", "Zone Transform Rules Read", "Mass URL Redirects Write", "Mass URL Redirects Read", "Magic Firewall Write", "Magic Firewall Read", "L4 DDoS Managed Ruleset Write", "L4 DDoS Managed Ruleset Read", "HTTP DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Read", "Sanitize Write", "Sanitize Read", "Transform Rules Write", "Transform Rules Read", "Select Configuration Write", "Select Configuration Read", "Bot Management Write", "Bot Management Read", "Zone WAF Write", "Zone WAF Read", "Account WAF Write", "Account WAF Read", "Account Rulesets Read", "Account Rulesets Write", "Logs Write", "Logs Read", "Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
