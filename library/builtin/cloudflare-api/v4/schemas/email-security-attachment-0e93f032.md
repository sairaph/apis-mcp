---
title: email-security_Attachment
page_id: schema-email-security-attachment-0e93f032
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_Attachment

```yaml
{"type": "object", "properties": {"content_type": {"description": "MIME type of the attachment.", "type": "string", "nullable": true}, "detection": {"description": "Detection result for this attachment.", "type": "string", "allOf": [{"$ref": "#/components/schemas/email-security_DispositionLabel"}], "nullable": true}, "encrypted": {"description": "Whether the attachment is encrypted.", "type": "boolean", "nullable": true}, "filename": {"description": "Name of the attached file.", "type": "string", "nullable": true}, "md5": {"description": "MD5 hash of the attachment.", "type": "string", "nullable": true}, "name": {"description": "Attachment name (alternative to filename).", "type": "string", "nullable": true}, "sha1": {"description": "SHA1 hash of the attachment.", "type": "string", "nullable": true}, "sha256": {"description": "SHA256 hash of the attachment.", "type": "string", "nullable": true}, "size": {"description": "Size of the attachment in bytes.", "type": "integer", "minimum": 0}}, "required": ["size"]}
```
