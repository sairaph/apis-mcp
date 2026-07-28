---
title: Show email unsubscribe details
page_id: operation-get-accounts-account-id-alerting-v3-policies-policy-id-email-unsubscribe-2fa9a7c1
path: operations/notification-policies
description: Shows details for unsubscribing an email address from a notification policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/policies/{policy_id}/email/unsubscribe
operation_ids:
    - notification-policies-show-email-unsubscribe-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Show email unsubscribe details

`GET /accounts/{account_id}/alerting/v3/policies/{policy_id}/email/unsubscribe`

Operation ID: `notification-policies-show-email-unsubscribe-details`

Shows details for unsubscribing an email address from a notification policy.

## Definition

```yaml
{"operationId": "notification-policies-show-email-unsubscribe-details", "summary": "Show email unsubscribe details", "description": "Shows details for unsubscribing an email address from a notification policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_policy-id"}}, {"name": "email", "in": "query", "required": true, "schema": {"type": "string", "format": "email"}}, {"name": "token", "in": "query", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Show email unsubscribe details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_unsubscribe_email_single_response"}}, "text/html": {"schema": {"type": "string"}}}}, "4XX": {"description": "Show email unsubscribe details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}}}}}, "security": [], "tags": ["Notification policies"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"]}
```
