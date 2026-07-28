---
title: Acknowledge + Retry Queue Messages
page_id: operation-post-accounts-account-id-queues-queue-id-messages-ack-8eb8295f
path: operations/queue
description: Acknowledge + Retry messages from a Queue
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/messages/ack
operation_ids:
    - queues-ack-messages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Acknowledge + Retry Queue Messages

`POST /accounts/{account_id}/queues/{queue_id}/messages/ack`

Operation ID: `queues-ack-messages`

Acknowledge + Retry messages from a Queue

## Definition

```yaml
{"operationId": "queues-ack-messages", "summary": "Acknowledge + Retry Queue Messages", "description": "Acknowledge + Retry messages from a Queue", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"acks": {"type": "array", "items": {"properties": {"lease_id": {"$ref": "#/components/schemas/mq_lease-id"}}, "type": "object"}}, "retries": {"type": "array", "items": {"properties": {"delay_seconds": {"$ref": "#/components/schemas/mq_retry-delay"}, "lease_id": {"$ref": "#/components/schemas/mq_lease-id"}}, "type": "object"}}}}}}}, "responses": {"200": {"description": "Details of ACKs and retries", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "object", "properties": {"ackCount": {"description": "The number of messages that were succesfully acknowledged.", "type": "number", "example": 5}, "retryCount": {"description": "The number of messages that were succesfully retried.", "type": "number", "example": 5}, "warnings": {"description": "Map of lease IDs to warning messages encountered during acknowledgement.", "type": "object", "additionalProperties": {"type": "string"}}}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-auditable": false, "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.ack"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.messages", "x-fern-sdk-method-name": "ack", "x-forge-hidden": true}
```
