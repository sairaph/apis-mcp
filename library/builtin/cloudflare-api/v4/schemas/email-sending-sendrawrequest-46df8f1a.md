---
title: email-sending_SendRawRequest
page_id: schema-email-sending-sendrawrequest-46df8f1a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-sending_SendRawRequest

```yaml
{"type": "object", "properties": {"from": {"description": "Sender email address.", "type": "string", "example": "sender@example.com", "pattern": "^[\\x20-\\x7E]+$"}, "mime_message": {"description": "The full MIME-encoded email message. Should include standard RFC 5322 headers such as From, To, Subject, and Content-Type. The from and recipients fields in the request body control SMTP envelope routing; the From and To headers in the MIME message control what the recipient's email client displays.", "type": "string", "example": "From: sender@example.com\r\nTo: recipient@example.com\r\nSubject: Hello\r\nContent-Type: text/plain\r\n\r\nHello, World!"}, "recipients": {"description": "List of recipient email addresses.", "type": "array", "items": {"pattern": "^[\\x20-\\x7E]+$", "type": "string"}, "example": ["recipient@example.com"]}}, "example": {"from": "sender@example.com", "mime_message": "From: sender@example.com\r\nTo: recipient@example.com\r\nSubject: Hello\r\nContent-Type: text/plain\r\n\r\nHello, World!", "recipients": ["recipient@example.com"]}, "additionalProperties": false, "required": ["from", "recipients", "mime_message"]}
```
