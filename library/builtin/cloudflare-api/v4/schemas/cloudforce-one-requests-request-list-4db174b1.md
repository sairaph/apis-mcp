---
title: cloudforce-one-requests_request-list
page_id: schema-cloudforce-one-requests-request-list-4db174b1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_request-list

```yaml
{"type": "object", "properties": {"completed_after": {"description": "Retrieve requests completed after this time.", "example": "2022-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "completed_before": {"description": "Retrieve requests completed before this time.", "example": "2024-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "created_after": {"description": "Retrieve requests created after this time.", "example": "2022-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "created_before": {"description": "Retrieve requests created before this time.", "example": "2024-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "page": {"description": "Page number of results.", "type": "integer"}, "per_page": {"description": "Number of results per page.", "type": "integer", "example": 10}, "request_type": {"$ref": "#/components/schemas/cloudforce-one-requests_request-type"}, "sort_by": {"description": "Field to sort results by.", "type": "string", "example": "created"}, "sort_order": {"description": "Sort order (asc or desc).", "type": "string", "enum": ["asc", "desc"]}, "status": {"$ref": "#/components/schemas/cloudforce-one-requests_request-status"}}, "required": ["page", "per_page"], "title": "Request List Parameters"}
```
