---
title: WebSearchResponse
page_id: schema-websearchresponse-8bc345db
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# WebSearchResponse

```yaml
{"type": "object", "properties": {"id": {"description": "Task ID.", "type": "string"}, "created": {"description": "Request creation time, Unix timestamp in seconds.", "type": "integer"}, "search_result": {"description": "Search results.", "type": "array", "items": {"$ref": "#/components/schemas/WebSearchObjectResponse"}}}}
```
