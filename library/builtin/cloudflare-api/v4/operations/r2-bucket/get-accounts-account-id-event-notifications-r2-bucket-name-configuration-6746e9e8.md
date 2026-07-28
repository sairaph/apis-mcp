---
title: List Event Notification Rules
page_id: operation-get-accounts-account-id-event-notifications-r2-bucket-name-configuration-cb2b92ef
path: operations/r2-bucket
description: List all event notification rules for a bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/event_notifications/r2/{bucket_name}/configuration
operation_ids:
    - r2-get-event-notification-configs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Event Notification Rules

`GET /accounts/{account_id}/event_notifications/r2/{bucket_name}/configuration`

Operation ID: `r2-get-event-notification-configs`

List all event notification rules for a bucket.

## Definition

```yaml
{"operationId": "r2-get-event-notification-configs", "summary": "List Event Notification Rules", "description": "List all event notification rules for a bucket.", "parameters": [{"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "responses": {"200": {"description": "Read Configuration response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_bucket-config"}}, "type": "object"}]}}}}, "404": {"description": "No Configuration Found response.", "content": {"application/json": {"example": {"errors": [{"code": 11015, "message": "workers.api.error.no_configs_found_for_bucket"}], "messages": [], "result": null, "success": false}, "schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}, "4XX": {"description": "Read Configuration failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write", "Workers R2 Storage Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.event-notifications", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
