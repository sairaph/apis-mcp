---
title: abuse-reports_CSAMReport
page_id: schema-abuse-reports-csamreport-b4a0a5be
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_CSAMReport

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/abuse-reports_BaseReportFields"}, {"properties": {"act": {"enum": ["abuse_children"]}, "country": {"description": "Text not exceeding 255 characters. This field may be released by Cloudflare to third parties such as the Lumen Database (https://lumendatabase.org/).\n", "type": "string", "maxLength": 255, "minLength": 1}, "host_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon"]}, "justification": {"description": "A detailed description of the infringement, including any necessary access details and the exact steps needed to view the content, not exceeding 5000 characters.\n", "type": "string", "maxLength": 5000, "minLength": 1}, "ncmec_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon"]}, "owner_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon", "none"]}}, "required": ["host_notification", "justification", "ncmec_notification", "owner_notification"], "title": "Children Abuse Report"}]}
```
