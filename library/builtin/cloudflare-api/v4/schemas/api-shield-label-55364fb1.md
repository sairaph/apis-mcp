---
title: api-shield_label
page_id: schema-api-shield-label-55364fb1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_label

```yaml
{"type": "object", "allOf": [{"properties": {"created_at": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "description": {"$ref": "#/components/schemas/api-shield_label_description"}, "last_updated": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "metadata": {"$ref": "#/components/schemas/api-shield_label_metadata"}, "name": {"$ref": "#/components/schemas/api-shield_label_name"}, "source": {"$ref": "#/components/schemas/api-shield_label_source"}}, "required": ["name", "description", "metadata", "source", "last_updated", "created_at"], "type": "object"}]}
```
