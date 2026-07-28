---
title: Get an Access application
page_id: operation-get-accounts-account-id-access-apps-app-id-655570dc
path: operations/access-applications
description: Fetches information about an Access application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}
operation_ids:
    - access-applications-get-an-access-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access application

`GET /accounts/{account_id}/access/apps/{app_id}`

Operation ID: `access-applications-get-an-access-application`

Fetches information about an Access application.

## Definition

```yaml
{"operationId": "access-applications-get-an-access-application", "summary": "Get an Access application", "description": "Fetches information about an Access application.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_app_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get an Access application response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-8"}}}}, "4XX": {"description": "Get an Access application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access applications"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications", "x-fern-sdk-method-name": "get", "x-forge-hidden": true, "x-stability": "beta"}
```
