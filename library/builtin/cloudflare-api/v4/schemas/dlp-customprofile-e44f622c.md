---
title: dlp_CustomProfile
page_id: schema-dlp-customprofile-e44f622c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_CustomProfile

```yaml
{"type": "object", "properties": {"ai_context_enabled": {"type": "boolean", "default": false}, "allowed_match_count": {"description": "Related DLP policies will trigger when the match count exceeds the number set.", "type": "integer", "format": "int32", "example": 5, "default": 0, "maximum": 1000, "minimum": 0}, "confidence_threshold": {"default": "low", "allOf": [{"$ref": "#/components/schemas/dlp_Confidence"}]}, "context_awareness": {"$ref": "#/components/schemas/dlp_ContextAwareness"}, "created_at": {"description": "When the profile was created.", "type": "string", "format": "date-time"}, "data_classes": {"description": "Data classes associated with this profile.", "type": "array", "items": {"format": "uuid", "type": "string"}}, "data_tags": {"description": "Data tags associated with this profile.", "type": "array", "items": {"format": "uuid", "type": "string"}}, "description": {"description": "The description of the profile.", "type": "string", "nullable": true}, "entries": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_Entry"}, "deprecated": true, "x-stainless-skip": ["terraform"]}, "id": {"description": "The id of the profile (uuid).", "type": "string", "format": "uuid"}, "name": {"description": "The name of the profile.", "type": "string", "x-auditable": true}, "ocr_enabled": {"type": "boolean", "default": false}, "sensitivity_levels": {"description": "Sensitivity levels associated with this profile.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_SensitivityLevelRef"}}, "shared_entries": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_Entry"}, "x-stainless-terraform-configurability": "computed"}, "updated_at": {"description": "When the profile was lasted updated.", "type": "string", "format": "date-time"}}, "required": ["id", "name", "created_at", "updated_at", "allowed_match_count", "ocr_enabled"]}
```
