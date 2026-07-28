---
title: magic_cf1_site_ramp_body
page_id: schema-magic-cf1-site-ramp-body-61ff32a1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_cf1_site_ramp_body

```yaml
{"type": "object", "properties": {"source_ramp_id": {"description": "Identifier of the source network resource to associate as a ramp.", "allOf": [{"$ref": "#/components/schemas/magic_identifier"}]}, "type": {"$ref": "#/components/schemas/magic_cf1_site_ramp_type"}}, "required": ["source_ramp_id", "type"]}
```
