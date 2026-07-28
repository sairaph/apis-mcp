---
title: dlp_PredefinedProfileConfigUpdate
page_id: schema-dlp-predefinedprofileconfigupdate-dfc2f5ef
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PredefinedProfileConfigUpdate

```yaml
{"type": "object", "properties": {"ai_context_enabled": {"type": "boolean", "default": false}, "allowed_match_count": {"type": "integer", "format": "int32", "example": 5, "default": 0, "maximum": 1000, "minimum": 0, "nullable": true}, "confidence_threshold": {"type": "string", "default": "low", "nullable": true}, "enabled_entries": {"type": "array", "items": {"format": "uuid", "type": "string"}, "nullable": true}, "entries": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_PredefinedProfileEntryUpdate"}, "deprecated": true, "x-stainless-terraform-configurability": "computed_optional"}, "ocr_enabled": {"type": "boolean", "default": false}}}
```
