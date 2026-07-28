---
title: Start Access policy test
page_id: operation-post-accounts-account-id-access-policy-tests-37664088
path: operations/access-policy-tester
description: Starts an Access policy test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/policy-tests
operation_ids:
    - access-policy-tests
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start Access policy test

`POST /accounts/{account_id}/access/policy-tests`

Operation ID: `access-policy-tests`

Starts an Access policy test.

## Definition

```yaml
{"operationId": "access-policy-tests", "summary": "Start Access policy test", "description": "Starts an Access policy test.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_policy_init_req"}}}}, "responses": {"200": {"description": "Start Access policy test response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_policy_init_resp"}}}}, "400": {"description": "Start Access policy test response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access policy tester"], "x-api-token-group": ["Access: Policy Test Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.policy-tests", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
