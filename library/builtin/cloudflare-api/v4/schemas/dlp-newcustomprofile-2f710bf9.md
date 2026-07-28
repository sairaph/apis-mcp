---
title: dlp_NewCustomProfile
page_id: schema-dlp-newcustomprofile-2f710bf9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_NewCustomProfile

```yaml
{"type": "object", "properties": {"ai_context_enabled": {"type": "boolean", "default": false}, "allowed_match_count": {"description": "Related DLP policies will trigger when the match count exceeds the number set.", "type": "integer", "format": "int32", "example": 5, "default": 0, "maximum": 1000, "minimum": 0}, "confidence_threshold": {"type": "string", "default": "low", "nullable": true}, "context_awareness": {"$ref": "#/components/schemas/dlp_ContextAwareness"}, "data_classes": {"description": "Data class IDs to associate with the profile.", "type": "array", "items": {"format": "uuid", "type": "string"}}, "data_tags": {"description": "Data tag IDs to associate with the profile.", "type": "array", "items": {"format": "uuid", "type": "string"}}, "description": {"description": "The description of the profile.", "type": "string", "nullable": true}, "entries": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_EntryOfNewProfile"}, "x-stainless-skip": ["terraform"]}, "name": {"type": "string"}, "ocr_enabled": {"type": "boolean", "default": false}, "sensitivity_levels": {"description": "Sensitivity levels to associate with the profile.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_SensitivityLevelRef"}}, "shared_entries": {"description": "Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your Microsoft Information Protection profiles).", "type": "array", "items": {"$ref": "#/components/schemas/dlp_NewSharedEntry"}}}, "required": ["name"]}
```
