---
title: access_jit_request_event
page_id: schema-access-jit-request-event-f0af0603
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_jit_request_event

```yaml
{"type": "object", "properties": {"actor_email": {"type": "string", "format": "email"}, "actor_idp": {"type": "string"}, "actor_uuid": {"type": "string", "format": "uuid"}, "decision_outcome": {"type": "string"}, "ray_id": {"type": "string"}, "session_duration_seconds": {"type": "integer"}, "success": {"type": "boolean"}, "timestamp": {"$ref": "#/components/schemas/access_timestamp"}, "type": {"$ref": "#/components/schemas/access_jit_request_event_type"}}}
```
