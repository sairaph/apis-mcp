---
title: Update an Access reusable policy
page_id: operation-put-accounts-account-id-access-policies-policy-id-da48fcbd
path: operations/access-reusable-policies
description: Updates a Access reusable policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/policies/{policy_id}
operation_ids:
    - access-policies-update-an-access-reusable-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an Access reusable policy

`PUT /accounts/{account_id}/access/policies/{policy_id}`

Operation ID: `access-policies-update-an-access-reusable-policy`

Updates a Access reusable policy.

## Definition

```yaml
{"operationId": "access-policies-update-an-access-reusable-policy", "summary": "Update an Access reusable policy", "description": "Updates a Access reusable policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_policy_req"}}}}, "responses": {"200": {"description": "Update an Access reusable policy response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-10"}}}}, "4XX": {"description": "Update an Access reusable policy response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access reusable policies"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.policies", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
