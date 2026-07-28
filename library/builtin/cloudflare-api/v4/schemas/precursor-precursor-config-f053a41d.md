---
title: precursor_precursor_config
page_id: schema-precursor-precursor-config-f053a41d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# precursor_precursor_config

```yaml
{"type": "object", "properties": {"default_mode": {"$ref": "#/components/schemas/precursor_default_mode"}, "enforcement_rules": {"$ref": "#/components/schemas/precursor_enforcement_rules"}}, "example": {"default_mode": "min-friction", "enforcement_rules": [{"description": "Enforce max-security on the shop page", "enabled": true, "expression": "http.request.uri.path eq \"/shop\"", "mode": "max-security"}]}, "minProperties": 1, "title": "Precursor Config"}
```
