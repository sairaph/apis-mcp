---
title: posture-api_FindingCategory
page_id: schema-posture-api-findingcategory-915e7d19
path: schemas
description: Category information for a finding.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingCategory

Category information for a finding.

```yaml
{"description": "Category information for a finding.", "type": "object", "properties": {"observation": {"$ref": "#/components/schemas/posture-api_ObservationEnum"}, "product": {"$ref": "#/components/schemas/posture-api_ProductEnum"}, "type": {"$ref": "#/components/schemas/posture-api_FindingCategoryTypeEnum"}}, "example": {"observation": "Issue", "product": "SaaS", "type": "Posture"}, "required": ["observation", "product", "type"]}
```
