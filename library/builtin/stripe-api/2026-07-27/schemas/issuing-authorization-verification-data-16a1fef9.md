---
title: issuing_authorization_verification_data
page_id: schema-issuing-authorization-verification-data-16a1fef9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_verification_data

```yaml
{"title": "IssuingAuthorizationVerificationData", "required": ["address_line1_check", "address_postal_code_check", "cvc_check", "expiry_check"], "type": "object", "properties": {"address_line1_check": {"type": "string", "description": "Whether the cardholder provided an address first line and if it matched the cardholder’s `billing.address.line1`.", "enum": ["match", "mismatch", "not_provided"]}, "address_postal_code_check": {"type": "string", "description": "Whether the cardholder provided a postal code and if it matched the cardholder’s `billing.address.postal_code`.", "enum": ["match", "mismatch", "not_provided"]}, "authentication_exemption": {"description": "The exemption applied to this authorization.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_authorization_authentication_exemption"}]}, "cvc_check": {"type": "string", "description": "Whether the cardholder provided a CVC and if it matched Stripe’s record.", "enum": ["match", "mismatch", "not_provided"]}, "expiry_check": {"type": "string", "description": "Whether the cardholder provided an expiry date and if it matched Stripe’s record.", "enum": ["match", "mismatch", "not_provided"]}, "postal_code": {"maxLength": 5000, "type": "string", "description": "The postal code submitted as part of the authorization used for postal code verification.", "nullable": true}, "three_d_secure": {"description": "3D Secure details.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_authorization_three_d_secure"}]}}, "description": "", "x-expandableFields": ["authentication_exemption", "three_d_secure"]}
```
