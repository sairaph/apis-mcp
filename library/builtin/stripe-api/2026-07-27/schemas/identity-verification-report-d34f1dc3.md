---
title: identity.verification_report
page_id: schema-identity-verification-report-d34f1dc3
path: schemas
description: |-
    A VerificationReport is the result of an attempt to collect and verify data from a user.
    The collection of verification checks performed is determined from the `type` and `options`
    parameters used. You can find the result of each verification check performed in the
    appropriate sub-resource: `document`, `id_number`, `selfie`.

    Each VerificationReport contains a copy of any data collected by the user as well as
    reference IDs which can be used to access collected images through the [FileUpload](https://docs.stripe.com/api/files)
    API. To configure and create VerificationReports, use the
    [VerificationSession](https://docs.stripe.com/api/identity/verification_sessions) API.

    Related guide: [Accessing verification results](https://docs.stripe.com/identity/verification-sessions#results).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# identity.verification_report

A VerificationReport is the result of an attempt to collect and verify data from a user.
The collection of verification checks performed is determined from the `type` and `options`
parameters used. You can find the result of each verification check performed in the
appropriate sub-resource: `document`, `id_number`, `selfie`.

Each VerificationReport contains a copy of any data collected by the user as well as
reference IDs which can be used to access collected images through the [FileUpload](https://docs.stripe.com/api/files)
API. To configure and create VerificationReports, use the
[VerificationSession](https://docs.stripe.com/api/identity/verification_sessions) API.

Related guide: [Accessing verification results](https://docs.stripe.com/identity/verification-sessions#results).

```yaml
{"title": "GelatoVerificationReport", "required": ["created", "id", "livemode", "object", "type"], "type": "object", "properties": {"client_reference_id": {"maxLength": 5000, "type": "string", "description": "A string to reference this user. This can be a customer ID, a session ID, or similar, and can be used to reconcile this verification with your internal systems.", "nullable": true}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "document": {"$ref": "#/components/schemas/gelato_document_report"}, "email": {"$ref": "#/components/schemas/gelato_email_report"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "id_number": {"$ref": "#/components/schemas/gelato_id_number_report"}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["identity.verification_report"]}, "options": {"$ref": "#/components/schemas/gelato_verification_report_options"}, "phone": {"$ref": "#/components/schemas/gelato_phone_report"}, "selfie": {"$ref": "#/components/schemas/gelato_selfie_report"}, "type": {"type": "string", "description": "Type of report.", "enum": ["document", "id_number", "verification_flow"], "x-stripeBypassValidation": true}, "verification_flow": {"maxLength": 5000, "type": "string", "description": "The configuration token of a verification flow from the dashboard."}, "verification_session": {"maxLength": 5000, "type": "string", "description": "ID of the VerificationSession that created this report.", "nullable": true}}, "description": "A VerificationReport is the result of an attempt to collect and verify data from a user.\nThe collection of verification checks performed is determined from the `type` and `options`\nparameters used. You can find the result of each verification check performed in the\nappropriate sub-resource: `document`, `id_number`, `selfie`.\n\nEach VerificationReport contains a copy of any data collected by the user as well as\nreference IDs which can be used to access collected images through the [FileUpload](https://docs.stripe.com/api/files)\nAPI. To configure and create VerificationReports, use the\n[VerificationSession](https://docs.stripe.com/api/identity/verification_sessions) API.\n\nRelated guide: [Accessing verification results](https://docs.stripe.com/identity/verification-sessions#results).", "x-expandableFields": ["document", "email", "id_number", "options", "phone", "selfie"], "x-resourceId": "identity.verification_report"}
```
