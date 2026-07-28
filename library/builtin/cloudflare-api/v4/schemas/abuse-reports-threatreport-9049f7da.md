---
title: abuse-reports_ThreatReport
page_id: schema-abuse-reports-threatreport-9049f7da
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_ThreatReport

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/abuse-reports_BaseReportFields"}, {"properties": {"act": {"enum": ["abuse_threat"]}, "host_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon"]}, "justification": {"description": "A detailed description of the infringement, including any necessary access details and the exact steps needed to view the content, not exceeding 5000 characters.\n", "type": "string", "maxLength": 5000, "minLength": 1}, "owner_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon"]}}, "required": ["host_notification", "justification", "owner_notification"], "title": "Threat Report"}]}
```
