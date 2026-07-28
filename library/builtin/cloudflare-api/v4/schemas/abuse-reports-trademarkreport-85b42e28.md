---
title: abuse-reports_TrademarkReport
page_id: schema-abuse-reports-trademarkreport-85b42e28
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_TrademarkReport

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/abuse-reports_BaseReportFields"}, {"properties": {"act": {"enum": ["abuse_trademark"]}, "host_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send"]}, "justification": {"description": "A detailed description of the infringement, including any necessary access details and the exact steps needed to view the content, not exceeding 5000 characters.\n", "type": "string", "maxLength": 5000, "minLength": 1}, "owner_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send"]}, "trademark_number": {"description": "Text not exceeding 1000 characters", "type": "string", "maxLength": 1000, "minLength": 1}, "trademark_office": {"description": "Text not exceeding 1000 characters", "type": "string", "maxLength": 1000, "minLength": 1}, "trademark_symbol": {"description": "Text not exceeding 1000 characters", "type": "string", "maxLength": 1000, "minLength": 1}}, "required": ["host_notification", "justification", "owner_notification", "trademark_number", "trademark_office", "trademark_symbol"], "title": "Trademark Report"}]}
```
