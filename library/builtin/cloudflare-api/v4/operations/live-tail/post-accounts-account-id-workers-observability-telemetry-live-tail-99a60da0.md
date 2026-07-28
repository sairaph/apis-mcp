---
title: Prepare live tail
page_id: operation-post-accounts-account-id-workers-observability-telemetry-live-tail-3f846784
path: operations/live-tail
description: Prepare websocket server for live tail.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/observability/telemetry/live-tail
operation_ids:
    - telemetry.live-tail.post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Prepare live tail

`POST /accounts/{account_id}/workers/observability/telemetry/live-tail`

Operation ID: `telemetry.live-tail.post`

Prepare websocket server for live tail.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Your Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "telemetry.live-tail.post", "summary": "Prepare live tail", "description": "Prepare websocket server for live tail.", "requestBody": {"description": "Create websocket server for live tail.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"filterCombination": {"description": "Set a flag to describe how to combine the filters on the query.", "type": "string", "default": "and", "enum": ["and", "or", "AND", "OR"]}, "filters": {"description": "Apply filters to the query. Supports nested groups via kind: 'group'.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/workers-observability_filter_node"}, {"anyOf": [{"description": "A group node that nests child filters combined with a logical operator.", "properties": {"filterCombination": {"description": "Logical operator for combining child filters: 'and' (all must match) or 'or' (any must match).", "type": "string", "enum": ["and", "or", "AND", "OR"]}, "filters": {"description": "Child filter nodes. Each can be a leaf filter or another nested group.", "type": "array", "items": {"$ref": "#/components/schemas/workers-observability_filter_node"}, "minItems": 1}, "kind": {"description": "Discriminator indicating this is a nested filter group.", "type": "string", "enum": ["group"]}}, "required": ["kind", "filterCombination", "filters"], "type": "object"}, {"$ref": "#/components/schemas/workers-observability_filter_leaf"}], "description": "Supports nested groups via kind: 'group'. Maximum nesting depth is 4."}]}, "default": []}, "scriptId": {"type": "string"}}}}}}, "responses": {"200": {"description": "Successful request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string", "enum": ["Successful request"]}}, "required": ["message"], "type": "object"}}, "result": {"type": "object", "properties": {"wsUrl": {"description": "WebSocket URL clients connect to in order to stream live tail events.", "type": "string", "format": "uri"}}, "required": ["wsUrl"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["messages", "success", "errors", "result"]}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Unauthorized"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Forbidden"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "500": {"description": "Internal error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Internal error"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}}, "tags": ["Live Tail"], "x-api-token-group": ["Workers Observability Write"]}
```
