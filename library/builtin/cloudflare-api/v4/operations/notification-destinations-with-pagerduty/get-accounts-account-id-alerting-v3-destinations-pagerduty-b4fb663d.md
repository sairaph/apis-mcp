---
title: List PagerDuty services
page_id: operation-get-accounts-account-id-alerting-v3-destinations-pagerduty-0768af76
path: operations/notification-destinations-with-pagerduty
description: Get a list of all configured PagerDuty services.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/pagerduty
operation_ids:
    - notification-destinations-with-pager-duty-list-pager-duty-services
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List PagerDuty services

`GET /accounts/{account_id}/alerting/v3/destinations/pagerduty`

Operation ID: `notification-destinations-with-pager-duty-list-pager-duty-services`

Get a list of all configured PagerDuty services.

## Definition

```yaml
{"operationId": "notification-destinations-with-pager-duty-list-pager-duty-services", "summary": "List PagerDuty services", "description": "Get a list of all configured PagerDuty services.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"200": {"description": "List PagerDuty services response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_components-schemas-response_collection"}}}}, "4XX": {"description": "List PagerDuty services response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification destinations with PagerDuty"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
