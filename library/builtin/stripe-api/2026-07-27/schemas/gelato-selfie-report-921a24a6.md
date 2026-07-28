---
title: gelato_selfie_report
page_id: schema-gelato-selfie-report-921a24a6
path: schemas
description: Result from a selfie check
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_selfie_report

Result from a selfie check

```yaml
{"title": "GelatoSelfieReport", "required": ["status"], "type": "object", "properties": {"document": {"maxLength": 5000, "type": "string", "description": "ID of the [File](https://docs.stripe.com/api/files) holding the image of the identity document used in this check.", "nullable": true}, "error": {"description": "Details on the verification error. Present when status is `unverified`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_selfie_report_error"}]}, "selfie": {"maxLength": 5000, "type": "string", "description": "ID of the [File](https://docs.stripe.com/api/files) holding the image of the selfie used in this check.", "nullable": true}, "status": {"type": "string", "description": "Status of this `selfie` check.", "enum": ["unverified", "verified"], "x-stripeBypassValidation": true}}, "description": "Result from a selfie check", "x-expandableFields": ["error"]}
```
