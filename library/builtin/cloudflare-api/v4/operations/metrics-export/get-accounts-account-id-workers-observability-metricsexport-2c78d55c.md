---
title: List Metrics Exports
page_id: operation-get-accounts-account-id-workers-observability-metricsexport-79f6bfcc
path: operations/metrics-export
description: List resources configured for Workers Observability metrics export.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/observability/metricsexport
operation_ids:
    - metricsExport.list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Metrics Exports

`GET /accounts/{account_id}/workers/observability/metricsexport`

Operation ID: `metricsExport.list`

List resources configured for Workers Observability metrics export.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Your Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "metricsExport.list", "summary": "List Metrics Exports", "description": "List resources configured for Workers Observability metrics export.", "responses": {"200": {"description": "Successful request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string", "enum": ["Successful request"]}}, "required": ["message"], "type": "object"}}, "result": {"type": "array", "items": {"properties": {"createdAt": {"type": "string"}, "destinations": {"type": "array", "items": {"minLength": 1, "type": "string"}, "minItems": 1}, "meta": {"type": "string", "minLength": 1}, "resourceId": {"type": "string", "minLength": 1}, "resourceType": {"type": "string", "minLength": 1}, "updatedAt": {"type": "string"}}, "required": ["resourceType", "resourceId", "destinations", "createdAt", "updatedAt"], "type": "object"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["messages", "success", "errors", "result"]}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Unauthorized"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "500": {"description": "Internal error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Internal error"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}}, "tags": ["Metrics Export"], "x-api-token-group": ["Workers Observability Write", "Workers Observability Read"]}
```
