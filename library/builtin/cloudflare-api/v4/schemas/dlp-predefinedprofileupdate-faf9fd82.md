---
title: dlp_PredefinedProfileUpdate
page_id: schema-dlp-predefinedprofileupdate-faf9fd82
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PredefinedProfileUpdate

```yaml
{"type": "object", "properties": {"ai_context_enabled": {"type": "boolean", "default": false}, "allowed_match_count": {"type": "integer", "format": "int32", "example": 5, "default": 0, "maximum": 1000, "minimum": 0, "nullable": true}, "confidence_threshold": {"type": "string", "default": "low", "nullable": true}, "context_awareness": {"$ref": "#/components/schemas/dlp_ContextAwareness"}, "entries": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_PredefinedProfileEntryUpdate"}, "deprecated": true}, "ocr_enabled": {"type": "boolean", "default": false}}}
```
