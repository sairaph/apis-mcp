---
title: Create PagerDuty integration token
page_id: operation-post-accounts-account-id-alerting-v3-destinations-pagerduty-connect-5a367dd1
path: operations/notification-destinations-with-pagerduty
description: Creates a new token for integrating with PagerDuty.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/pagerduty/connect
operation_ids:
    - notification-destinations-with-pager-duty-connect-pager-duty
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create PagerDuty integration token

`POST /accounts/{account_id}/alerting/v3/destinations/pagerduty/connect`

Operation ID: `notification-destinations-with-pager-duty-connect-pager-duty`

Creates a new token for integrating with PagerDuty.

## Definition

```yaml
{"operationId": "notification-destinations-with-pager-duty-connect-pager-duty", "summary": "Create PagerDuty integration token", "description": "Creates a new token for integrating with PagerDuty.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"201": {"description": "Token for PagerDuty integration", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_sensitive_id_response"}}}}, "4XX": {"description": "Create a token for PagerDuty integration failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_id_response"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification destinations with PagerDuty"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
