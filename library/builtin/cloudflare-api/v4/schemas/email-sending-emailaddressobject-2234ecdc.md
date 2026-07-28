---
title: email-sending_EmailAddressObject
page_id: schema-email-sending-emailaddressobject-2234ecdc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-sending_EmailAddressObject

```yaml
{"type": "object", "properties": {"address": {"description": "Email address (e.g., 'user@example.com').", "type": "string", "example": "user@example.com", "pattern": "^[\\x20-\\x7E]+$"}, "name": {"description": "Display name for the email address (e.g., 'John Doe'). Optional; set to null or leave it unset to send the address on its own.", "type": "string", "example": "John Doe", "nullable": true}}, "example": {"address": "user@example.com", "name": "John Doe"}, "required": ["address"]}
```
