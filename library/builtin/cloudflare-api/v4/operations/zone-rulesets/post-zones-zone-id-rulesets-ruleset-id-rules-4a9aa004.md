---
title: Create a zone ruleset rule
page_id: operation-post-zones-zone-id-rulesets-ruleset-id-rules-fb4d13fb
path: operations/zone-rulesets
description: Adds a new rule to a zone ruleset. The rule will be added to the end of the existing list of rules in the ruleset by default.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/rulesets/{ruleset_id}/rules
operation_ids:
    - createZoneRulesetRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a zone ruleset rule

`POST /zones/{zone_id}/rulesets/{ruleset_id}/rules`

Operation ID: `createZoneRulesetRule`

Adds a new rule to a zone ruleset. The rule will be added to the end of the existing list of rules in the ruleset by default.

## Definition

```yaml
{"operationId": "createZoneRulesetRule", "summary": "Create a zone ruleset rule", "description": "Adds a new rule to a zone ruleset. The rule will be added to the end of the existing list of rules in the ruleset by default.", "parameters": [{"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "requestBody": {"$ref": "#/components/requestBodies/rulesets_Rule"}, "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Rulesets"], "x-api-token-group": ["Response Compression Write", "Config Settings Write", "Dynamic URL Redirects Write", "Cache Settings Write", "Custom Errors Write", "Origin Write", "Managed headers Write", "Zone Transform Rules Write", "Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Write", "Sanitize Write", "Transform Rules Write", "Select Configuration Write", "Bot Management Write", "Zone WAF Write", "Account WAF Write", "Account Rulesets Write", "Logs Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rulesets.rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
