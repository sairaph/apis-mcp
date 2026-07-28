---
title: Get the current status of a given Access policy test
page_id: operation-get-accounts-account-id-access-policy-tests-policy-test-id-e489a757
path: operations/access-policy-tester
description: Fetches the current status of a given Access policy test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/policy-tests/{policy_test_id}
operation_ids:
    - access-policy-tests-get-an-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the current status of a given Access policy test

`GET /accounts/{account_id}/access/policy-tests/{policy_test_id}`

Operation ID: `access-policy-tests-get-an-update`

Fetches the current status of a given Access policy test.

## Definition

```yaml
{"operationId": "access-policy-tests-get-an-update", "summary": "Get the current status of a given Access policy test", "description": "Fetches the current status of a given Access policy test.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "policy_test_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_policy_test_id"}}], "responses": {"200": {"description": "Get an Access policy test update response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_policy_update_resp"}}}}, "400": {"description": "Get an Access policy test update response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access policy tester"], "x-api-token-group": ["Access: Policy Test Write", "Access: Policy Test Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.policy-tests", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
