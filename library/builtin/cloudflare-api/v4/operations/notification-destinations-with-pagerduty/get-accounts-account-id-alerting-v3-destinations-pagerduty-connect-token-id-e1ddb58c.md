---
title: Connect PagerDuty
page_id: operation-get-accounts-account-id-alerting-v3-destinations-pagerduty-connect-token-79643132
path: operations/notification-destinations-with-pagerduty
description: Links PagerDuty with the account using the integration token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/pagerduty/connect/{token_id}
operation_ids:
    - notification-destinations-with-pager-duty-connect-pager-duty-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Connect PagerDuty

`GET /accounts/{account_id}/alerting/v3/destinations/pagerduty/connect/{token_id}`

Operation ID: `notification-destinations-with-pager-duty-connect-pager-duty-token`

Links PagerDuty with the account using the integration token.

## Definition

```yaml
{"operationId": "notification-destinations-with-pager-duty-connect-pager-duty-token", "summary": "Connect PagerDuty", "description": "Links PagerDuty with the account using the integration token.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_integration-token"}}], "responses": {"200": {"description": "Create a Notification policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_id_response"}}}}, "4XX": {"description": "Create a Notification policy response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_id_response"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification destinations with PagerDuty"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
