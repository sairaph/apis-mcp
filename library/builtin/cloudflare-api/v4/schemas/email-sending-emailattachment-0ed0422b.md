---
title: email-sending_EmailAttachment
page_id: schema-email-sending-emailattachment-0ed0422b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-sending_EmailAttachment

```yaml
{"type": "object", "properties": {"content": {"description": "Base64-encoded content of the attachment.", "example": "JVBERi0xLjQK...", "type": "string"}, "disposition": {"description": "Must be 'attachment'. Adds a standard file attachment.", "type": "string", "enum": ["attachment"]}, "filename": {"description": "Filename for the attachment.", "type": "string", "example": "report.pdf"}, "type": {"description": "MIME type of the attachment (e.g., 'application/pdf', 'text/plain').", "type": "string", "example": "application/pdf"}}, "example": {"content": "JVBERi0xLjQK...", "disposition": "attachment", "filename": "report.pdf", "type": "application/pdf"}, "required": ["disposition", "filename", "type", "content"]}
```
