---
title: Get an Access application policy
page_id: operation-get-accounts-account-id-access-apps-app-id-policies-policy-id-dd3e7efc
path: operations/access-application-scoped-policies
description: Fetches a single Access policy configured for an application. Returns both exclusively owned and reusable policies used by the application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/policies/{policy_id}
operation_ids:
    - access-policies-get-an-access-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access application policy

`GET /accounts/{account_id}/access/apps/{app_id}/policies/{policy_id}`

Operation ID: `access-policies-get-an-access-policy`

Fetches a single Access policy configured for an application. Returns both exclusively owned and reusable policies used by the application.

## Definition

```yaml
{"operationId": "access-policies-get-an-access-policy", "summary": "Get an Access application policy", "description": "Fetches a single Access policy configured for an application. Returns both exclusively owned and reusable policies used by the application.", "parameters": [{"name": "app_id", "in": "path", "description": "The application ID.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "policy_id", "in": "path", "description": "The policy ID.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get an Access policy response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-9"}}}}, "4XX": {"description": "Get an Access policy response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Access application-scoped policies"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.policies", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
