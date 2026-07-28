---
title: gelato_email_report_error
page_id: schema-gelato-email-report-error-ca27fa9c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_email_report_error

```yaml
{"title": "GelatoEmailReportError", "type": "object", "properties": {"code": {"type": "string", "description": "A short machine-readable string giving the reason for the verification failure.", "nullable": true, "enum": ["email_unverified_other", "email_verification_declined"]}, "reason": {"maxLength": 5000, "type": "string", "description": "A human-readable message giving the reason for the failure. These messages can be shown to your users.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
