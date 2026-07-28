---
title: email-sending_EmailInlineAttachment
page_id: schema-email-sending-emailinlineattachment-0cd38f5c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-sending_EmailInlineAttachment

```yaml
{"type": "object", "properties": {"content": {"description": "Base64-encoded content of the attachment.", "example": "iVBORw0KGgoAAAANSUhEUgAA...", "type": "string"}, "content_id": {"description": "Content ID used to reference this attachment in HTML via cid: URI (e.g., <img src=\"cid:logo\">).", "type": "string", "example": "logo"}, "disposition": {"description": "Must be 'inline'. Embeds the attachment in the email body.", "type": "string", "enum": ["inline"]}, "filename": {"description": "Filename for the attachment.", "type": "string", "example": "logo.png"}, "type": {"description": "MIME type of the attachment (e.g., 'image/png', 'text/plain').", "type": "string", "example": "image/png"}}, "example": {"content": "iVBORw0KGgoAAAANSUhEUgAA...", "content_id": "logo", "disposition": "inline", "filename": "logo.png", "type": "image/png"}, "required": ["disposition", "content_id", "filename", "type", "content"]}
```
