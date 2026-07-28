---
title: zaraz_zaraz-config-body
page_id: schema-zaraz-zaraz-config-body-7f9c2709
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_zaraz-config-body

```yaml
{"allOf": [{"$ref": "#/components/schemas/zaraz_zaraz-config-base"}, {"properties": {"tools": {"description": "Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID and value is the tool configuration object.", "type": "object", "additionalProperties": {"anyOf": [{"$ref": "#/components/schemas/zaraz_managed-component"}, {"$ref": "#/components/schemas/zaraz_custom-managed-component"}]}}}, "type": "object"}]}
```
