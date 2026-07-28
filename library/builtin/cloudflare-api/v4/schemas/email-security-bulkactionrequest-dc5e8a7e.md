---
title: email-security_BulkActionRequest
page_id: schema-email-security-bulkactionrequest-dc5e8a7e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_BulkActionRequest

```yaml
{"type": "object", "properties": {"action": {"type": "string", "enum": ["MOVE", "RELEASE"]}, "comment": {"type": "string", "nullable": true}, "destination": {"$ref": "#/components/schemas/email-security_MailboxDestination"}, "expected_disposition": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "search_params": {"$ref": "#/components/schemas/email-security_BulkSearchParams"}}, "required": ["action", "search_params"]}
```
