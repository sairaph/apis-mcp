---
title: zones_security_level
page_id: schema-zones-security-level-ae4daad3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_security_level

```yaml
{"type": "object", "properties": {"id": {"description": "Control options for the **Security Level** feature from the **Security** app.\n", "type": "string", "enum": ["security_level"], "x-auditable": true}, "value": {"type": "string", "example": "under_attack", "enum": ["off", "essentially_off", "low", "medium", "high", "under_attack"], "x-auditable": true}}, "title": "Security Level", "x-stainless-skip": ["terraform"]}
```
