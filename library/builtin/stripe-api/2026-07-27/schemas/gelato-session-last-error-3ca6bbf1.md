---
title: gelato_session_last_error
page_id: schema-gelato-session-last-error-3ca6bbf1
path: schemas
description: Shows last VerificationSession error
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_session_last_error

Shows last VerificationSession error

```yaml
{"title": "GelatoSessionLastError", "type": "object", "properties": {"code": {"type": "string", "description": "A short machine-readable string giving the reason for the verification or user-session failure.", "nullable": true, "enum": ["abandoned", "consent_declined", "country_not_supported", "device_not_supported", "document_expired", "document_type_not_supported", "document_unverified_other", "email_unverified_other", "email_verification_declined", "id_number_insufficient_document_data", "id_number_mismatch", "id_number_unverified_other", "phone_unverified_other", "phone_verification_declined", "selfie_document_missing_photo", "selfie_face_mismatch", "selfie_manipulated", "selfie_unverified_other", "under_supported_age"], "x-stripeBypassValidation": true}, "reason": {"maxLength": 5000, "type": "string", "description": "A message that explains the reason for verification or user-session failure.", "nullable": true}}, "description": "Shows last VerificationSession error", "x-expandableFields": []}
```
