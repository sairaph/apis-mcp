---
title: Add an Access application
page_id: operation-post-zones-zone-id-access-apps-2356518f
path: operations/zone-level-access-applications
description: Adds a new application to Access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/access/apps
operation_ids:
    - zone-level-access-applications-add-a-bookmark-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add an Access application

`POST /zones/{zone_id}/access/apps`

Operation ID: `zone-level-access-applications-add-a-bookmark-application`

Adds a new application to Access.

## Definition

```yaml
{"operationId": "zone-level-access-applications-add-a-bookmark-application", "summary": "Add an Access application", "description": "Adds a new application to Access.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_apps"}}}}, "responses": {"201": {"description": "Add an Access application response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/access_single_response-21"}, {"properties": {"result": {"$ref": "#/components/schemas/access_apps"}}}]}}}}, "4XX": {"description": "Add an Access application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access applications"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
