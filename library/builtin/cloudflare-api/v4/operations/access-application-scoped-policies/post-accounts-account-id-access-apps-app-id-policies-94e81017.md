---
title: Create an Access application policy
page_id: operation-post-accounts-account-id-access-apps-app-id-policies-66de2420
path: operations/access-application-scoped-policies
description: Creates a policy applying exclusive to a single application that defines the users or groups who can reach it. We recommend creating a reusable policy instead and subsequently referencing its ID in the application's 'policies' array.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/policies
operation_ids:
    - access-policies-create-an-access-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an Access application policy

`POST /accounts/{account_id}/access/apps/{app_id}/policies`

Operation ID: `access-policies-create-an-access-policy`

Creates a policy applying exclusive to a single application that defines the users or groups who can reach it. We recommend creating a reusable policy instead and subsequently referencing its ID in the application's 'policies' array.

## Definition

```yaml
{"operationId": "access-policies-create-an-access-policy", "summary": "Create an Access application policy", "description": "Creates a policy applying exclusive to a single application that defines the users or groups who can reach it. We recommend creating a reusable policy instead and subsequently referencing its ID in the application's 'policies' array.", "parameters": [{"name": "app_id", "in": "path", "description": "The application ID.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_app_policy_request"}}}}, "responses": {"201": {"description": "Create an Access application policy response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-9"}}}}, "4XX": {"description": "Create an Access application policy response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access application-scoped policies"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.policies", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
