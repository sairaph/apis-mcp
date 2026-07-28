---
title: cloudforce-one-requests_request-item
page_id: schema-cloudforce-one-requests-request-item-feceecaa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_request-item

```yaml
{"type": "object", "properties": {"completed": {"$ref": "#/components/schemas/cloudforce-one-requests_time"}, "content": {"$ref": "#/components/schemas/cloudforce-one-requests_request-content"}, "created": {"$ref": "#/components/schemas/cloudforce-one-requests_time"}, "id": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}, "message_tokens": {"description": "Tokens for the request messages.", "type": "integer", "example": 1, "x-auditable": true}, "priority": {"$ref": "#/components/schemas/cloudforce-one-requests_time"}, "readable_id": {"$ref": "#/components/schemas/cloudforce-one-requests_request-readable-id"}, "request": {"$ref": "#/components/schemas/cloudforce-one-requests_request-type"}, "status": {"$ref": "#/components/schemas/cloudforce-one-requests_request-status"}, "summary": {"$ref": "#/components/schemas/cloudforce-one-requests_request-summary"}, "tlp": {"$ref": "#/components/schemas/cloudforce-one-requests_tlp"}, "tokens": {"description": "Tokens for the request.", "type": "integer", "example": 16, "x-auditable": true}, "updated": {"$ref": "#/components/schemas/cloudforce-one-requests_time"}}, "required": ["id", "created", "updated", "content", "priority", "request", "summary", "tlp"], "title": "Request Item"}
```
