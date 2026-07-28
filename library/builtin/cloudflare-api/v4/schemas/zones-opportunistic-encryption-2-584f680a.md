---
title: zones_opportunistic_encryption-2
page_id: schema-zones-opportunistic-encryption-2-584f680a
path: schemas
description: Enables the Opportunistic Encryption feature for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_opportunistic_encryption-2

Enables the Opportunistic Encryption feature for a zone.

```yaml
{"description": "Enables the Opportunistic Encryption feature for a zone.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "opportunistic_encryption", "enum": ["opportunistic_encryption"]}, "value": {"$ref": "#/components/schemas/zones_opportunistic_encryption_value"}}}], "title": "Enable Opportunistic Encryption for a zone"}
```
