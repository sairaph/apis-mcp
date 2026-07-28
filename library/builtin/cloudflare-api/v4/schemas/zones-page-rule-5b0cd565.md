---
title: zones_page_rule
page_id: schema-zones-page-rule-5b0cd565
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_page_rule

```yaml
{"type": "object", "properties": {"actions": {"$ref": "#/components/schemas/zones_actions"}, "created_on": {"$ref": "#/components/schemas/zones_created_on"}, "id": {"$ref": "#/components/schemas/zones_identifier-2"}, "modified_on": {"$ref": "#/components/schemas/zones_modified_on"}, "priority": {"$ref": "#/components/schemas/zones_priority"}, "status": {"$ref": "#/components/schemas/zones_status"}, "targets": {"$ref": "#/components/schemas/zones_targets"}}, "required": ["id", "targets", "actions", "priority", "status", "modified_on", "created_on"]}
```
