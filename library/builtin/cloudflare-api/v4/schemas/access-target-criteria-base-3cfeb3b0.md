---
title: access_target_criteria_base
page_id: schema-access-target-criteria-base-3cfeb3b0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_target_criteria_base

```yaml
{"type": "object", "properties": {"port": {"$ref": "#/components/schemas/access_port"}, "target_attributes": {"$ref": "#/components/schemas/access_target_attributes"}}, "required": ["target_attributes", "port", "protocol"], "title": "Target Criteria"}
```
