---
title: OutputComputerCallItem
page_id: schema-outputcomputercallitem-2af21dfb
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputComputerCallItem

```yaml
{"example": {"action": {"type": "screenshot"}, "call_id": "call-abc123", "id": "cu-abc123", "pending_safety_checks": [], "status": "completed", "type": "computer_call"}, "properties": {"action": {}, "call_id": {"type": "string"}, "id": {"type": "string"}, "pending_safety_checks": {"items": {"properties": {"code": {"type": "string"}, "id": {"type": "string"}, "message": {"type": "string"}}, "required": ["id", "code", "message"], "type": "object"}, "type": "array"}, "status": {"enum": ["completed", "incomplete", "in_progress"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["computer_call"], "type": "string"}}, "required": ["type", "call_id", "status", "pending_safety_checks"], "type": "object"}
```
