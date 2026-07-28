---
title: List Access application policies
page_id: operation-get-accounts-account-id-access-apps-app-id-policies-5b58b3ad
path: operations/access-application-scoped-policies
description: Lists Access policies configured for an application. Returns both exclusively scoped and reusable policies used by the application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/policies
operation_ids:
    - access-policies-list-access-app-policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access application policies

`GET /accounts/{account_id}/access/apps/{app_id}/policies`

Operation ID: `access-policies-list-access-app-policies`

Lists Access policies configured for an application. Returns both exclusively scoped and reusable policies used by the application.

## Definition

```yaml
{"operationId": "access-policies-list-access-app-policies", "summary": "List Access application policies", "description": "Lists Access policies configured for an application. Returns both exclusively scoped and reusable policies used by the application.", "parameters": [{"name": "app_id", "in": "path", "description": "The application ID.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 200, "maximum": 1000}}], "responses": {"200": {"description": "List Access application policies response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-8"}}}}, "4XX": {"description": "List Access application policies response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access application-scoped policies"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.policies", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
