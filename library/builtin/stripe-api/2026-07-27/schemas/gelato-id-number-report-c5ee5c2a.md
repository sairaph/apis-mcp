---
title: gelato_id_number_report
page_id: schema-gelato-id-number-report-c5ee5c2a
path: schemas
description: Result from an id_number check
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_id_number_report

Result from an id_number check

```yaml
{"title": "GelatoIdNumberReport", "required": ["status"], "type": "object", "properties": {"dob": {"description": "Date of birth.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_data_id_number_report_date"}]}, "error": {"description": "Details on the verification error. Present when status is `unverified`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_id_number_report_error"}]}, "first_name": {"maxLength": 5000, "type": "string", "description": "First name.", "nullable": true}, "id_number": {"maxLength": 5000, "type": "string", "description": "ID number. When `id_number_type` is `us_ssn`, only the last 4 digits are present.", "nullable": true}, "id_number_type": {"type": "string", "description": "Type of ID number.", "nullable": true, "enum": ["br_cpf", "sg_nric", "us_ssn"]}, "last_name": {"maxLength": 5000, "type": "string", "description": "Last name.", "nullable": true}, "status": {"type": "string", "description": "Status of this `id_number` check.", "enum": ["unverified", "verified"], "x-stripeBypassValidation": true}}, "description": "Result from an id_number check", "x-expandableFields": ["dob", "error"]}
```
