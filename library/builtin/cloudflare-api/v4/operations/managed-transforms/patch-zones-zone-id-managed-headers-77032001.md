---
title: Update Managed Transforms
page_id: operation-patch-zones-zone-id-managed-headers-c58bae12
path: operations/managed-transforms
description: Updates the status of one or more Managed Transforms.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/managed_headers
operation_ids:
    - updateManagedTransforms
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Managed Transforms

`PATCH /zones/{zone_id}/managed_headers`

Operation ID: `updateManagedTransforms`

Updates the status of one or more Managed Transforms.

## Definition

```yaml
{"operationId": "updateManagedTransforms", "summary": "Update Managed Transforms", "description": "Updates the status of one or more Managed Transforms.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rulesets_ZoneId"}}], "requestBody": {"$ref": "#/components/requestBodies/rulesets_ManagedTransforms"}, "responses": {"200": {"$ref": "#/components/responses/rulesets_ManagedTransforms"}, "4XX": {"$ref": "#/components/responses/rulesets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Managed Transforms"], "x-api-token-group": ["Response Compression Write", "Config Settings Write", "Dynamic URL Redirects Write", "Cache Settings Write", "Custom Errors Write", "Origin Write", "Managed headers Write", "Zone Transform Rules Write", "Mass URL Redirects Write", "Magic Firewall Write", "L4 DDoS Managed Ruleset Write", "HTTP DDoS Managed Ruleset Write", "Sanitize Write", "Transform Rules Write", "Select Configuration Write", "Bot Management Write", "Zone WAF Write", "Account WAF Write", "Account Rulesets Write", "Logs Write", "Logs Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
