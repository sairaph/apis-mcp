---
title: gelato_email_report
page_id: schema-gelato-email-report-ee836a0f
path: schemas
description: Result from a email check
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_email_report

Result from a email check

```yaml
{"title": "GelatoEmailReport", "required": ["status"], "type": "object", "properties": {"email": {"maxLength": 5000, "type": "string", "description": "Email to be verified.", "nullable": true}, "error": {"description": "Details on the verification error. Present when status is `unverified`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_email_report_error"}]}, "status": {"type": "string", "description": "Status of this `email` check.", "enum": ["unverified", "verified"], "x-stripeBypassValidation": true}}, "description": "Result from a email check", "x-expandableFields": ["error"]}
```
