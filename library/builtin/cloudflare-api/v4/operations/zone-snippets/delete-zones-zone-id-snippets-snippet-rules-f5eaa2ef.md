---
title: Delete zone snippet rules
page_id: operation-delete-zones-zone-id-snippets-snippet-rules-239abaf2
path: operations/zone-snippets
description: Deletes all snippet rules belonging to the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/snippets/snippet_rules
operation_ids:
    - deleteZoneSnippetRules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete zone snippet rules

`DELETE /zones/{zone_id}/snippets/snippet_rules`

Operation ID: `deleteZoneSnippetRules`

Deletes all snippet rules belonging to the zone.

## Definition

```yaml
{"operationId": "deleteZoneSnippetRules", "summary": "Delete zone snippet rules", "description": "Deletes all snippet rules belonging to the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/snippets_ZoneId"}}], "responses": {"200": {"$ref": "#/components/responses/snippets_SnippetRules"}, "4XX": {"$ref": "#/components/responses/snippets_Failure"}, "5XX": {"$ref": "#/components/responses/snippets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Snippets"], "x-api-token-group": ["Snippets Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
