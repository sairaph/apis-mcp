---
title: Get zone email suppression
page_id: operation-get-zones-zone-id-email-sending-suppression-suppression-id-5741ffc7
path: operations/email-sending-suppressions
description: Retrieves a single email suppression for the specified zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/email/sending/suppression/{suppression_id}
operation_ids:
    - get_publicGetSuppressionZoneSending
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get zone email suppression

`GET /zones/{zone_id}/email/sending/suppression/{suppression_id}`

Operation ID: `get_publicGetSuppressionZoneSending`

Retrieves a single email suppression for the specified zone.

## Definition

```yaml
{"operationId": "get_publicGetSuppressionZoneSending", "summary": "Get zone email suppression", "description": "Retrieves a single email suppression for the specified zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"type": "string", "example": "88228822"}}, {"name": "suppression_id", "in": "path", "required": true, "schema": {"type": "string", "example": "396a5436-d4b0-42a6-b3fc-48e8fa522321"}}], "responses": {"200": {"description": "Returns suppression.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "email": {"type": "string", "format": "email", "example": "user@example.com"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}, "id": {"type": "string", "format": "uuid"}, "reason": {"type": "string", "example": "hard_bounce"}, "zones": {"type": "array", "items": {"type": "string"}, "default": []}}, "required": ["id", "email", "reason", "created_at", "expires_at"]}}, "required": ["result"]}}}}, "404": {"description": "Suppression not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "error"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
