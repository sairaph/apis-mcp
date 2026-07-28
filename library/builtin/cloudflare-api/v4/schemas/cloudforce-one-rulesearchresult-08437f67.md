---
title: cloudforce-one_RuleSearchResult
page_id: schema-cloudforce-one-rulesearchresult-08437f67
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_RuleSearchResult

```yaml
{"allOf": [{"$ref": "#/components/schemas/cloudforce-one_Rule"}, {"properties": {"score": {"description": "Relevance score in [0,1]. Present only when AI Search powers the query.", "type": "number", "example": 0.87, "maximum": 1, "minimum": 0}, "scoring_details": {"description": "Per-component scoring breakdown returned by AI Search hybrid retrieval.", "type": "object", "properties": {"fusion_method": {"type": "string"}, "keyword_rank": {"type": "number"}, "keyword_score": {"type": "number"}, "reranking_score": {"type": "number"}, "vector_rank": {"type": "number"}, "vector_score": {"type": "number"}}}}, "type": "object"}]}
```
