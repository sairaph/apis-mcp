---
title: intel_collection_response
page_id: schema-intel-collection-response-115398ca
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_collection_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/intel_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"additional_information": {"$ref": "#/components/schemas/intel_additional_information"}, "application": {"$ref": "#/components/schemas/intel_application"}, "content_categories": {"$ref": "#/components/schemas/intel_content_categories"}, "domain": {"$ref": "#/components/schemas/intel_domain_name"}, "inherited_content_categories": {"$ref": "#/components/schemas/intel_categories_with_super_category_ids_example_empty"}, "inherited_from": {"$ref": "#/components/schemas/intel_inherited_from"}, "inherited_risk_types": {"$ref": "#/components/schemas/intel_categories_with_super_category_ids_example_empty"}, "popularity_rank": {"$ref": "#/components/schemas/intel_popularity_rank"}, "risk_score": {"$ref": "#/components/schemas/intel_risk_score"}, "risk_types": {"$ref": "#/components/schemas/intel_categories_with_super_category_ids_example_empty"}}, "type": "object"}}}, "type": "object"}]}
```
