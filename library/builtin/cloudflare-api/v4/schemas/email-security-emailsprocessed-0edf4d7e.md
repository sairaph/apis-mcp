---
title: email-security_EmailsProcessed
page_id: schema-email-security-emailsprocessed-0edf4d7e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_EmailsProcessed

```yaml
{"type": "object", "properties": {"timestamp": {"type": "string", "format": "date-time", "readOnly": true}, "total_emails_processed": {"type": "integer", "minimum": 0}, "total_emails_processed_previous": {"type": "integer", "minimum": 0}}, "required": ["total_emails_processed", "total_emails_processed_previous", "timestamp"]}
```
