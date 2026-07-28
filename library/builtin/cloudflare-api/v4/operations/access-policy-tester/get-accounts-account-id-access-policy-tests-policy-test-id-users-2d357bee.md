---
title: Get an Access policy test users page
page_id: operation-get-accounts-account-id-access-policy-tests-policy-test-id-users-c22b8f24
path: operations/access-policy-tester
description: Fetches a single page of user results from an Access policy test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/policy-tests/{policy_test_id}/users
operation_ids:
    - access-policy-tests-get-a-user-page
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access policy test users page

`GET /accounts/{account_id}/access/policy-tests/{policy_test_id}/users`

Operation ID: `access-policy-tests-get-a-user-page`

Fetches a single page of user results from an Access policy test.

## Definition

```yaml
{"operationId": "access-policy-tests-get-a-user-page", "summary": "Get an Access policy test users page", "description": "Fetches a single page of user results from an Access policy test.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "policy_test_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_policy_test_id"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 25, "maximum": 1000}}, {"name": "status", "in": "query", "description": "Filter users by their policy evaluation status.", "schema": {"type": "string", "enum": ["success", "fail", "error"]}}], "responses": {"200": {"description": "Get an Access policy tester users page response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_policy_users_resp"}}}}, "400": {"description": "Get an Access policy tester users page response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access policy tester"], "x-api-token-group": ["Access: Policy Test Write", "Access: Policy Test Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.policy-tests.users", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
