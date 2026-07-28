---
title: Get a zone entry point ruleset
page_id: operation-get-zones-zone-id-rulesets-phases-ruleset-phase-entrypoint-8c39d0fd
path: operations/zone-rulesets
description: Fetches the latest version of the zone entry point ruleset for a given phase.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/rulesets/phases/{ruleset_phase}/entrypoint
operation_ids:
    - getZoneEntrypointRuleset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a zone entry point ruleset

`GET /zones/{zone_id}/rulesets/phases/{ruleset_phase}/entrypoint`

Operation ID: `getZoneEntrypointRuleset`

Fetches the latest version of the zone entry point ruleset for a given phase.

## Definition

```yaml
{"operationId": "getZoneEntrypointRuleset", "summary": "Get a zone entry point ruleset", "description": "Fetches the latest version of the zone entry point ruleset for a given phase.", "parameters": [{"name": "ruleset_phase", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetPhase"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Rulesets"], "x-api-token-group": ["Response Compression Write", "Response Compression Read", "Config Settings Write", "Config Settings Read", "Dynamic URL Redirects Write", "Dynamic URL Redirects Read", "Cache Settings Write", "Cache Settings Read", "Custom Errors Write", "Custom Errors Read", "Origin Write", "Origin Read", "Managed headers Write", "Managed headers Read", "Zone Transform Rules Write", "Zone Transform Rules Read", "Mass URL Redirects Write", "Mass URL Redirects Read", "Magic Firewall Write", "Magic Firewall Read", "L4 DDoS Managed Ruleset Write", "L4 DDoS Managed Ruleset Read", "HTTP DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Read", "Sanitize Write", "Sanitize Read", "Transform Rules Write", "Transform Rules Read", "Select Configuration Write", "Select Configuration Read", "Bot Management Write", "Bot Management Read", "Zone WAF Write", "Zone WAF Read", "Account WAF Write", "Account WAF Read", "Account Rulesets Read", "Account Rulesets Write", "Logs Write", "Logs Read", "Logs Write", "Logs Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
