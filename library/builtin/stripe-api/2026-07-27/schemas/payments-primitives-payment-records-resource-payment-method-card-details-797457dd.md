---
title: payments_primitives_payment_records_resource_payment_method_card_details_resource_three_d_secure
page_id: schema-payments-primitives-payment-records-resource-payment-method-card-details-797457dd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_payment_method_card_details_resource_three_d_secure

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCardDetailsResourceThreeDSecure", "type": "object", "properties": {"authentication_flow": {"type": "string", "description": "For authenticated transactions: Indicates how the issuing bank authenticated the customer.", "nullable": true, "enum": ["challenge", "frictionless"]}, "cryptogram": {"maxLength": 5000, "type": "string", "description": "The 3D Secure cryptogram, also known as the \"authentication value\" (AAV, CAVV or AEVV).", "nullable": true}, "electronic_commerce_indicator": {"type": "string", "description": "The Electronic Commerce Indicator (ECI). A protocol-level field indicating what degree of authentication was performed.", "nullable": true, "enum": ["01", "02", "03", "04", "05", "06", "07"]}, "exemption_indicator": {"type": "string", "description": "The exemption requested via 3DS and accepted by the issuer at authentication time.", "nullable": true, "enum": ["low_risk", "none"]}, "exemption_indicator_applied": {"type": "boolean", "description": "Whether Stripe requested the value of `exemption_indicator` in the transaction. This will depend on the outcome of Stripe's internal risk assessment.", "nullable": true}, "result": {"type": "string", "description": "Indicates the outcome of 3D Secure authentication.", "nullable": true, "enum": ["attempt_acknowledged", "authenticated", "exempted", "failed", "not_supported", "processing_error"]}, "result_reason": {"type": "string", "description": "Additional information about why 3D Secure succeeded or failed, based on the `result`.", "nullable": true, "enum": ["abandoned", "bypassed", "canceled", "card_not_enrolled", "network_not_supported", "protocol_error", "rejected"]}, "version": {"type": "string", "description": "The version of 3D Secure that was used.", "nullable": true, "enum": ["1.0.2", "2.1.0", "2.2.0"]}}, "description": "", "x-expandableFields": []}
```
