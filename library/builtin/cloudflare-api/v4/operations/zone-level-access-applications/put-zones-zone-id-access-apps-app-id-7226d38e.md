---
title: Update an Access application
page_id: operation-put-zones-zone-id-access-apps-app-id-d1bd75f7
path: operations/zone-level-access-applications
description: Updates an Access application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}
operation_ids:
    - zone-level-access-applications-update-a-bookmark-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an Access application

`PUT /zones/{zone_id}/access/apps/{app_id}`

Operation ID: `zone-level-access-applications-update-a-bookmark-application`

Updates an Access application.

## Definition

```yaml
{"operationId": "zone-level-access-applications-update-a-bookmark-application", "summary": "Update an Access application", "description": "Updates an Access application.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_app_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_apps"}}}}, "responses": {"200": {"description": "Update an Access application response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/access_single_response-21"}, {"properties": {"result": {"$ref": "#/components/schemas/access_apps"}}}]}}}}, "4XX": {"description": "Update an Access application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access applications"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps", "x-fern-sdk-method-name": "update", "x-forge-hidden": true, "x-stability": "beta"}
```
