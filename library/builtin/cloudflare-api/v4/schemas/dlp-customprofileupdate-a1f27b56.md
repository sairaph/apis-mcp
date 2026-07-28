---
title: dlp_CustomProfileUpdate
page_id: schema-dlp-customprofileupdate-a1f27b56
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_CustomProfileUpdate

```yaml
{"type": "object", "properties": {"ai_context_enabled": {"type": "boolean", "default": false}, "allowed_match_count": {"type": "integer", "format": "int32", "nullable": true}, "confidence_threshold": {"type": "string", "default": "low", "nullable": true}, "context_awareness": {"$ref": "#/components/schemas/dlp_ContextAwareness"}, "data_classes": {"description": "Data class IDs to associate with the profile. If omitted, existing associations are unchanged.", "type": "array", "items": {"format": "uuid", "type": "string"}, "nullable": true}, "data_tags": {"description": "Data tag IDs to associate with the profile. If omitted, existing associations are unchanged.", "type": "array", "items": {"format": "uuid", "type": "string"}, "nullable": true}, "description": {"description": "The description of the profile.", "type": "string", "nullable": true}, "entries": {"description": "Custom entries from this profile.\nIf this field is omitted, entries owned by this profile will not be changed.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_ProfileEntryUpdate"}, "deprecated": true, "nullable": true}, "name": {"type": "string"}, "ocr_enabled": {"type": "boolean", "default": false}, "sensitivity_levels": {"description": "Sensitivity levels to associate with the profile. If omitted, existing associations are unchanged.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_SensitivityLevelRef"}, "nullable": true}, "shared_entries": {"description": "Other entries, e.g. predefined or integration.", "type": "array", "items": {"$ref": "#/components/schemas/dlp_SharedEntryUpdate"}}}, "required": ["name"]}
```
