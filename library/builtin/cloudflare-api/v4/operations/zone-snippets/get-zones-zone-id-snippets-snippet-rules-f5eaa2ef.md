---
title: List zone snippet rules
page_id: operation-get-zones-zone-id-snippets-snippet-rules-0224227a
path: operations/zone-snippets
description: Fetches all snippet rules belonging to the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/snippets/snippet_rules
operation_ids:
    - listZoneSnippetRules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List zone snippet rules

`GET /zones/{zone_id}/snippets/snippet_rules`

Operation ID: `listZoneSnippetRules`

Fetches all snippet rules belonging to the zone.

## Definition

```yaml
{"operationId": "listZoneSnippetRules", "summary": "List zone snippet rules", "description": "Fetches all snippet rules belonging to the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/snippets_ZoneId"}}], "responses": {"200": {"$ref": "#/components/responses/snippets_SnippetRules"}, "4XX": {"$ref": "#/components/responses/snippets_Failure"}, "5XX": {"$ref": "#/components/responses/snippets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Snippets"], "x-api-token-group": ["Snippets Write", "Snippets Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
