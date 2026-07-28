---
title: email-security_CursorResultInfo
page_id: schema-email-security-cursorresultinfo-1d7706f5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_CursorResultInfo

```yaml
{"type": "object", "properties": {"count": {"description": "Number of items in current page.", "type": "integer", "minimum": 0}, "next": {"description": "Cursor for next page.", "type": "string", "nullable": true}, "page": {"description": "Deprecated: Always returns 0. End of life: November 1, 2026.", "type": "integer", "deprecated": true, "minimum": 0, "x-stainless-deprecation-message": "This field always returns 0 and is deprecated."}, "per_page": {"description": "Number of items per page.", "type": "integer", "minimum": 0}, "previous": {"description": "Cursor for previous page.", "type": "string", "nullable": true}, "total_count": {"description": "Deprecated: Always returns 0. End of life: November 1, 2026.", "type": "integer", "deprecated": true, "minimum": 0, "x-stainless-deprecation-message": "This field always returns 0 and is deprecated."}}, "required": ["per_page", "count", "total_count"]}
```
