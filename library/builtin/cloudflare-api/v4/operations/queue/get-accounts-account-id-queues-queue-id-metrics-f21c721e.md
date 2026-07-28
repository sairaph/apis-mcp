---
title: Get Queue Metrics
page_id: operation-get-accounts-account-id-queues-queue-id-metrics-8b5ba349
path: operations/queue
description: Return best-effort metrics for a queue. Values may be approximate due to the distributed nature of queues.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/metrics
operation_ids:
    - queues-get-metrics
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Queue Metrics

`GET /accounts/{account_id}/queues/{queue_id}/metrics`

Operation ID: `queues-get-metrics`

Return best-effort metrics for a queue. Values may be approximate due to the distributed nature of queues.

## Definition

```yaml
{"operationId": "queues-get-metrics", "summary": "Get Queue Metrics", "description": "Return best-effort metrics for a queue. Values may be approximate due to the distributed nature of queues.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "Queue metrics.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_queue-metrics"}}, "type": "object"}]}}, "text/event-stream": {"schema": {"$ref": "#/components/schemas/mq_queue-metrics"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Queues Read", "Workers Scripts Write", "Workers Scripts Read"]}
```
