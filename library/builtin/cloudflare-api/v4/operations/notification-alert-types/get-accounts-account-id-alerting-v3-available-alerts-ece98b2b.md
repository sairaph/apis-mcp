---
title: Get Alert Types
page_id: operation-get-accounts-account-id-alerting-v3-available-alerts-fc0b29b3
path: operations/notification-alert-types
description: Gets a list of all alert types for which an account is eligible.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/available_alerts
operation_ids:
    - notification-alert-types-get-alert-types
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Alert Types

`GET /accounts/{account_id}/alerting/v3/available_alerts`

Operation ID: `notification-alert-types-get-alert-types`

Gets a list of all alert types for which an account is eligible.

## Definition

```yaml
{"operationId": "notification-alert-types-get-alert-types", "summary": "Get Alert Types", "description": "Gets a list of all alert types for which an account is eligible.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"200": {"description": "Get Alert Types response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_alerts-response_collection"}}}}, "4XX": {"description": "Get Alert Types response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification Alert Types"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"]}
```
