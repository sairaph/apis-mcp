---
title: Add an Access application
page_id: operation-post-accounts-account-id-access-apps-1d65a38d
path: operations/access-applications
description: Adds a new application to Access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/apps
operation_ids:
    - access-applications-add-an-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add an Access application

`POST /accounts/{account_id}/access/apps`

Operation ID: `access-applications-add-an-application`

Adds a new application to Access.

## Definition

```yaml
{"operationId": "access-applications-add-an-application", "summary": "Add an Access application", "description": "Adds a new application to Access.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_app_request"}}}}, "responses": {"201": {"description": "Add an Access application response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/access_single_response-8"}, {"properties": {"result": {"$ref": "#/components/schemas/access_app_response"}}}]}}}}, "4XX": {"description": "Add an Access application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access applications"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications", "x-fern-sdk-method-name": "create", "x-forge-hidden": true, "x-stability": "beta"}
```
