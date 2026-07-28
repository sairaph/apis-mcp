---
title: billing_meter_resource_customer_mapping_settings
page_id: schema-billing-meter-resource-customer-mapping-settings-c4be4c8d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_meter_resource_customer_mapping_settings

```yaml
{"title": "BillingMeterResourceCustomerMappingSettings", "required": ["event_payload_key", "type"], "type": "object", "properties": {"event_payload_key": {"maxLength": 5000, "type": "string", "description": "The key in the meter event payload to use for mapping the event to a customer."}, "type": {"type": "string", "description": "The method for mapping a meter event to a customer.", "enum": ["by_id"]}}, "description": "", "x-expandableFields": []}
```
