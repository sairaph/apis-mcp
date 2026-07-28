---
title: email-security_BulkMessageActionParams
page_id: schema-email-security-bulkmessageactionparams-20b40dfa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_BulkMessageActionParams

```yaml
{"discriminator": {"propertyName": "type"}, "oneOf": [{"properties": {"client_recipient": {"type": "string"}, "destination": {"$ref": "#/components/schemas/email-security_MailboxDestination"}, "expected_disposition": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "type": {"type": "string", "enum": ["MOVE"]}}, "required": ["type", "destination", "client_recipient"], "type": "object"}, {"properties": {"client_recipient": {"type": "string"}, "type": {"type": "string", "enum": ["RELEASE"]}}, "required": ["type", "client_recipient"], "type": "object"}]}
```
