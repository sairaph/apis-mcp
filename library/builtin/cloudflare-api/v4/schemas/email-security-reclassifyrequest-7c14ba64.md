---
title: email-security_ReclassifyRequest
page_id: schema-email-security-reclassifyrequest-7c14ba64
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_ReclassifyRequest

```yaml
{"type": "object", "properties": {"eml_content": {"description": "Base64 encoded content of the EML file.", "type": "string"}, "escalated_submission_id": {"type": "string", "x-auditable": true}, "expected_disposition": {"type": "string", "enum": ["NONE", "BULK", "MALICIOUS", "SPAM", "SPOOF", "SUSPICIOUS"]}}, "required": ["expected_disposition"]}
```
