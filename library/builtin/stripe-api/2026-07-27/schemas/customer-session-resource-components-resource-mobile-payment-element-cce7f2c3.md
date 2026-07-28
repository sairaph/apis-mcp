---
title: customer_session_resource_components_resource_mobile_payment_element
page_id: schema-customer-session-resource-components-resource-mobile-payment-element-cce7f2c3
path: schemas
description: This hash contains whether the mobile payment element is enabled and the features it supports.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_session_resource_components_resource_mobile_payment_element

This hash contains whether the mobile payment element is enabled and the features it supports.

```yaml
{"title": "CustomerSessionResourceComponentsResourceMobilePaymentElement", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Whether the mobile payment element is enabled."}, "features": {"description": "This hash defines whether the mobile payment element supports certain features.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/customer_session_resource_components_resource_mobile_payment_element_resource_features"}]}}, "description": "This hash contains whether the mobile payment element is enabled and the features it supports.", "x-expandableFields": ["features"]}
```
