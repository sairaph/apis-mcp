---
title: email-sending_EmailSendResponse
page_id: schema-email-sending-emailsendresponse-fb65c506
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-sending_EmailSendResponse

```yaml
{"type": "object", "properties": {"delivered": {"description": "Email addresses to which the message was delivered immediately.", "type": "array", "items": {"pattern": "^[\\x20-\\x7E]+$", "type": "string"}, "example": ["recipient@example.com"]}, "message_id": {"description": "Message ID of the sent email.", "type": "string", "example": "<aB3xK9mP2qR5sT8uV0wX1yZ4cD6fG7hJ9kL0@example.com>", "pattern": "^[\\x20-\\x7E]+$"}, "permanent_bounces": {"description": "Email addresses that permanently bounced.", "type": "array", "items": {"pattern": "^[\\x20-\\x7E]+$", "type": "string"}, "example": []}, "queued": {"description": "Email addresses for which delivery was queued for later.", "type": "array", "items": {"pattern": "^[\\x20-\\x7E]+$", "type": "string"}, "example": []}}, "example": {"delivered": ["recipient@example.com"], "message_id": "<aB3xK9mP2qR5sT8uV0wX1yZ4cD6fG7hJ9kL0@example.com>", "permanent_bounces": [], "queued": []}, "required": ["message_id", "delivered", "queued", "permanent_bounces"]}
```
