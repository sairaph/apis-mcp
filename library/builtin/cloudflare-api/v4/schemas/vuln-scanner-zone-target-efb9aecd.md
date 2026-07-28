---
title: vuln_scanner_zone-target
page_id: schema-vuln-scanner-zone-target-efb9aecd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_zone-target

```yaml
{"type": "object", "properties": {"type": {"type": "string", "enum": ["zone"]}, "zone_tag": {"description": "Cloudflare zone tag. The zone must belong to the account.\n", "type": "string", "example": "d8e8fca2dc0f896fd7cb4cb0031ba249", "maxLength": 32}}, "required": ["type", "zone_tag"]}
```
