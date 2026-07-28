---
title: Update application settings
page_id: operation-patch-zones-zone-id-access-apps-app-id-settings-393748ce
path: operations/zone-level-access-applications
description: Updates application settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/access/apps/{app_id}/settings
operation_ids:
    - zone-level-access-applications-patch-update-access-application-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update application settings

`PATCH /zones/{zone_id}/access/apps/{app_id}/settings`

Operation ID: `zone-level-access-applications-patch-update-access-application-settings`

Updates application settings.

## Definition

```yaml
{"operationId": "zone-level-access-applications-patch-update-access-application-settings", "summary": "Update application settings", "description": "Updates application settings.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_app_id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_app_settings_request"}}}}, "responses": {"202": {"description": "Update application settings response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/access_single_response_update"}]}}}}, "4XX": {"description": "Update application settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access applications"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
