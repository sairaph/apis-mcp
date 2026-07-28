---
title: posture-api_result-info
page_id: schema-posture-api-result-info-f9e0c43b
path: schemas
description: Pagination and result information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_result-info

Pagination and result information.

```yaml
{"description": "Pagination and result information.", "type": "object", "properties": {"count": {"description": "Total number of results for the requested service.", "type": "integer", "example": 1}, "cursor": {"description": "Cursor for cursor-based pagination.", "type": "string", "example": "eyJpZCI6IjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCIsImFmZmxpY3Rpb25fZGF0ZSI6IjE5NzAtMDEtMDFUMDA6MDA6MDAuMDAwMDAwWiJ9", "nullable": true}, "next": {"description": "URL to the next page of results.", "type": "string", "format": "uri", "nullable": true}, "page": {"description": "Current page within paginated list of results.", "type": "integer", "example": 1}, "per_page": {"description": "Number of results per page of results.", "type": "integer", "example": 20}, "previous": {"description": "URL to the previous page of results.", "type": "string", "format": "uri", "nullable": true}, "total_count": {"description": "Total results available without any search parameters.", "type": "integer", "example": 2000}}}
```
