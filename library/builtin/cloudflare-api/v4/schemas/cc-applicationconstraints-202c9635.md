---
title: cc_ApplicationConstraints
page_id: schema-cc-applicationconstraints-202c9635
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ApplicationConstraints

```yaml
{"type": "object", "properties": {"jurisdiction": {"description": "Currently only \"eu\" and \"fedramp\" are supported. Overlap between jurisdiction and region is allowed for ENAM, WNAM (FedRAMP) and EEUR, WEUR (EU).", "type": "string"}, "regions": {"type": "array", "items": {"$ref": "#/components/schemas/cc_Region"}}}}
```
