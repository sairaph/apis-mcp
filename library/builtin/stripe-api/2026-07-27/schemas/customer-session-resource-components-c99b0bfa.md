---
title: customer_session_resource_components
page_id: schema-customer-session-resource-components-c99b0bfa
path: schemas
description: Configuration for the components supported by this Customer Session.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_session_resource_components

Configuration for the components supported by this Customer Session.

```yaml
{"title": "CustomerSessionResourceComponents", "required": ["buy_button", "customer_sheet", "mobile_payment_element", "payment_element", "pricing_table"], "type": "object", "properties": {"buy_button": {"$ref": "#/components/schemas/customer_session_resource_components_resource_buy_button"}, "customer_sheet": {"$ref": "#/components/schemas/customer_session_resource_components_resource_customer_sheet"}, "mobile_payment_element": {"$ref": "#/components/schemas/customer_session_resource_components_resource_mobile_payment_element"}, "payment_element": {"$ref": "#/components/schemas/customer_session_resource_components_resource_payment_element"}, "pricing_table": {"$ref": "#/components/schemas/customer_session_resource_components_resource_pricing_table"}}, "description": "Configuration for the components supported by this Customer Session.", "x-expandableFields": ["buy_button", "customer_sheet", "mobile_payment_element", "payment_element", "pricing_table"]}
```
