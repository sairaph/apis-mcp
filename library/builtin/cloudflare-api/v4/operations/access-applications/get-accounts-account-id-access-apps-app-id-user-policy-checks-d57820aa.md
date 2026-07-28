---
title: Test Access policies
page_id: operation-get-accounts-account-id-access-apps-app-id-user-policy-checks-f94dd12b
path: operations/access-applications
description: Tests if a specific user has permission to access an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/apps/{app_id}/user_policy_checks
operation_ids:
    - access-applications-test-access-policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Test Access policies

`GET /accounts/{account_id}/access/apps/{app_id}/user_policy_checks`

Operation ID: `access-applications-test-access-policies`

Tests if a specific user has permission to access an application.

## Definition

```yaml
{"operationId": "access-applications-test-access-policies", "summary": "Test Access policies", "description": "Tests if a specific user has permission to access an application.", "parameters": [{"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_app_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Test Access policies response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_policy_check_response"}}}}, "4XX": {"description": "Test Access policies response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access applications"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.applications.user-policy-checks", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
