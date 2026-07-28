---
title: Get Event Notification Rule
page_id: operation-get-accounts-account-id-event-notifications-r2-bucket-name-configuration-72249759
path: operations/r2-bucket
description: Get a single event notification rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/event_notifications/r2/{bucket_name}/configuration/queues/{queue_id}
operation_ids:
    - r2-get-event-notification-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Event Notification Rule

`GET /accounts/{account_id}/event_notifications/r2/{bucket_name}/configuration/queues/{queue_id}`

Operation ID: `r2-get-event-notification-config`

Get a single event notification rule.

## Definition

```yaml
{"operationId": "r2-get-event-notification-config", "summary": "Get Event Notification Rule", "description": "Get a single event notification rule.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_queue_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"description": "The bucket jurisdiction.", "type": "string", "default": "default", "enum": ["default", "eu", "fedramp"], "x-stainless-param": "jurisdiction"}}], "responses": {"200": {"description": "Read Configuration response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_queues-config"}}, "type": "object"}]}}}}, "404": {"description": "No Configuration Found response.", "content": {"application/json": {"example": {"errors": [{"code": 11015, "message": "workers.api.error.no_configs_found_for_bucket"}], "messages": [], "result": null, "success": false}, "schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}, "4XX": {"description": "Read Configuration failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write", "Workers R2 Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.event-notifications", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
