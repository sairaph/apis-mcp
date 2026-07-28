---
title: Get Destinations
page_id: operation-get-accounts-account-id-workers-observability-destinations-66e34979
path: operations/destinations
description: List your Workers Observability Telemetry Destinations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/observability/destinations
operation_ids:
    - destination.list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Destinations

`GET /accounts/{account_id}/workers/observability/destinations`

Operation ID: `destination.list`

List your Workers Observability Telemetry Destinations.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Your Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "destination.list", "summary": "Get Destinations", "description": "List your Workers Observability Telemetry Destinations.", "parameters": [{"name": "page", "in": "query", "schema": {"type": "number", "default": 1, "minimum": 1}}, {"name": "perPage", "in": "query", "schema": {"type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"type": "string", "default": "desc", "enum": ["asc", "desc"]}}, {"name": "orderBy", "in": "query", "schema": {"type": "string", "default": "updated", "enum": ["created", "updated"]}}], "responses": {"200": {"description": "Successful request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string", "enum": ["Successful request"]}}, "required": ["message"], "type": "object"}}, "result": {"type": "array", "items": {"properties": {"configuration": {"type": "object", "properties": {"destination_conf": {"type": "string"}, "headers": {"type": "object", "additionalProperties": {"type": "string"}}, "jobStatus": {"type": "object", "properties": {"error_message": {"type": "string"}, "last_complete": {"type": "string"}, "last_error": {"type": "string"}}, "required": ["last_complete", "last_error", "error_message"]}, "logpushDataset": {"anyOf": [{"enum": ["opentelemetry-traces"], "type": "string"}, {"enum": ["opentelemetry-logs"], "type": "string"}, {"enum": ["opentelemetry-metrics"], "type": "string"}]}, "type": {"type": "string", "enum": ["logpush"]}, "url": {"type": "string"}}, "required": ["type", "logpushDataset", "destination_conf", "url", "headers", "jobStatus"]}, "enabled": {"type": "boolean"}, "name": {"type": "string", "pattern": "^[a-z0-9][a-z0-9-]*[a-z0-9]$"}, "scripts": {"type": "array", "items": {"type": "string"}}, "slug": {"type": "string"}}, "required": ["slug", "name", "enabled", "configuration", "scripts"], "type": "object"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["messages", "success", "errors", "result"]}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Unauthorized"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "404": {"description": "Not found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Not found"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "500": {"description": "Internal error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Internal error"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}}, "tags": ["Destinations"], "x-api-token-group": ["Workers Observability Write", "Workers Observability Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "observability.destinations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
