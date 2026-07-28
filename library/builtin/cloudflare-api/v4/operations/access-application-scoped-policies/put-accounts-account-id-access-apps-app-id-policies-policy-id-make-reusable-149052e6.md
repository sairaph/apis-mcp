---
title: Convert an Access application policy to a reusable policy
page_id: operation-put-accounts-account-id-access-apps-app-id-policies-policy-id-make-reusa-7fb6b2e5
path: operations/access-application-scoped-policies
description: Converts an application-scoped policy to a reusable policy. The policy will no longer be exclusively scoped to the application. Further updates to the policy should go through the /accounts/{account_id}/policies/{uid} endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/policies/{policy_id}/make_reusable
operation_ids:
    - access-policies-convert-reusable
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Convert an Access application policy to a reusable policy

`PUT /accounts/{account_id}/access/apps/{app_id}/policies/{policy_id}/make_reusable`

Operation ID: `access-policies-convert-reusable`

Converts an application-scoped policy to a reusable policy. The policy will no longer be exclusively scoped to the application. Further updates to the policy should go through the /accounts/{account_id}/policies/{uid} endpoint.

## Definition

```yaml
{"operationId": "access-policies-convert-reusable", "summary": "Convert an Access application policy to a reusable policy", "description": "Converts an application-scoped policy to a reusable policy. The policy will no longer be exclusively scoped to the application. Further updates to the policy should go through the /accounts/{account_id}/policies/{uid} endpoint.", "parameters": [{"name": "app_id", "in": "path", "description": "The application ID.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "policy_id", "in": "path", "description": "The policy ID.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Convert an Access application policy to a reusable policy", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-8"}}}}, "4XX": {"description": "Convert an Access application policy to a reusable policy failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access application-scoped policies"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.apps.policies.make.reusable", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
