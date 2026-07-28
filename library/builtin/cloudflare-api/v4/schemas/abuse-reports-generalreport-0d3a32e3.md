---
title: abuse-reports_GeneralReport
page_id: schema-abuse-reports-generalreport-0d3a32e3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_GeneralReport

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/abuse-reports_BaseReportFields"}, {"properties": {"act": {"enum": ["abuse_general"]}, "destination_ips": {"description": "A list of IP addresses separated by ‘\\n’ (new line character). The list of destination IPs should not exceed 30 IP addresses. Each one of the IP addresses ought to be unique.", "type": "string"}, "host_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon"]}, "justification": {"description": "A detailed description of the infringement, including any necessary access details and the exact steps needed to view the content, not exceeding 5000 characters.\n", "type": "string", "maxLength": 5000, "minLength": 1}, "owner_notification": {"description": "Notification type based on the abuse type. NOTE: Copyright (DMCA) and Trademark reports cannot be anonymous.\n", "type": "string", "enum": ["send", "send-anon"]}, "ports_protocols": {"description": "A comma separated list of ports and protocols e.g. 80/TCP, 22/UDP. The total size of the field should not exceed 2000 characters. Each individual port/protocol should not exceed 100 characters. The list should not have more than 30 unique ports and protocols.", "type": "string"}, "source_ips": {"description": "A list of IP addresses separated by ‘\\n’ (new line character). The list of source IPs should not exceed 30 IP addresses. Each one of the IP addresses ought to be unique.", "type": "string"}}, "required": ["host_notification", "justification", "owner_notification"], "title": "General Report"}]}
```
