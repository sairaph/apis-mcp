---
title: Update zone snippet rules
page_id: operation-put-zones-zone-id-snippets-snippet-rules-37cd4650
path: operations/zone-snippets
description: Updates all snippet rules belonging to the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/snippets/snippet_rules
operation_ids:
    - updateZoneSnippetRules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update zone snippet rules

`PUT /zones/{zone_id}/snippets/snippet_rules`

Operation ID: `updateZoneSnippetRules`

Updates all snippet rules belonging to the zone.

## Definition

```yaml
{"operationId": "updateZoneSnippetRules", "summary": "Update zone snippet rules", "description": "Updates all snippet rules belonging to the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/snippets_ZoneId"}}], "requestBody": {"$ref": "#/components/requestBodies/snippets_SnippetRules"}, "responses": {"200": {"$ref": "#/components/responses/snippets_SnippetRules"}, "4XX": {"$ref": "#/components/responses/snippets_Failure"}, "5XX": {"$ref": "#/components/responses/snippets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Snippets"], "x-api-token-group": ["Snippets Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
