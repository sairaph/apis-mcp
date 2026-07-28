---
title: registrar-api-sandbox_cursor_result_info
page_id: schema-registrar-api-sandbox-cursor-result-info-a3f2ca2a
path: schemas
description: |-
    Cursor-based pagination metadata. Used by list endpoints that support
    cursor pagination. Pass the `cursor` value as a query parameter in the
    next request to fetch the next page. An empty string indicates there
    are no more pages.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_cursor_result_info

Cursor-based pagination metadata. Used by list endpoints that support
cursor pagination. Pass the `cursor` value as a query parameter in the
next request to fetch the next page. An empty string indicates there
are no more pages.

```yaml
{"description": "Cursor-based pagination metadata. Used by list endpoints that support\ncursor pagination. Pass the `cursor` value as a query parameter in the\nnext request to fetch the next page. An empty string indicates there\nare no more pages.\n", "type": "object", "properties": {"count": {"description": "Number of items in the current result set.", "type": "integer", "example": 20}, "cursor": {"description": "Opaque cursor for fetching the next page. Pass this value as the\n`cursor` query parameter in a subsequent request. An empty string\nindicates there are no more pages.\n", "type": "string", "example": "eyJ0IjoiMjAyNS0wNi0xNVQxMjowMDowMC4wMDAwMDBaIiwibiI6ImJyYXZvLm5ldCJ9"}, "per_page": {"description": "Maximum number of items per page.", "type": "integer", "example": 20}}, "required": ["cursor", "per_page", "count"]}
```
