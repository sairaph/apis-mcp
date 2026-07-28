---
title: List zone email suppressions
page_id: operation-get-zones-zone-id-email-sending-suppression-b2a078ba
path: operations/email-sending-suppressions
description: Lists email suppressions for the specified zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/sending/suppression
operation_ids:
    - get_publicListSuppressionZoneSending
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List zone email suppressions

`GET /zones/{zone_id}/email/sending/suppression`

Operation ID: `get_publicListSuppressionZoneSending`

Lists email suppressions for the specified zone.

## Definition

```yaml
{"operationId": "get_publicListSuppressionZoneSending", "summary": "List zone email suppressions", "description": "Lists email suppressions for the specified zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"type": "string", "example": "88228822"}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1, "minimum": 0}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 100, "maximum": 1000, "minimum": 0}}, {"name": "order", "in": "query", "schema": {"type": "string", "default": "created_at", "enum": ["email", "expires_at", "created_at"]}}, {"name": "direction", "in": "query", "schema": {"type": "string", "default": "desc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "Returns suppressed emails.", "content": {"application/json": {"schema": {"type": "object", "properties": {"page": {"type": "integer"}, "per_page": {"type": "integer", "example": 100}, "result": {"type": "array", "items": {"properties": {"created_at": {"type": "string", "format": "date-time"}, "email": {"type": "string", "format": "email", "example": "user@example.com"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}, "id": {"type": "string", "format": "uuid"}, "reason": {"type": "string", "example": "hard_bounce"}, "zones": {"type": "array", "items": {"type": "string"}, "default": []}}, "required": ["id", "email", "reason", "created_at", "expires_at"], "type": "object"}}, "total": {"type": "integer", "example": 1}}, "required": ["page", "per_page", "total", "result"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "string"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
