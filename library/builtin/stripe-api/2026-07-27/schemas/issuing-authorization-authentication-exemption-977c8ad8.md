---
title: issuing_authorization_authentication_exemption
page_id: schema-issuing-authorization-authentication-exemption-977c8ad8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_authentication_exemption

```yaml
{"title": "IssuingAuthorizationAuthenticationExemption", "required": ["claimed_by", "type"], "type": "object", "properties": {"claimed_by": {"type": "string", "description": "The entity that requested the exemption, either the acquiring merchant or the Issuing user.", "enum": ["acquirer", "issuer"]}, "type": {"type": "string", "description": "The specific exemption claimed for this authorization.", "enum": ["low_value_transaction", "transaction_risk_analysis", "unknown"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
