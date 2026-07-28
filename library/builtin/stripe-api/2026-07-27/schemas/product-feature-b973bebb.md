---
title: product_feature
page_id: schema-product-feature-b973bebb
path: schemas
description: |-
    A product_feature represents an attachment between a feature and a product.
    When a product is purchased that has a feature attached, Stripe will create an entitlement to the feature for the purchasing customer.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# product_feature

A product_feature represents an attachment between a feature and a product.
When a product is purchased that has a feature attached, Stripe will create an entitlement to the feature for the purchasing customer.

```yaml
{"title": "ProductFeature", "required": ["entitlement_feature", "id", "livemode", "object"], "type": "object", "properties": {"entitlement_feature": {"$ref": "#/components/schemas/entitlements.feature"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["product_feature"]}}, "description": "A product_feature represents an attachment between a feature and a product.\nWhen a product is purchased that has a feature attached, Stripe will create an entitlement to the feature for the purchasing customer.", "x-expandableFields": ["entitlement_feature"], "x-resourceId": "product_feature"}
```
