---
title: gelato_report_document_options
page_id: schema-gelato-report-document-options-2adaab28
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_report_document_options

```yaml
{"title": "GelatoReportDocumentOptions", "type": "object", "properties": {"allowed_types": {"type": "array", "description": "Array of strings of allowed identity document types. If the provided identity document isn’t one of the allowed types, the verification check will fail with a document_type_not_allowed error code.", "items": {"type": "string", "enum": ["driving_license", "id_card", "passport"]}}, "require_id_number": {"type": "boolean", "description": "Collect an ID number and perform an [ID number check](https://docs.stripe.com/identity/verification-checks?type=id-number) with the document’s extracted name and date of birth."}, "require_live_capture": {"type": "boolean", "description": "Disable image uploads, identity document images have to be captured using the device’s camera."}, "require_matching_selfie": {"type": "boolean", "description": "Capture a face image and perform a [selfie check](https://docs.stripe.com/identity/verification-checks?type=selfie) comparing a photo ID and a picture of your user’s face. [Learn more](https://docs.stripe.com/identity/selfie)."}}, "description": "", "x-expandableFields": []}
```
