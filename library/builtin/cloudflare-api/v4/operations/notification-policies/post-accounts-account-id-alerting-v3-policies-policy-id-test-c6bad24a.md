---
title: Test a Notification policy
page_id: operation-post-accounts-account-id-alerting-v3-policies-policy-id-test-e448f5f0
path: operations/notification-policies
description: Send a test notification for a policy to verify delivery mechanisms are working as expected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/alerting/v3/policies/{policy_id}/test
operation_ids:
    - notification-policies-test-a-notification-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Test a Notification policy

`POST /accounts/{account_id}/alerting/v3/policies/{policy_id}/test`

Operation ID: `notification-policies-test-a-notification-policy`

Send a test notification for a policy to verify delivery mechanisms are working as expected.

## Definition

```yaml
{"operationId": "notification-policies-test-a-notification-policy", "summary": "Test a Notification policy", "description": "Send a test notification for a policy to verify delivery mechanisms are working as expected.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_policy-id"}}], "requestBody": {"description": "Optional configuration for the test notification. When omitted, a default INFO-severity test alert is sent.", "content": {"application/json": {"schema": {"type": "object", "properties": {"severity": {"description": "Severity level for the test alert. Defaults to INFO (1) if omitted.", "type": "integer", "example": 1, "enum": [0, 1, 2, 3, 4]}, "source": {"description": "Source identifier for the test alert.", "type": "string"}, "state_correlation_id": {"description": "Correlation ID for stateful test alerts. Required when state_event is set.", "type": "string"}, "state_event": {"description": "State event type for stateful test alerts. Use with state_correlation_id.", "type": "integer", "enum": [0, 1, 2]}}}}}}, "responses": {"200": {"description": "Test a Notification policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-2"}}}}, "4XX": {"description": "Test a Notification policy response failure. Common causes include an invalid or non-existent policy ID, or a delivery mechanism that is unreachable.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification policies"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
