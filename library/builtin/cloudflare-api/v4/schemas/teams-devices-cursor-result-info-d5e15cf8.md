---
title: teams-devices_cursor_result_info
page_id: schema-teams-devices-cursor-result-info-d5e15cf8
path: schemas
description: V4 public API Pagination/Cursor info.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_cursor_result_info

V4 public API Pagination/Cursor info.

```yaml
{"description": "V4 public API Pagination/Cursor info.", "type": "object", "properties": {"count": {"description": "Number of records in the response.", "type": "integer"}, "cursor": {"description": "Opaque token to request the next set of records.", "type": "string"}, "per_page": {"description": "The limit for the number of records in the response.", "type": "integer"}, "total_count": {"description": "Total number of records available.", "type": "integer", "nullable": true}}, "example": {"count": 1, "cursor": "ais86dftf.asdf7ba8", "per_page": 10, "total_count": null}, "required": ["per_page", "count", "cursor"]}
```
