---
title: realtimekit_stopReason
page_id: schema-realtimekit-stopreason-98f54b8f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_stopReason

```yaml
{"type": "object", "properties": {"caller": {"type": "object", "properties": {"name": {"description": "Name of the user who stopped the recording.", "type": "string", "example": "RealtimeKit_test"}, "type": {"description": "The type can be an App or a user. If the type is `user`, then only the `user_Id` and `name` are returned.", "type": "string", "enum": ["ORGANIZATION", "USER"]}, "user_Id": {"description": "The user ID of the person who stopped the recording.", "type": "string", "format": "uuid", "example": "d61f6956-e68f-4375-bf10-c38a704d1bec"}}}, "reason": {"description": "Specifies the reason why the recording stopped.", "type": "string", "enum": ["API_CALL", "INTERNAL_ERROR", "ALL_PEERS_LEFT"]}}, "title": "stopReason"}
```
