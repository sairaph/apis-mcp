---
title: Update Spectrum application configuration using a name for the origin
page_id: operation-put-zones-zone-id-spectrum-apps-app-id-a4e1cb7b
path: operations/spectrum-applications
description: Updates a previously existing application's configuration that uses a name for the origin.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/spectrum/apps/{app_id}
operation_ids:
    - spectrum-applications-update-spectrum-application-configuration-using-a-name-for-the-origin
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Spectrum application configuration using a name for the origin

`PUT /zones/{zone_id}/spectrum/apps/{app_id}`

Operation ID: `spectrum-applications-update-spectrum-application-configuration-using-a-name-for-the-origin`

Updates a previously existing application's configuration that uses a name for the origin.

## Definition

```yaml
{"operationId": "spectrum-applications-update-spectrum-application-configuration-using-a-name-for-the-origin", "summary": "Update Spectrum application configuration using a name for the origin", "description": "Updates a previously existing application's configuration that uses a name for the origin.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-config_app_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-config_zone_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_update_app_config"}}}}, "responses": {"200": {"description": "Update Spectrum application configuration using a name for the origin response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_app_config_single"}}}}, "4XX": {"description": "Update Spectrum application configuration using a name for the origin response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Spectrum Applications"], "x-api-token-group": ["Zone Settings Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "spectrum.apps", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
