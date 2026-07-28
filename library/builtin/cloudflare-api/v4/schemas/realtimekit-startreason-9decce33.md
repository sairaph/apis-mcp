---
title: realtimekit_startReason
page_id: schema-realtimekit-startreason-9decce33
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_startReason

```yaml
{"type": "object", "properties": {"caller": {"type": "object", "properties": {"name": {"description": "Name of the user who started the recording.", "type": "string", "example": "RealtimeKit_test"}, "type": {"description": "The type can be an App or a user. If the type is `user`, then only the `user_Id` and `name` are returned.", "type": "string", "enum": ["ORGANIZATION", "USER"]}, "user_Id": {"description": "The user ID of the person who started the recording.", "type": "string", "format": "uuid", "example": "d61f6956-e68f-4375-bf10-c38a704d1bec"}}}, "reason": {"description": "Specifies if the recording was started using the \"Start a Recording\"API or using the parameter RECORD_ON_START in the \"Create a meeting\" API. \n\nIf the recording is initiated using the \"RECORD_ON_START\" parameter, the user details will not be populated.", "type": "string", "enum": ["API_CALL", "RECORD_ON_START"]}}, "title": "startReason"}
```
