---
title: cloudforce-one-requests_request-message-item
page_id: schema-cloudforce-one-requests-request-message-item-f4e2605f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_request-message-item

```yaml
{"type": "object", "properties": {"author": {"description": "Author of message.", "type": "string", "example": "user@domain.com", "x-auditable": true}, "content": {"$ref": "#/components/schemas/cloudforce-one-requests_message-content"}, "created": {"description": "Defines the message creation time.", "example": "2022-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "id": {"description": "Message ID.", "type": "integer", "x-auditable": true}, "is_follow_on_request": {"description": "Whether the message is a follow-on request.", "type": "boolean", "x-auditable": true}, "updated": {"description": "Defines the message last updated time.", "example": "2022-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}}, "required": ["id", "updated", "content", "author", "is_follow_on_request"], "title": "Request Message Item"}
```
