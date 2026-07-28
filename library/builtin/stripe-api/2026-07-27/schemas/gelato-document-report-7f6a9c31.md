---
title: gelato_document_report
page_id: schema-gelato-document-report-7f6a9c31
path: schemas
description: Result from a document check
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_document_report

Result from a document check

```yaml
{"title": "GelatoDocumentReport", "required": ["status"], "type": "object", "properties": {"address": {"description": "Address as it appears in the document.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "dob": {"description": "Date of birth as it appears in the document.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_data_document_report_date_of_birth"}]}, "error": {"description": "Details on the verification error. Present when status is `unverified`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_document_report_error"}]}, "expiration_date": {"description": "Expiration date of the document.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_data_document_report_expiration_date"}]}, "files": {"type": "array", "description": "Array of [File](https://docs.stripe.com/api/files) ids containing images for this document.", "nullable": true, "items": {"maxLength": 5000, "type": "string"}}, "first_name": {"maxLength": 5000, "type": "string", "description": "First name as it appears in the document.", "nullable": true}, "issued_date": {"description": "Issued date of the document.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_data_document_report_issued_date"}]}, "issuing_country": {"maxLength": 5000, "type": "string", "description": "Issuing country of the document.", "nullable": true}, "last_name": {"maxLength": 5000, "type": "string", "description": "Last name as it appears in the document.", "nullable": true}, "number": {"maxLength": 5000, "type": "string", "description": "Document ID number.", "nullable": true}, "sex": {"type": "string", "description": "Sex of the person in the document.", "nullable": true, "enum": ["[redacted]", "female", "male", "unknown"]}, "status": {"type": "string", "description": "Status of this `document` check.", "enum": ["unverified", "verified"], "x-stripeBypassValidation": true}, "type": {"type": "string", "description": "Type of the document.", "nullable": true, "enum": ["driving_license", "id_card", "passport"]}, "unparsed_place_of_birth": {"maxLength": 5000, "type": "string", "description": "Place of birth as it appears in the document.", "nullable": true}, "unparsed_sex": {"maxLength": 5000, "type": "string", "description": "Sex as it appears in the document.", "nullable": true}}, "description": "Result from a document check", "x-expandableFields": ["address", "dob", "error", "expiration_date", "issued_date"]}
```
