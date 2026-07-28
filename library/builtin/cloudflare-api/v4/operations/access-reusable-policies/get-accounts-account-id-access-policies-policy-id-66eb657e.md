---
title: Get an Access reusable policy
page_id: operation-get-accounts-account-id-access-policies-policy-id-ce6e2215
path: operations/access-reusable-policies
description: Fetches a single Access reusable policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/policies/{policy_id}
operation_ids:
    - access-policies-get-an-access-reusable-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access reusable policy

`GET /accounts/{account_id}/access/policies/{policy_id}`

Operation ID: `access-policies-get-an-access-reusable-policy`

Fetches a single Access reusable policy.

## Definition

```yaml
{"operationId": "access-policies-get-an-access-reusable-policy", "summary": "Get an Access reusable policy", "description": "Fetches a single Access reusable policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid-2"}}], "responses": {"200": {"description": "Get an Access reusable policy response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-10"}}}}, "4XX": {"description": "Get an Access reusable policy response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access reusable policies"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.policies", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
