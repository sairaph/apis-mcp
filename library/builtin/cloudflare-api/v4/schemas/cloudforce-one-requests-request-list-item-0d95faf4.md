---
title: cloudforce-one-requests_request-list-item
page_id: schema-cloudforce-one-requests-request-list-item-0d95faf4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_request-list-item

```yaml
{"type": "object", "properties": {"completed": {"description": "Request completion time.", "example": "2024-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "created": {"description": "Request creation time.", "example": "2022-04-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "id": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}, "message_tokens": {"description": "Tokens for the request messages.", "type": "integer", "example": 16, "x-auditable": true}, "priority": {"$ref": "#/components/schemas/cloudforce-one-requests_priority"}, "readable_id": {"$ref": "#/components/schemas/cloudforce-one-requests_request-readable-id"}, "request": {"$ref": "#/components/schemas/cloudforce-one-requests_request-type"}, "status": {"$ref": "#/components/schemas/cloudforce-one-requests_request-status"}, "summary": {"$ref": "#/components/schemas/cloudforce-one-requests_request-summary"}, "tlp": {"$ref": "#/components/schemas/cloudforce-one-requests_tlp"}, "tokens": {"description": "Tokens for the request.", "type": "integer", "x-auditable": true}, "updated": {"description": "Request last updated time.", "example": "2022-04-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}}, "required": ["id", "created", "updated", "priority", "request", "summary", "tlp"], "title": "Request List Item"}
```
