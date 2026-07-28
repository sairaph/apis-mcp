---
title: Update Access application settings
page_id: operation-patch-accounts-account-id-access-apps-app-id-settings-c15fa717
path: operations/access-applications
description: Updates Access application settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/settings
operation_ids:
    - access-applications-patch-update-access-application-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Access application settings

`PATCH /accounts/{account_id}/access/apps/{app_id}/settings`

Operation ID: `access-applications-patch-update-access-application-settings`

Updates Access application settings.

## Definition

```yaml
{"operationId": "access-applications-patch-update-access-application-settings", "summary": "Update Access application settings", "description": "Updates Access application settings.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_app_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_app_settings_request"}}}}, "responses": {"202": {"description": "Update Access application settings response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/access_single_response_update"}]}}}}, "4XX": {"description": "Update Access application settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access applications"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
