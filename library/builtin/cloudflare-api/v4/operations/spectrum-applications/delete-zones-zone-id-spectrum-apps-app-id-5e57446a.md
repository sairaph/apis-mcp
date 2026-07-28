---
title: Delete Spectrum application
page_id: operation-delete-zones-zone-id-spectrum-apps-app-id-f165a57a
path: operations/spectrum-applications
description: Deletes a previously existing application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/spectrum/apps/{app_id}
operation_ids:
    - spectrum-applications-delete-spectrum-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Spectrum application

`DELETE /zones/{zone_id}/spectrum/apps/{app_id}`

Operation ID: `spectrum-applications-delete-spectrum-application`

Deletes a previously existing application.

## Definition

```yaml
{"operationId": "spectrum-applications-delete-spectrum-application", "summary": "Delete Spectrum application", "description": "Deletes a previously existing application.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-config_app_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/spectrum-config_zone_identifier"}}], "responses": {"200": {"description": "Delete Spectrum application response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_api-response-single-id"}}}}, "4XX": {"description": "Delete Spectrum application response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/spectrum-config_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Spectrum Applications"], "x-api-token-group": ["Zone Settings Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "spectrum.apps", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
