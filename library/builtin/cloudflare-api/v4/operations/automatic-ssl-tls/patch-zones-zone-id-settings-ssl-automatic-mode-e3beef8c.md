---
title: Patch Automatic SSL/TLS Enrollment status for given zone
page_id: operation-patch-zones-zone-id-settings-ssl-automatic-mode-9e84a9b5
path: operations/automatic-ssl-tls
description: The automatic system is enabled when this endpoint is hit with value in the request body is set to "auto", and disabled when the request body value is set to "custom".
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/ssl_automatic_mode
operation_ids:
    - ssl-detector-automatic-mode-patch-enrollment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Automatic SSL/TLS Enrollment status for given zone

`PATCH /zones/{zone_id}/settings/ssl_automatic_mode`

Operation ID: `ssl-detector-automatic-mode-patch-enrollment`

The automatic system is enabled when this endpoint is hit with value in the request body is set to "auto", and disabled when the request body value is set to "custom".

## Definition

```yaml
{"operationId": "ssl-detector-automatic-mode-patch-enrollment", "summary": "Patch Automatic SSL/TLS Enrollment status for given zone", "description": "The automatic system is enabled when this endpoint is hit with value in the request body is set to \"auto\", and disabled when the request body value is set to \"custom\".", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache_schemas_patch"}}}}, "responses": {"200": {"description": "Patch Automatic SSL/TLS Enrollment status response.", "content": {"application/json": {"examples": {"Enabled": {"$ref": "#/components/examples/cache_dummy_automatic_upgrader_enabled_response"}}, "schema": {"$ref": "#/components/schemas/cache_api-response-single-id"}}}}, "4XX": {"description": "Patch Automatic SSL/TLS Enrollment status failure.", "content": {"application/json": {"examples": {"Error": {"$ref": "#/components/examples/cache_dummy_error_response"}}, "schema": {"$ref": "#/components/schemas/cache_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Automatic SSL/TLS"], "x-api-token-group": ["Zone Settings Write", "SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read", "#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
