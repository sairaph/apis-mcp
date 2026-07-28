---
title: zones_opportunistic_encryption
page_id: schema-zones-opportunistic-encryption-887b3f3f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_opportunistic_encryption

```yaml
{"type": "object", "properties": {"id": {"description": "Opportunistic Encryption allows browsers to access HTTP URIs over an encrypted TLS channel.\nIt's not a substitute for HTTPS, but provides additional security for otherwise vulnerable requests.\n", "type": "string", "example": "opportunistic_encryption", "enum": ["opportunistic_encryption"], "x-auditable": true}, "value": {"description": "The status of Opportunistic Encryption.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Opportunistic Encryption", "x-stainless-skip": ["terraform"]}
```
