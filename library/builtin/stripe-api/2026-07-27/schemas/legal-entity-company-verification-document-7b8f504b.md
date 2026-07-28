---
title: legal_entity_company_verification_document
page_id: schema-legal-entity-company-verification-document-7b8f504b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# legal_entity_company_verification_document

```yaml
{"title": "LegalEntityCompanyVerificationDocument", "type": "object", "properties": {"back": {"description": "The back of a document returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `additional_verification`. Note that `additional_verification` files are [not downloadable](/file-upload#uploading-a-file).", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "details": {"maxLength": 5000, "type": "string", "description": "A user-displayable string describing the verification state of this document.", "nullable": true}, "details_code": {"maxLength": 5000, "type": "string", "description": "One of `document_corrupt`, `document_expired`, `document_failed_copy`, `document_failed_greyscale`, `document_failed_other`, `document_failed_test_mode`, `document_fraudulent`, `document_incomplete`, `document_invalid`, `document_manipulated`, `document_not_readable`, `document_not_uploaded`, `document_type_not_supported`, or `document_too_large`. A machine-readable code specifying the verification state for this document.", "nullable": true}, "front": {"description": "The front of a document returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `additional_verification`. Note that `additional_verification` files are [not downloadable](/file-upload#uploading-a-file).", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}}, "description": "", "x-expandableFields": ["back", "front"]}
```
