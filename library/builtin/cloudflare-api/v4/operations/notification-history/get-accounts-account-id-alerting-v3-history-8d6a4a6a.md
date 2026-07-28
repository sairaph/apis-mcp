---
title: List History
page_id: operation-get-accounts-account-id-alerting-v3-history-ce3cc44a
path: operations/notification-history
description: Gets a list of history records for notifications sent to an account. The records are displayed for last `x` number of days based on the zone plan (free = 30, pro = 30, biz = 30, ent = 90).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/history
operation_ids:
    - notification-history-list-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List History

`GET /accounts/{account_id}/alerting/v3/history`

Operation ID: `notification-history-list-history`

Gets a list of history records for notifications sent to an account. The records are displayed for last `x` number of days based on the zone plan (free = 30, pro = 30, biz = 30, ent = 90).

## Definition

```yaml
{"operationId": "notification-history-list-history", "summary": "List History", "description": "Gets a list of history records for notifications sent to an account. The records are displayed for last `x` number of days based on the zone plan (free = 30, pro = 30, biz = 30, ent = 90).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/aaa_per_page"}}, {"name": "before", "in": "query", "schema": {"$ref": "#/components/schemas/aaa_before"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "since", "in": "query", "schema": {"description": "Limit the returned results to history records newer than the specified date. This must be a timestamp that conforms to RFC3339.", "type": "string", "format": "date-time", "example": "2022-05-19T20:29:58.679897Z"}}], "responses": {"200": {"description": "List History response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_history_components-schemas-response_collection"}}}}, "4XX": {"description": "List History response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_history_components-schemas-response_collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification History"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"]}
```
