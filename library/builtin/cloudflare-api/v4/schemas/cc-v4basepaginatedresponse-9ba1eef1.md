---
title: cc_V4BasePaginatedResponse
page_id: schema-cc-v4basepaginatedresponse-9ba1eef1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_V4BasePaginatedResponse

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result_info": {"type": "object", "properties": {"next_page_token": {"description": "The token to use to retrieve the next page of results.", "type": "string"}, "page_token": {"description": "The page token sent in the request.", "type": "string"}, "per_page": {"description": "The number of items per page requested.", "type": "integer"}}}}, "required": ["result_info"], "type": "object"}]}
```
