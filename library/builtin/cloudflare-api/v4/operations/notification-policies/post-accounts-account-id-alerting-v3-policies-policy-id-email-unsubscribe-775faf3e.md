---
title: Unsubscribe email from a Notification policy
page_id: operation-post-accounts-account-id-alerting-v3-policies-policy-id-email-unsubscrib-690550d2
path: operations/notification-policies
description: Unsubscribes an email address from a notification policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/alerting/v3/policies/{policy_id}/email/unsubscribe
operation_ids:
    - notification-policies-unsubscribe-email-from-notification-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Unsubscribe email from a Notification policy

`POST /accounts/{account_id}/alerting/v3/policies/{policy_id}/email/unsubscribe`

Operation ID: `notification-policies-unsubscribe-email-from-notification-policy`

Unsubscribes an email address from a notification policy.

## Definition

```yaml
{"operationId": "notification-policies-unsubscribe-email-from-notification-policy", "summary": "Unsubscribe email from a Notification policy", "description": "Unsubscribes an email address from a notification policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_policy-id"}}, {"name": "email", "in": "query", "required": true, "schema": {"type": "string", "format": "email"}}, {"name": "token", "in": "query", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Unsubscribe email from Notification policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_unsubscribe_email_post_single_response"}}, "text/html": {"schema": {"type": "string"}}}}, "4XX": {"description": "Unsubscribe email from Notification policy response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}}}}}, "security": [], "tags": ["Notification policies"]}
```
