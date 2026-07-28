---
title: treasury_outbound_payments_resource_outbound_payment_resource_end_user_details
page_id: schema-treasury-outbound-payments-resource-outbound-payment-resource-end-user-d-c68058d4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_outbound_payments_resource_outbound_payment_resource_end_user_details

```yaml
{"title": "TreasuryOutboundPaymentsResourceOutboundPaymentResourceEndUserDetails", "required": ["present"], "type": "object", "properties": {"ip_address": {"maxLength": 5000, "type": "string", "description": "IP address of the user initiating the OutboundPayment. Set if `present` is set to `true`. IP address collection is required for risk and compliance reasons. This will be used to help determine if the OutboundPayment is authorized or should be blocked.", "nullable": true}, "present": {"type": "boolean", "description": "`true` if the OutboundPayment creation request is being made on behalf of an end user by a platform. Otherwise, `false`."}}, "description": "", "x-expandableFields": []}
```
