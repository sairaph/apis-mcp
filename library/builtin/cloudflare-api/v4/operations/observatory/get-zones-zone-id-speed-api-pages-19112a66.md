---
title: List tested webpages
page_id: operation-get-zones-zone-id-speed-api-pages-d341afa4
path: operations/observatory
description: Lists all webpages which have been tested.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/speed_api/pages
operation_ids:
    - speed-list-pages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tested webpages

`GET /zones/{zone_id}/speed_api/pages`

Operation ID: `speed-list-pages`

Lists all webpages which have been tested.

## Definition

```yaml
{"operationId": "speed-list-pages", "summary": "List tested webpages", "description": "Lists all webpages which have been tested.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/observatory_identifier"}}], "responses": {"200": {"description": "List of pages.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_pages-response-collection"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/observatory_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Observatory"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "speed.pages", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
