---
title: Create Event Notification Rule
page_id: operation-put-accounts-account-id-event-notifications-r2-bucket-name-configuration-e6cbe8e9
path: operations/r2-bucket
description: Create event notification rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/event_notifications/r2/{bucket_name}/configuration/queues/{queue_id}
operation_ids:
    - r2-put-event-notification-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Event Notification Rule

`PUT /accounts/{account_id}/event_notifications/r2/{bucket_name}/configuration/queues/{queue_id}`

Operation ID: `r2-put-event-notification-config`

Create event notification rule.

## Definition

```yaml
{"operationId": "r2-put-event-notification-config", "summary": "Create Event Notification Rule", "description": "Create event notification rule.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_queue_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"rules": {"description": "Array of rules to drive notifications.", "type": "array", "items": {"$ref": "#/components/schemas/r2_rule"}, "minItems": 1}}, "required": ["rules"]}}}}, "responses": {"200": {"description": "Create Configuration response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"type": "object"}]}}}}, "4XX": {"description": "Create Configuration failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.event-notifications", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
