---
title: Delete Event Notification Rules
page_id: operation-delete-accounts-account-id-event-notifications-r2-bucket-name-configurat-c37ff964
path: operations/r2-bucket
description: Delete an event notification rule. **If no body is provided, all rules for specified queue will be deleted**.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/event_notifications/r2/{bucket_name}/configuration/queues/{queue_id}
operation_ids:
    - r2-event-notification-delete-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Event Notification Rules

`DELETE /accounts/{account_id}/event_notifications/r2/{bucket_name}/configuration/queues/{queue_id}`

Operation ID: `r2-event-notification-delete-config`

Delete an event notification rule. **If no body is provided, all rules for specified queue will be deleted**.

## Definition

```yaml
{"operationId": "r2-event-notification-delete-config", "summary": "Delete Event Notification Rules", "description": "Delete an event notification rule. **If no body is provided, all rules for specified queue will be deleted**.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_queue_identifier"}}, {"name": "bucket_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_bucket_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}, {"name": "cf-r2-jurisdiction", "in": "header", "schema": {"$ref": "#/components/schemas/r2_jurisdiction"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"ruleIds": {"description": "Array of rule ids to delete.", "type": "array", "items": {"description": "rule ids to be deleted.", "type": "string"}}}}}}}, "responses": {"200": {"description": "Delete Configuration response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"type": "object"}]}}}}, "4XX": {"description": "Delete Configuration failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Bucket"], "x-api-token-group": ["Workers R2 Storage Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.event-notifications", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
