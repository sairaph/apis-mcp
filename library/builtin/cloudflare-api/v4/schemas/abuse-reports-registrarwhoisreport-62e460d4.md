---
title: abuse-reports_RegistrarWhoisReport
page_id: schema-abuse-reports-registrarwhoisreport-62e460d4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_RegistrarWhoisReport

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/abuse-reports_BaseReportFields"}, {"properties": {"act": {"enum": ["abuse_registrar_whois"]}, "owner_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon", "none"]}, "reg_who_request": {"$ref": "#/components/schemas/abuse-reports_RegistrarWhoIsFields"}}, "required": ["owner_notification"], "title": "Registrar Whois Report"}]}
```
