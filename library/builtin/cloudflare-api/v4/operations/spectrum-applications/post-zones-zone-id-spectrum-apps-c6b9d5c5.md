---
title: Create Spectrum application using a name for the origin
page_id: operation-post-zones-zone-id-spectrum-apps-ee3b01e8
path: operations/spectrum-applications
description: Creates a new Spectrum application from a configuration using a name for the origin.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/spectrum/apps
operation_ids:
    - spectrum-applications-create-spectrum-application-using-a-name-for-the-origin
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Spectrum application using a name for the origin

`POST /zones/{zone_id}/spectrum/apps`

Operation ID: `spectrum-applications-create-spectrum-application-using-a-name-for-the-origin`

Creates a new Spectrum application from a configuration using a name for the origin.

## Definition

```yaml
{"operationId": "spectrum-applications-create-spectrum-application-using-a-name-for-the-origin", "summary": "Create Spectrum application using a name for the origin", "description": "Creates a new Spectrum application from a configuration using a name for the origin.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-config_zone_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_update_app_config"}}}}, "responses": {"200": {"description": "Create Spectrum application using a name for the origin response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_app_config_single"}}}}, "4XX": {"description": "Create Spectrum application using a name for the origin response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Spectrum Applications"], "x-api-token-group": ["Zone Settings Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "spectrum.apps", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
