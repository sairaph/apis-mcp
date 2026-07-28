---
title: abuse-reports_NCSEIReport
page_id: schema-abuse-reports-ncseireport-4f8ed3ce
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_NCSEIReport

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/abuse-reports_BaseReportFields"}, {"properties": {"act": {"enum": ["abuse_ncsei"]}, "country": {"description": "Text not exceeding 255 characters. This field may be released by Cloudflare to third parties such as the Lumen Database (https://lumendatabase.org/).\n", "type": "string", "maxLength": 255, "minLength": 1}, "host_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon"]}, "ncsei_subject_representation": {"description": "If the submitter is the target of NCSEI in the URLs of the abuse report.", "type": "boolean"}, "owner_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon", "none"]}}, "required": ["host_notification", "ncsei_subject_representation", "owner_notification"], "title": "NCSEI Report"}]}
```
