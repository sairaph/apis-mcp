---
title: payment_flows_amount_details
page_id: schema-payment-flows-amount-details-372fda73
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_amount_details

```yaml
{"title": "PaymentFlowsAmountDetails", "type": "object", "properties": {"discount_amount": {"type": "integer", "description": "The total discount applied on the transaction represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). An integer greater than 0.\n\nThis field is mutually exclusive with the `amount_details[line_items][#][discount_amount]` field."}, "error": {"$ref": "#/components/schemas/payment_flows_amount_details_resource_error"}, "line_items": {"title": "PaymentFlowsAmountDetailsResourceLineItemsList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "Details about each object.", "items": {"$ref": "#/components/schemas/payment_intent_amount_details_line_item"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "A list of line items, each containing information about a product in the PaymentIntent. There is a maximum of 200 line items.", "x-expandableFields": ["data"]}, "shipping": {"$ref": "#/components/schemas/payment_flows_amount_details_resource_shipping"}, "tax": {"$ref": "#/components/schemas/payment_flows_amount_details_resource_tax"}, "tip": {"$ref": "#/components/schemas/payment_flows_amount_details_client_resource_tip"}}, "description": "", "x-expandableFields": ["error", "line_items", "shipping", "tax", "tip"]}
```
