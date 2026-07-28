---
title: email-security_MoveResponseItem
page_id: schema-email-security-moveresponseitem-f254018a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_MoveResponseItem

```yaml
{"type": "object", "properties": {"completed_at": {"description": "When the move operation completed (UTC).", "type": "string", "format": "date-time", "nullable": true}, "completed_timestamp": {"description": "Deprecated, use `completed_at` instead. End of life: November 1, 2026.", "type": "string", "format": "date-time", "deprecated": true, "x-stainless-deprecation-message": "Use `completed_at` instead."}, "destination": {"description": "Destination folder for the message.", "type": "string", "nullable": true}, "item_count": {"description": "Number of items moved. End of life: November 1, 2026.", "type": "integer", "deprecated": true, "x-stainless-deprecation-message": "This field is deprecated."}, "message_id": {"description": "Message identifier.", "type": "string", "nullable": true}, "operation": {"description": "Type of operation performed.", "type": "string", "nullable": true}, "recipient": {"description": "Recipient email address.", "type": "string", "nullable": true}, "status": {"description": "Operation status.", "type": "string", "nullable": true}, "success": {"description": "Whether the operation succeeded.", "type": "boolean"}}, "required": ["success"]}
```
