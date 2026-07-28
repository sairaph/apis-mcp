---
title: digital-experience-monitoring_dex-rule
page_id: schema-digital-experience-monitoring-dex-rule-a75963a1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_dex-rule

```yaml
{"type": "object", "properties": {"created_at": {"type": "string", "example": "2023-07-16 15:00:00+00", "x-auditable": true, "x-stainless-terraform-configurability": "computed"}, "description": {"type": "string"}, "id": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}], "x-auditable": true, "x-stainless-terraform-configurability": "computed"}, "match": {"type": "string"}, "name": {"type": "string", "x-auditable": true}, "targeted_tests": {"type": "array", "items": {"$ref": "#/components/schemas/digital-experience-monitoring_dex-targeted-test"}, "x-auditable": true, "x-stainless-terraform-configurability": "computed"}, "updated_at": {"type": "string", "example": "2023-07-16 15:00:00+00", "x-auditable": true, "x-stainless-terraform-configurability": "computed"}}, "required": ["id", "name", "match", "created_at"]}
```
