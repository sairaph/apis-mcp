---
title: legal_entity_person_verification
page_id: schema-legal-entity-person-verification-e757fc61
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# legal_entity_person_verification

```yaml
{"title": "LegalEntityPersonVerification", "required": ["status"], "type": "object", "properties": {"additional_document": {"description": "A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/legal_entity_person_verification_document"}]}, "details": {"maxLength": 5000, "type": "string", "description": "A user-displayable string describing the verification state for the person. For example, this may say \"Provided identity information could not be verified\".", "nullable": true}, "details_code": {"maxLength": 5000, "type": "string", "description": "One of `document_address_mismatch`, `document_dob_mismatch`, `document_duplicate_type`, `document_id_number_mismatch`, `document_name_mismatch`, `document_nationality_mismatch`, `failed_keyed_identity`, or `failed_other`. A machine-readable code specifying the verification state for the person.", "nullable": true}, "document": {"$ref": "#/components/schemas/legal_entity_person_verification_document"}, "status": {"maxLength": 5000, "type": "string", "description": "The state of verification for the person. Possible values are `unverified`, `pending`, or `verified`. Please refer [guide](https://docs.stripe.com/connect/handling-api-verification) to handle verification updates."}}, "description": "", "x-expandableFields": ["additional_document", "document"]}
```
