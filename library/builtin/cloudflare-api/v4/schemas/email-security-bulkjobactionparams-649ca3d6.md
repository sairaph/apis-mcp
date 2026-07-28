---
title: email-security_BulkJobActionParams
page_id: schema-email-security-bulkjobactionparams-649ca3d6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_BulkJobActionParams

```yaml
{"discriminator": {"propertyName": "type"}, "oneOf": [{"properties": {"destination": {"$ref": "#/components/schemas/email-security_MailboxDestination"}, "expected_disposition": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "type": {"type": "string", "enum": ["MOVE"]}}, "required": ["type", "destination"], "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["RELEASE"]}}, "required": ["type"], "type": "object"}]}
```
