---
title: zones_security_level-2
page_id: schema-zones-security-level-2-adb21a5f
path: schemas
description: Choose the appropriate security profile for your website, which will automatically adjust each of the security settings. If you choose to customize an individual security setting, the profile will become Custom. (https://support.cloudflare.com/hc/en-us/articles/200170056).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_security_level-2

Choose the appropriate security profile for your website, which will automatically adjust each of the security settings. If you choose to customize an individual security setting, the profile will become Custom. (https://support.cloudflare.com/hc/en-us/articles/200170056).

```yaml
{"description": "Choose the appropriate security profile for your website, which will automatically adjust each of the security settings. If you choose to customize an individual security setting, the profile will become Custom. (https://support.cloudflare.com/hc/en-us/articles/200170056).", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "security_level", "enum": ["security_level"]}, "value": {"$ref": "#/components/schemas/zones_security_level_value"}}}], "title": "Security Level"}
```
