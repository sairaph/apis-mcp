---
title: payment_flows_amount_details_resource_line_items_list_resource_line_item_resource_payment_method_options
page_id: schema-payment-flows-amount-details-resource-line-items-list-resource-line-item-38cb1bc3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_amount_details_resource_line_items_list_resource_line_item_resource_payment_method_options

```yaml
{"title": "PaymentFlowsAmountDetailsResourceLineItemsListResourceLineItemResourcePaymentMethodOptions", "type": "object", "properties": {"card": {"$ref": "#/components/schemas/payment_flows_private_payment_methods_card_payment_intent_amount_details_line_item_payment_method_options"}, "card_present": {"$ref": "#/components/schemas/payment_flows_private_payment_methods_card_present_amount_details_line_item_payment_method_options"}, "klarna": {"$ref": "#/components/schemas/payment_flows_private_payment_methods_klarna_payment_intent_amount_details_line_item_payment_method_options"}, "paypal": {"$ref": "#/components/schemas/payment_flows_private_payment_methods_paypal_amount_details_line_item_payment_method_options"}}, "description": "", "x-expandableFields": ["card", "card_present", "klarna", "paypal"]}
```
