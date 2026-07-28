---
title: cloudforce-one_RulesSearchResponse
page_id: schema-cloudforce-one-rulessearchresponse-1a1d5fff
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_RulesSearchResponse

```yaml
{"type": "object", "properties": {"fallback": {"description": "True when AI Search was unavailable and the response is from a fallback path.", "type": "boolean", "example": false}, "interpreted": {"description": "Parsed natural-language interpretation of the query, when available.", "type": "object", "properties": {"filters": {"description": "Filters applied during retrieval (account ACL plus user-supplied facets).", "type": "object", "additionalProperties": {"nullable": true, "type": "object"}}, "retrieval_type": {"type": "string"}}, "required": ["retrieval_type"]}, "mode": {"description": "Retrieval strategy actually used to produce results.", "type": "string", "example": "hybrid"}, "results": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one_RuleSearchResult"}}, "total": {"type": "integer", "example": 12, "minimum": 0}}, "required": ["results", "total", "mode"]}
```
