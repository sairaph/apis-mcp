---
title: Upsert Metrics Exports
page_id: operation-post-accounts-account-id-workers-observability-metricsexport-7e873520
path: operations/metrics-export
description: Create or replace resources configured for Workers Observability metrics export.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/observability/metricsexport
operation_ids:
    - metricsExport.upsert
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upsert Metrics Exports

`POST /accounts/{account_id}/workers/observability/metricsexport`

Operation ID: `metricsExport.upsert`

Create or replace resources configured for Workers Observability metrics export.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Your Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "metricsExport.upsert", "summary": "Upsert Metrics Exports", "description": "Create or replace resources configured for Workers Observability metrics export.", "requestBody": {"content": {"application/json": {"schema": {"anyOf": [{"properties": {"destinations": {"type": "array", "items": {"minLength": 1, "type": "string"}, "minItems": 1}, "meta": {"type": "string", "minLength": 1}, "resourceId": {"type": "string", "minLength": 1}, "resourceType": {"type": "string", "minLength": 1}}, "required": ["resourceType", "resourceId", "destinations"], "type": "object"}, {"properties": {"requester": {"type": "object", "properties": {"requesterId": {"type": "string", "minLength": 1}, "requesterType": {"type": "string", "minLength": 1}}, "required": ["requesterType", "requesterId"]}, "resources": {"type": "array", "items": {"properties": {"destinations": {"type": "array", "items": {"minLength": 1, "type": "string"}, "minItems": 1}, "meta": {"type": "string", "minLength": 1}, "resourceId": {"type": "string", "minLength": 1}, "resourceType": {"type": "string", "minLength": 1}}, "required": ["resourceType", "resourceId", "destinations"], "type": "object"}}}, "required": ["requester", "resources"], "type": "object"}, {"properties": {"resources": {"type": "array", "items": {"properties": {"destinations": {"type": "array", "items": {"minLength": 1, "type": "string"}, "minItems": 1}, "meta": {"type": "string", "minLength": 1}, "resourceId": {"type": "string", "minLength": 1}, "resourceType": {"type": "string", "minLength": 1}}, "required": ["resourceType", "resourceId", "destinations"], "type": "object"}, "minItems": 1}}, "required": ["resources"], "type": "object"}]}}}}, "responses": {"201": {"description": "Resource created", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string", "enum": ["Resource created"]}}, "required": ["message"], "type": "object"}}, "result": {"type": "array", "items": {"properties": {"createdAt": {"type": "string"}, "destinations": {"type": "array", "items": {"minLength": 1, "type": "string"}, "minItems": 1}, "meta": {"type": "string", "minLength": 1}, "resourceId": {"type": "string", "minLength": 1}, "resourceType": {"type": "string", "minLength": 1}, "updatedAt": {"type": "string"}}, "required": ["resourceType", "resourceId", "destinations", "createdAt", "updatedAt"], "type": "object"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["messages", "success", "errors", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Bad Request"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Unauthorized"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "409": {"description": "Conflict", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Conflict"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "500": {"description": "Internal error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Internal error"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}}, "tags": ["Metrics Export"], "x-api-token-group": ["Workers Observability Write"]}
```
