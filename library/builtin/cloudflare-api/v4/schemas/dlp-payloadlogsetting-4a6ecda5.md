---
title: dlp_PayloadLogSetting
page_id: schema-dlp-payloadlogsetting-4a6ecda5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_PayloadLogSetting

```yaml
{"type": "object", "properties": {"masking_level": {"allOf": [{"$ref": "#/components/schemas/dlp_PayloadLogMaskingLevel"}]}, "public_key": {"description": "Base64-encoded public key for encrypting payload logs. Null when payload logging is disabled.", "type": "string", "nullable": true}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["updated_at"]}
```
