---
title: Create zone email suppression
page_id: operation-post-zones-zone-id-email-routing-suppression-d3b7312a
path: operations/email-routing-suppressions
description: Creates a new email suppression for the specified zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/email/routing/suppression
operation_ids:
    - post_publicNewSuppressionZoneRouting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create zone email suppression

`POST /zones/{zone_id}/email/routing/suppression`

Operation ID: `post_publicNewSuppressionZoneRouting`

Creates a new email suppression for the specified zone.

## Definition

```yaml
{"operationId": "post_publicNewSuppressionZoneRouting", "summary": "Create zone email suppression", "description": "Creates a new email suppression for the specified zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"type": "string", "example": "88228822"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"email": {"type": "string", "format": "email", "example": "user@example.com"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}}, "required": ["email"]}}}}, "responses": {"200": {"description": "Returns suppression.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"id": {"type": "string", "format": "uuid"}}, "required": ["id"]}}, "required": ["result"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "string"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Routing suppressions"]}
```
