---
title: Create Destination
page_id: operation-post-accounts-account-id-workers-observability-destinations-2e62def1
path: operations/destinations
description: Create a new Workers Observability Telemetry Destination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/observability/destinations
operation_ids:
    - destination.create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Destination

`POST /accounts/{account_id}/workers/observability/destinations`

Operation ID: `destination.create`

Create a new Workers Observability Telemetry Destination.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Your Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "destination.create", "summary": "Create Destination", "description": "Create a new Workers Observability Telemetry Destination.", "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"configuration": {"type": "object", "properties": {"headers": {"type": "object", "additionalProperties": {"type": "string"}}, "logpushDataset": {"anyOf": [{"enum": ["opentelemetry-traces"], "type": "string"}, {"enum": ["opentelemetry-logs"], "type": "string"}, {"enum": ["opentelemetry-metrics"], "type": "string"}]}, "type": {"type": "string", "enum": ["logpush"]}, "url": {"type": "string"}}, "required": ["type", "logpushDataset", "url", "headers"]}, "enabled": {"type": "boolean"}, "name": {"type": "string", "pattern": "^[a-z0-9][a-z0-9-]*[a-z0-9]$"}, "skipPreflightCheck": {"type": "boolean"}}, "required": ["name", "enabled", "configuration"]}}}}, "responses": {"201": {"description": "Resource created", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string", "enum": ["Resource created"]}}, "required": ["message"], "type": "object"}}, "result": {"type": "object", "properties": {"configuration": {"type": "object", "properties": {"destination_conf": {"type": "string"}, "logpushDataset": {"anyOf": [{"enum": ["opentelemetry-traces"], "type": "string"}, {"enum": ["opentelemetry-logs"], "type": "string"}, {"enum": ["opentelemetry-metrics"], "type": "string"}]}, "logpushJob": {"type": "number"}, "type": {"type": "string", "enum": ["logpush"]}, "url": {"type": "string"}}, "required": ["type", "logpushDataset", "logpushJob", "destination_conf", "url"]}, "enabled": {"type": "boolean"}, "name": {"type": "string", "pattern": "^[a-z0-9][a-z0-9-]*[a-z0-9]$"}, "scripts": {"type": "array", "items": {"type": "string"}}, "slug": {"type": "string"}}, "required": ["slug", "name", "enabled", "configuration", "scripts"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["messages", "success", "errors", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Bad Request"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Unauthorized"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "403": {"description": "Forbidden", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Forbidden"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "500": {"description": "Internal error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Internal error"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}}, "tags": ["Destinations"], "x-api-token-group": ["Workers Observability Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "observability.destinations", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
