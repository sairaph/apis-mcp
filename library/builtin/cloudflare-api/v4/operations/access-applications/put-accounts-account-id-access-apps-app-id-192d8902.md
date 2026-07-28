---
title: Update an Access application
page_id: operation-put-accounts-account-id-access-apps-app-id-c02169e0
path: operations/access-applications
description: Updates an Access application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}
operation_ids:
    - access-applications-update-an-access-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an Access application

`PUT /accounts/{account_id}/access/apps/{app_id}`

Operation ID: `access-applications-update-an-access-application`

Updates an Access application.

## Definition

```yaml
{"operationId": "access-applications-update-an-access-application", "summary": "Update an Access application", "description": "Updates an Access application.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_app_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_app_request"}}}}, "responses": {"200": {"description": "Update an Access application response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/access_single_response-8"}, {"properties": {"result": {"$ref": "#/components/schemas/access_app_response"}}}]}}}}, "4XX": {"description": "Update an Access application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access applications"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications", "x-fern-sdk-method-name": "update", "x-forge-hidden": true, "x-stability": "beta"}
```
