---
title: List Notification policies
page_id: operation-get-accounts-account-id-alerting-v3-policies-8b8bf80e
path: operations/notification-policies
description: Get a list of all Notification policies.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/policies
operation_ids:
    - notification-policies-list-notification-policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Notification policies

`GET /accounts/{account_id}/alerting/v3/policies`

Operation ID: `notification-policies-list-notification-policies`

Get a list of all Notification policies.

## Definition

```yaml
{"operationId": "notification-policies-list-notification-policies", "summary": "List Notification policies", "description": "Get a list of all Notification policies.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"200": {"description": "List Notification policies response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_policies_components-schemas-response_collection"}}}}, "4XX": {"description": "List Notification policies response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification policies"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"]}
```
