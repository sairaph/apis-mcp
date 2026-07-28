---
title: dlp_PredefinedProfile
page_id: schema-dlp-predefinedprofile-756a9a17
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PredefinedProfile

```yaml
{"type": "object", "properties": {"ai_context_enabled": {"type": "boolean", "default": false}, "allowed_match_count": {"type": "integer", "format": "int32"}, "confidence_threshold": {"default": "low", "allOf": [{"$ref": "#/components/schemas/dlp_Confidence"}]}, "context_awareness": {"$ref": "#/components/schemas/dlp_ContextAwareness"}, "entries": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_Entry"}, "deprecated": true}, "id": {"description": "The id of the predefined profile (uuid).", "type": "string", "format": "uuid"}, "name": {"description": "The name of the predefined profile.", "type": "string"}, "ocr_enabled": {"type": "boolean", "default": false}, "open_access": {"description": "Whether this profile can be accessed by anyone.", "type": "boolean"}}, "required": ["id", "name", "entries", "allowed_match_count"]}
```
