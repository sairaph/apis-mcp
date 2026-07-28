---
title: three_d_secure_details_charge
page_id: schema-three-d-secure-details-charge-974a6d5c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# three_d_secure_details_charge

```yaml
{"title": "three_d_secure_details_charge", "type": "object", "properties": {"authentication_flow": {"type": "string", "description": "For authenticated transactions: how the customer was authenticated by\nthe issuing bank.", "nullable": true, "enum": ["challenge", "frictionless"]}, "electronic_commerce_indicator": {"type": "string", "description": "The Electronic Commerce Indicator (ECI). A protocol-level field\nindicating what degree of authentication was performed.", "nullable": true, "enum": ["01", "02", "05", "06", "07"], "x-stripeBypassValidation": true}, "exemption_indicator": {"type": "string", "description": "The exemption requested via 3DS and accepted by the issuer at authentication time.", "nullable": true, "enum": ["low_risk", "none"]}, "exemption_indicator_applied": {"type": "boolean", "description": "Whether Stripe requested the value of `exemption_indicator` in the transaction. This will depend on\nthe outcome of Stripe's internal risk assessment."}, "result": {"type": "string", "description": "Indicates the outcome of 3D Secure authentication.", "nullable": true, "enum": ["attempt_acknowledged", "authenticated", "exempted", "failed", "not_supported", "processing_error"]}, "result_reason": {"type": "string", "description": "Additional information about why 3D Secure succeeded or failed based\non the `result`.", "nullable": true, "enum": ["abandoned", "bypassed", "canceled", "card_not_enrolled", "network_not_supported", "protocol_error", "rejected"]}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "The 3D Secure 1 XID or 3D Secure 2 Directory Server Transaction ID\n(dsTransId) for this payment.", "nullable": true}, "version": {"type": "string", "description": "The version of 3D Secure that was used.", "nullable": true, "enum": ["1.0.2", "2.1.0", "2.2.0", "2.3.0", "2.3.1"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
