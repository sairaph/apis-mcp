---
title: Delete Managed Transforms
page_id: operation-delete-zones-zone-id-managed-headers-e207dbaa
path: operations/managed-transforms
description: Disables all Managed Transforms.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/managed_headers
operation_ids:
    - deleteManagedTransforms
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Managed Transforms

`DELETE /zones/{zone_id}/managed_headers`

Operation ID: `deleteManagedTransforms`

Disables all Managed Transforms.

## Definition

```yaml
{"operationId": "deleteManagedTransforms", "summary": "Delete Managed Transforms", "description": "Disables all Managed Transforms.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "responses": {"204": {"$ref": "#/components/responses/rulesets_Empty"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Managed Transforms"], "x-api-token-group": ["Response Compression Write", "Config Settings Write", "Dynamic URL Redirects Write", "Cache Settings Write", "Custom Errors Write", "Origin Write", "Managed headers Write", "Zone Transform Rules Write", "Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Write", "Sanitize Write", "Transform Rules Write", "Select Configuration Write", "Bot Management Write", "Zone WAF Write", "Account WAF Write", "Account Rulesets Write", "Logs Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
