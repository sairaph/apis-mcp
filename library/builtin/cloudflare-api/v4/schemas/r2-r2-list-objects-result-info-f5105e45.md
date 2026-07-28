---
title: r2_r2_list_objects_result_info
page_id: schema-r2-r2-list-objects-result-info-f5105e45
path: schemas
description: Pagination information for list objects responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_r2_list_objects_result_info

Pagination information for list objects responses.

```yaml
{"description": "Pagination information for list objects responses.", "type": "object", "properties": {"cursor": {"description": "Pagination cursor to use in the next List Objects call to retrieve the next page of results.", "type": "string", "example": "eyJrZXkiOiJwYXRoL3RvL215LW9iamVjdC50eHQifQ=="}, "delimited": {"description": "Common prefixes found when a delimiter is specified. Each entry represents a group\nof keys sharing a common prefix up to the delimiter. Equivalent to S3's `CommonPrefixes`\nin `ListObjectsV2`; the field name differs because of the existing R2 API wire format.\n", "type": "array", "items": {"type": "string"}, "example": ["path/to/", "another/path/"]}, "is_truncated": {"description": "Whether the result was truncated. If true, use the cursor to retrieve the next page.", "type": "boolean", "example": true}, "per_page": {"description": "The maximum number of objects returned per page.", "type": "integer", "example": 20}}}
```
