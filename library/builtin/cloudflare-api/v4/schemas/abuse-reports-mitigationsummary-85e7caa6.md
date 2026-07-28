---
title: abuse-reports_MitigationSummary
page_id: schema-abuse-reports-mitigationsummary-85e7caa6
path: schemas
description: A summary of the mitigations related to this report.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_MitigationSummary

A summary of the mitigations related to this report.

```yaml
{"description": "A summary of the mitigations related to this report.", "type": "object", "properties": {"accepted_url_count": {"description": "How many of the reported URLs were confirmed as abusive.", "type": "integer"}, "active_count": {"description": "How many mitigations are active.", "type": "integer"}, "external_host_notified": {"description": "Whether the report has been forwarded to an external hosting provider.", "type": "boolean"}, "in_review_count": {"description": "How many mitigations are under review.", "type": "integer"}, "pending_count": {"description": "How many mitigations are pending their effective date.", "type": "integer"}}, "required": ["accepted_url_count", "pending_count", "active_count", "in_review_count", "external_host_notified"]}
```
