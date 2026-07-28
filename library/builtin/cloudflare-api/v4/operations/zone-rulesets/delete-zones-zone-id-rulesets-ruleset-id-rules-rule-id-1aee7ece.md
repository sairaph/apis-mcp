---
title: Delete a zone ruleset rule
page_id: operation-delete-zones-zone-id-rulesets-ruleset-id-rules-rule-id-2ee4aa7b
path: operations/zone-rulesets
description: Deletes an existing rule from a zone ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/rulesets/{ruleset_id}/rules/{rule_id}
operation_ids:
    - deleteZoneRulesetRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a zone ruleset rule

`DELETE /zones/{zone_id}/rulesets/{ruleset_id}/rules/{rule_id}`

Operation ID: `deleteZoneRulesetRule`

Deletes an existing rule from a zone ruleset.

## Definition

```yaml
{"operationId": "deleteZoneRulesetRule", "summary": "Delete a zone ruleset rule", "description": "Deletes an existing rule from a zone ruleset.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RuleId"}}, {"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "responses": {"200": {"$ref": "#/components/responses/rulesets_Ruleset"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Rulesets"], "x-api-token-group": ["Response Compression Write", "Config Settings Write", "Dynamic URL Redirects Write", "Cache Settings Write", "Custom Errors Write", "Origin Write", "Managed headers Write", "Zone Transform Rules Write", "Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Write", "Sanitize Write", "Transform Rules Write", "Select Configuration Write", "Bot Management Write", "Zone WAF Write", "Account WAF Write", "Account Rulesets Write", "Logs Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rulesets.rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
