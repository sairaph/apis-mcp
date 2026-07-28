---
title: Delete an Access application policy
page_id: operation-delete-accounts-account-id-access-apps-app-id-policies-policy-id-9cdb5e8b
path: operations/access-application-scoped-policies
description: Deletes an Access policy specific to an application. To delete a reusable policy, use the /accounts/{account_id}/policies/{uid} endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/policies/{policy_id}
operation_ids:
    - access-policies-delete-an-access-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an Access application policy

`DELETE /accounts/{account_id}/access/apps/{app_id}/policies/{policy_id}`

Operation ID: `access-policies-delete-an-access-policy`

Deletes an Access policy specific to an application. To delete a reusable policy, use the /accounts/{account_id}/policies/{uid} endpoint.

## Definition

```yaml
{"operationId": "access-policies-delete-an-access-policy", "summary": "Delete an Access application policy", "description": "Deletes an Access policy specific to an application. To delete a reusable policy, use the /accounts/{account_id}/policies/{uid} endpoint.", "parameters": [{"name": "app_id", "in": "path", "description": "The application ID.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "policy_id", "in": "path", "description": "The policy ID.", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete an Access application policy response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response-4"}}}}, "4XX": {"description": "Delete an Access application policy response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access application-scoped policies"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.policies", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
