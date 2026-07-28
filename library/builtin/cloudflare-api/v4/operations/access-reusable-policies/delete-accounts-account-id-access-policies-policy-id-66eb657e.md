---
title: Delete an Access reusable policy
page_id: operation-delete-accounts-account-id-access-policies-policy-id-c8197884
path: operations/access-reusable-policies
description: Deletes an Access reusable policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/policies/{policy_id}
operation_ids:
    - access-policies-delete-an-access-reusable-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an Access reusable policy

`DELETE /accounts/{account_id}/access/policies/{policy_id}`

Operation ID: `access-policies-delete-an-access-reusable-policy`

Deletes an Access reusable policy.

## Definition

```yaml
{"operationId": "access-policies-delete-an-access-reusable-policy", "summary": "Delete an Access reusable policy", "description": "Deletes an Access reusable policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid-2"}}], "responses": {"202": {"description": "Delete an Access reusable policy response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response-5"}}}}, "4XX": {"description": "Delete an Access reusable policy response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access reusable policies"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.policies", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
