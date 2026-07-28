---
title: Delete a zone snippet
page_id: operation-delete-zones-zone-id-snippets-snippet-name-9c68411b
path: operations/zone-snippets
description: Deletes a snippet belonging to the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/snippets/{snippet_name}
operation_ids:
    - deleteZoneSnippet
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a zone snippet

`DELETE /zones/{zone_id}/snippets/{snippet_name}`

Operation ID: `deleteZoneSnippet`

Deletes a snippet belonging to the zone.

## Definition

```yaml
{"operationId": "deleteZoneSnippet", "summary": "Delete a zone snippet", "description": "Deletes a snippet belonging to the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/snippets_ZoneId"}}, {"name": "snippet_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/snippets_SnippetName"}}], "responses": {"200": {"$ref": "#/components/responses/snippets_Null"}, "4XX": {"$ref": "#/components/responses/snippets_Failure"}, "5XX": {"$ref": "#/components/responses/snippets_Failure"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Snippets"], "x-api-token-group": ["Snippets Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
