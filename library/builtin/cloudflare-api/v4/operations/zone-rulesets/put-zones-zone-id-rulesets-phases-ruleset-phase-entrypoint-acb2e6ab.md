---
title: Update a zone entry point ruleset
page_id: operation-put-zones-zone-id-rulesets-phases-ruleset-phase-entrypoint-9c0d4d89
path: operations/zone-rulesets
description: Updates a zone entry point ruleset, creating a new version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/rulesets/phases/{ruleset_phase}/entrypoint
operation_ids:
    - updateZoneEntrypointRuleset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a zone entry point ruleset

`PUT /zones/{zone_id}/rulesets/phases/{ruleset_phase}/entrypoint`

Operation ID: `updateZoneEntrypointRuleset`

Updates a zone entry point ruleset, creating a new version.

## Definition

```yaml
{"operationId": "updateZoneEntrypointRuleset", "summary": "Update a zone entry point ruleset", "description": "Updates a zone entry point ruleset, creating a new version.", "parameters": [{"name": "ruleset_phase", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetPhase"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "requestBody": {"$ref": "#/components/requestBodies/rulesets_UpdateEntrypointRuleset"}, "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Rulesets"], "x-api-token-group": ["Response Compression Write", "Config Settings Write", "Dynamic URL Redirects Write", "Cache Settings Write", "Custom Errors Write", "Origin Write", "Managed headers Write", "Zone Transform Rules Write", "Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Write", "Sanitize Write", "Transform Rules Write", "Select Configuration Write", "Bot Management Write", "Zone WAF Write", "Account WAF Write", "Account Rulesets Write", "Logs Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
