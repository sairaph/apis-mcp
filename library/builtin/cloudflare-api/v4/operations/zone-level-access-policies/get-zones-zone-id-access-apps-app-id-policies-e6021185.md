---
title: List Access policies
page_id: operation-get-zones-zone-id-access-apps-app-id-policies-f3c8c9b9
path: operations/zone-level-access-policies
description: Lists Access policies configured for an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}/policies
operation_ids:
    - zone-level-access-policies-list-access-policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access policies

`GET /zones/{zone_id}/access/apps/{app_id}/policies`

Operation ID: `zone-level-access-policies-list-access-policies`

Lists Access policies configured for an application.

## Definition

```yaml
{"operationId": "zone-level-access-policies-list-access-policies", "summary": "List Access policies", "description": "Lists Access policies configured for an application.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List Access policies response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-23"}}}}, "4XX": {"description": "List Access policies response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access policies"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.policies", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
