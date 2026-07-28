---
title: payment_flows_private_payment_methods_card_details_api_resource_enterprise_features_overcapture_overcapture
page_id: schema-payment-flows-private-payment-methods-card-details-api-resource-enterpri-4668a49b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_private_payment_methods_card_details_api_resource_enterprise_features_overcapture_overcapture

```yaml
{"title": "PaymentFlowsPrivatePaymentMethodsCardDetailsAPIResourceEnterpriseFeaturesOvercaptureOvercapture", "required": ["maximum_amount_capturable", "status"], "type": "object", "properties": {"maximum_amount_capturable": {"type": "integer", "description": "The maximum amount that can be captured."}, "status": {"type": "string", "description": "Indicates whether or not the authorized amount can be over-captured.", "enum": ["available", "unavailable"]}}, "description": "", "x-expandableFields": []}
```
