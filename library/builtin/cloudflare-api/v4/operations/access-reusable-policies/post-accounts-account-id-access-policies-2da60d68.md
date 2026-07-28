---
title: Create an Access reusable policy
page_id: operation-post-accounts-account-id-access-policies-00dfdbb2
path: operations/access-reusable-policies
description: Creates a new Access reusable policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/policies
operation_ids:
    - access-policies-create-an-access-reusable-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an Access reusable policy

`POST /accounts/{account_id}/access/policies`

Operation ID: `access-policies-create-an-access-reusable-policy`

Creates a new Access reusable policy.

## Definition

```yaml
{"operationId": "access-policies-create-an-access-reusable-policy", "summary": "Create an Access reusable policy", "description": "Creates a new Access reusable policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_policy_req"}}}}, "responses": {"201": {"description": "Create an Access reusable policy response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-10"}}}}, "4XX": {"description": "Create an Access reusable policy response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access reusable policies"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.policies", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
