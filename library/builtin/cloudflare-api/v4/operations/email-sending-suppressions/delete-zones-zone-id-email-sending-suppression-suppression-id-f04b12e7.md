---
title: Delete zone email suppression
page_id: operation-delete-zones-zone-id-email-sending-suppression-suppression-id-62f2b38e
path: operations/email-sending-suppressions
description: Deletes an email suppression for the specified zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/email/sending/suppression/{suppression_id}
operation_ids:
    - delete_publicDeleteSuppressionZoneSending
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete zone email suppression

`DELETE /zones/{zone_id}/email/sending/suppression/{suppression_id}`

Operation ID: `delete_publicDeleteSuppressionZoneSending`

Deletes an email suppression for the specified zone.

## Definition

```yaml
{"operationId": "delete_publicDeleteSuppressionZoneSending", "summary": "Delete zone email suppression", "description": "Deletes an email suppression for the specified zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"type": "string", "example": "88228822"}}, {"name": "suppression_id", "in": "path", "required": true, "schema": {"type": "string", "example": "396a5436-d4b0-42a6-b3fc-48e8fa522321"}}], "responses": {"200": {"description": "Deletes suppression.", "content": {"application/json": {"schema": {"type": "object", "properties": {"success": {"type": "boolean"}}, "required": ["success"]}}}}, "404": {"description": "Suppression not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "error"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
