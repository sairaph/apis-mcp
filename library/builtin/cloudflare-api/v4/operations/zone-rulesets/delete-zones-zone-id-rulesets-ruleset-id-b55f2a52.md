---
title: Delete a zone ruleset
page_id: operation-delete-zones-zone-id-rulesets-ruleset-id-e7b6a984
path: operations/zone-rulesets
description: Deletes all versions of an existing zone ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/rulesets/{ruleset_id}
operation_ids:
    - deleteZoneRuleset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a zone ruleset

`DELETE /zones/{zone_id}/rulesets/{ruleset_id}`

Operation ID: `deleteZoneRuleset`

Deletes all versions of an existing zone ruleset.

## Definition

```yaml
{"operationId": "deleteZoneRuleset", "summary": "Delete a zone ruleset", "description": "Deletes all versions of an existing zone ruleset.", "parameters": [{"name": "ruleset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_RulesetId"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "responses": {"204": {"$ref": "#/components/responses/rulesets_Empty"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Rulesets"], "x-api-token-group": ["Response Compression Write", "Config Settings Write", "Dynamic URL Redirects Write", "Cache Settings Write", "Custom Errors Write", "Origin Write", "Managed headers Write", "Zone Transform Rules Write", "Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Write", "Sanitize Write", "Transform Rules Write", "Select Configuration Write", "Bot Management Write", "Zone WAF Write", "Account WAF Write", "Account Rulesets Write", "Logs Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rulesets", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
