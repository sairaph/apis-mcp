---
title: Delete URL Normalization settings
page_id: operation-delete-zones-zone-id-url-normalization-819423f0
path: operations/url-normalization
description: Deletes the URL Normalization settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/url_normalization
operation_ids:
    - deleteUrlNormalization
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete URL Normalization settings

`DELETE /zones/{zone_id}/url_normalization`

Operation ID: `deleteUrlNormalization`

Deletes the URL Normalization settings.

## Definition

```yaml
{"operationId": "deleteUrlNormalization", "summary": "Delete URL Normalization settings", "description": "Deletes the URL Normalization settings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "responses": {"204": {"$ref": "#/components/responses/rulesets_Empty"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["URL Normalization"], "x-api-token-group": ["Response Compression Write", "Config Settings Write", "Dynamic URL Redirects Write", "Cache Settings Write", "Custom Errors Write", "Origin Write", "Managed headers Write", "Zone Transform Rules Write", "Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Write", "Sanitize Write", "Transform Rules Write", "Select Configuration Write", "Bot Management Write", "Zone WAF Write", "Account WAF Write", "Account Rulesets Write", "Logs Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
