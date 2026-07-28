---
title: customer_session_resource_components_resource_payment_element
page_id: schema-customer-session-resource-components-resource-payment-element-3a63d214
path: schemas
description: This hash contains whether the Payment Element is enabled and the features it supports.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_session_resource_components_resource_payment_element

This hash contains whether the Payment Element is enabled and the features it supports.

```yaml
{"title": "CustomerSessionResourceComponentsResourcePaymentElement", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Whether the Payment Element is enabled."}, "features": {"description": "This hash defines whether the Payment Element supports certain features.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/customer_session_resource_components_resource_payment_element_resource_features"}]}}, "description": "This hash contains whether the Payment Element is enabled and the features it supports.", "x-expandableFields": ["features"]}
```
