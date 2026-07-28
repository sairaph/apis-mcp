---
title: gelato_phone_report
page_id: schema-gelato-phone-report-c980d055
path: schemas
description: Result from a phone check
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_phone_report

Result from a phone check

```yaml
{"title": "GelatoPhoneReport", "required": ["status"], "type": "object", "properties": {"error": {"description": "Details on the verification error. Present when status is `unverified`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_phone_report_error"}]}, "phone": {"maxLength": 5000, "type": "string", "description": "Phone to be verified.", "nullable": true}, "status": {"type": "string", "description": "Status of this `phone` check.", "enum": ["unverified", "verified"], "x-stripeBypassValidation": true}}, "description": "Result from a phone check", "x-expandableFields": ["error"]}
```
