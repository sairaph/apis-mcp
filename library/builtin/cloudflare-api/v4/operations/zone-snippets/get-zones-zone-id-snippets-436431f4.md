---
title: List zone snippets
page_id: operation-get-zones-zone-id-snippets-e584cf96
path: operations/zone-snippets
description: Fetches all snippets belonging to the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/snippets
operation_ids:
    - listZoneSnippets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List zone snippets

`GET /zones/{zone_id}/snippets`

Operation ID: `listZoneSnippets`

Fetches all snippets belonging to the zone.

## Definition

```yaml
{"operationId": "listZoneSnippets", "summary": "List zone snippets", "description": "Fetches all snippets belonging to the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/snippets_ZoneId"}}, {"name": "page", "in": "query", "schema": {"$ref": "#/components/schemas/snippets_Page"}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/snippets_PerPage"}}], "responses": {"200": {"$ref": "#/components/responses/snippets_Snippets"}, "4XX": {"$ref": "#/components/responses/snippets_Failure"}, "5XX": {"$ref": "#/components/responses/snippets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Snippets"], "x-api-token-group": ["Snippets Write", "Snippets Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
