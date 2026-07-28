---
title: gelato_phone_report_error
page_id: schema-gelato-phone-report-error-fbd85cff
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_phone_report_error

```yaml
{"title": "GelatoPhoneReportError", "type": "object", "properties": {"code": {"type": "string", "description": "A short machine-readable string giving the reason for the verification failure.", "nullable": true, "enum": ["phone_unverified_other", "phone_verification_declined"]}, "reason": {"maxLength": 5000, "type": "string", "description": "A human-readable message giving the reason for the failure. These messages can be shown to your users.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
