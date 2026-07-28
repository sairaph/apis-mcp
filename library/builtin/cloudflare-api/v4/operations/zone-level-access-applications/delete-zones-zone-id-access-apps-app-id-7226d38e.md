---
title: Delete an Access application
page_id: operation-delete-zones-zone-id-access-apps-app-id-64d58292
path: operations/zone-level-access-applications
description: Deletes an application from Access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}
operation_ids:
    - zone-level-access-applications-delete-an-access-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an Access application

`DELETE /zones/{zone_id}/access/apps/{app_id}`

Operation ID: `zone-level-access-applications-delete-an-access-application`

Deletes an application from Access.

## Definition

```yaml
{"operationId": "zone-level-access-applications-delete-an-access-application", "summary": "Delete an Access application", "description": "Deletes an application from Access.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_app_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete an Access application response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response"}}}}, "4XX": {"description": "Delete an Access application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access applications"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
