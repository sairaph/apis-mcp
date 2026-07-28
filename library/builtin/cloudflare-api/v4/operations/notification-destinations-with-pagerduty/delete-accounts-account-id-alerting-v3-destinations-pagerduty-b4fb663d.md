---
title: Delete PagerDuty Services
page_id: operation-delete-accounts-account-id-alerting-v3-destinations-pagerduty-48e1057b
path: operations/notification-destinations-with-pagerduty
description: Deletes all the PagerDuty Services connected to the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/pagerduty
operation_ids:
    - notification-destinations-with-pager-duty-delete-pager-duty-services
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete PagerDuty Services

`DELETE /accounts/{account_id}/alerting/v3/destinations/pagerduty`

Operation ID: `notification-destinations-with-pager-duty-delete-pager-duty-services`

Deletes all the PagerDuty Services connected to the account.

## Definition

```yaml
{"operationId": "notification-destinations-with-pager-duty-delete-pager-duty-services", "summary": "Delete PagerDuty Services", "description": "Deletes all the PagerDuty Services connected to the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"200": {"description": "Delete PagerDuty Services response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-2"}}}}, "4XX": {"description": "Delete PagerDuty Services response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification destinations with PagerDuty"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
