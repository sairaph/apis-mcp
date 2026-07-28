---
title: Get a zone snippet content
page_id: operation-get-zones-zone-id-snippets-snippet-name-content-ac3251fa
path: operations/zone-snippets
description: Fetches the content of a snippet belonging to the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/snippets/{snippet_name}/content
operation_ids:
    - getZoneSnippetContent
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a zone snippet content

`GET /zones/{zone_id}/snippets/{snippet_name}/content`

Operation ID: `getZoneSnippetContent`

Fetches the content of a snippet belonging to the zone.

## Definition

```yaml
{"operationId": "getZoneSnippetContent", "summary": "Get a zone snippet content", "description": "Fetches the content of a snippet belonging to the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/snippets_ZoneId"}}, {"name": "snippet_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/snippets_SnippetName"}}], "responses": {"200": {"$ref": "#/components/responses/snippets_SnippetContent"}, "4XX": {"$ref": "#/components/responses/snippets_Failure"}, "5XX": {"$ref": "#/components/responses/snippets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Snippets"], "x-api-token-group": ["Snippets Write", "Snippets Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
