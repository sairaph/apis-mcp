---
title: customer_session_resource_components_resource_customer_sheet_resource_features
page_id: schema-customer-session-resource-components-resource-customer-sheet-resource-fe-320f7033
path: schemas
description: This hash contains the features the customer sheet supports.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_session_resource_components_resource_customer_sheet_resource_features

This hash contains the features the customer sheet supports.

```yaml
{"title": "CustomerSessionResourceComponentsResourceCustomerSheetResourceFeatures", "type": "object", "properties": {"payment_method_allow_redisplay_filters": {"type": "array", "description": "A list of [`allow_redisplay`](https://docs.stripe.com/api/payment_methods/object#payment_method_object-allow_redisplay) values that controls which saved payment methods the customer sheet displays by filtering to only show payment methods with an `allow_redisplay` value that is present in this list.\n\nIf not specified, defaults to [\"always\"]. In order to display all saved payment methods, specify [\"always\", \"limited\", \"unspecified\"].", "nullable": true, "items": {"type": "string", "enum": ["always", "limited", "unspecified"]}}, "payment_method_remove": {"type": "string", "description": "Controls whether the customer sheet displays the option to remove a saved payment method.\"\n\nAllowing buyers to remove their saved payment methods impacts subscriptions that depend on that payment method. Removing the payment method detaches the [`customer` object](https://docs.stripe.com/api/payment_methods/object#payment_method_object-customer) from that [PaymentMethod](https://docs.stripe.com/api/payment_methods).", "nullable": true, "enum": ["disabled", "enabled"], "x-stripeBypassValidation": true}}, "description": "This hash contains the features the customer sheet supports.", "x-expandableFields": []}
```
