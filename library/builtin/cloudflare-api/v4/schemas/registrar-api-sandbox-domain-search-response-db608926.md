---
title: registrar-api-sandbox_domain_search_response
page_id: schema-registrar-api-sandbox-domain-search-response-db608926
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_domain_search_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/registrar-api-sandbox_api-response-common"}, {"properties": {"result": {"description": "Contains the search results.", "type": "object", "properties": {"domains": {"description": "Array of domain suggestions sorted by relevance. May be empty if no domains match the search criteria.", "type": "array", "items": {"$ref": "#/components/schemas/registrar-api-sandbox_domain_search_result"}}}, "required": ["domains"]}}, "required": ["result"], "type": "object"}]}
```
