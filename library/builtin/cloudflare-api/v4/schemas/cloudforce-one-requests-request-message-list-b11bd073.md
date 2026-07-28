---
title: cloudforce-one-requests_request-message-list
page_id: schema-cloudforce-one-requests-request-message-list-b11bd073
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_request-message-list

```yaml
{"type": "object", "properties": {"after": {"description": "Retrieve mes  ges created after this time.", "example": "2022-01-01T00:00:Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "before": {"description": "Retrieve messages created before this time.", "example": "2024-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "page": {"description": "Page number of results.", "type": "integer"}, "per_page": {"description": "Number of results per page.", "type": "integer", "example": 10}, "sort_by": {"description": "Field to sort results by.", "type": "string", "example": "created"}, "sort_order": {"description": "Sort order (asc or desc).", "type": "string", "enum": ["asc", "desc"]}}, "required": ["page", "per_page"], "title": "Request Message List Parameters"}
```
