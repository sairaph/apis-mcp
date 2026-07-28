---
title: pay-per-crawl_ResultInfo
page_id: schema-pay-per-crawl-resultinfo-a3348e7a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pay-per-crawl_ResultInfo

```yaml
{"type": "object", "properties": {"count": {"type": "integer"}, "page": {"type": "integer"}, "per_page": {"type": "integer"}, "total_count": {"type": "integer"}, "total_pages": {"description": "TotalPages is a pointer so that if TotalPages == 0 we return that there\nare indeed 0 pages. omitempty would have removed the field otherwise.\nThis is important as a customer may be relying on always reading this\nproperty and it should not be absent just because it is 0, only absent\nif a value is never provided.", "type": "integer"}}}
```
