---
title: magic_cf1_site
page_id: schema-magic-cf1-site-84481b1f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_cf1_site

```yaml
{"type": "object", "properties": {"created_on": {"type": "string", "format": "date-time", "readOnly": true}, "description": {"description": "A human-provided description of the CF1 Site.", "type": "string", "example": "Launch Pad 34"}, "id": {"allOf": [{"$ref": "#/components/schemas/magic_identifier"}], "readOnly": true}, "location": {"$ref": "#/components/schemas/magic_cf1_site_location"}, "modified_on": {"type": "string", "format": "date-time", "readOnly": true}, "name": {"description": "A human-provided name describing the CF1 Site that should be unique within the account.", "type": "string", "example": "Pad 34"}}, "required": ["name"]}
```
