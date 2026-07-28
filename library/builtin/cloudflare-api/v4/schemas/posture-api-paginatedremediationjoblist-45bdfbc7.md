---
title: posture-api_PaginatedRemediationJobList
page_id: schema-posture-api-paginatedremediationjoblist-45bdfbc7
path: schemas
description: Paginated list of remediation jobs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_PaginatedRemediationJobList

Paginated list of remediation jobs.

```yaml
{"description": "Paginated list of remediation jobs.", "type": "object", "properties": {"errors": {"description": "Array of error messages.", "type": "array", "items": {"type": "object"}, "example": []}, "messages": {"description": "Array of informational messages.", "type": "array", "items": {"type": "object"}, "example": []}, "result": {"description": "Array of remediation job objects.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_RemediationJob"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Number of results on current page.", "type": "integer", "example": 2}, "cursor": {"description": "Cursor for pagination.", "type": "string", "example": "next_cursor_value", "nullable": true}, "page": {"description": "Current page number.", "type": "integer", "example": 1}, "per_page": {"description": "Number of results per page.", "type": "integer", "example": 10}, "total_count": {"description": "Total number of results.", "type": "integer", "example": 2}}}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true}}, "required": ["success", "errors", "messages", "result_info", "result"]}
```
