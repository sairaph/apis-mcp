---
title: billing.alert
page_id: schema-billing-alert-fe4b7887
path: schemas
description: A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing.alert

A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.

```yaml
{"title": "ThresholdsResourceAlert", "required": ["alert_type", "id", "livemode", "object", "title"], "type": "object", "properties": {"alert_type": {"type": "string", "description": "Defines the type of the alert.", "enum": ["usage_threshold"], "x-stripeBypassValidation": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["billing.alert"]}, "status": {"type": "string", "description": "Status of the alert. This can be active, inactive or archived.", "nullable": true, "enum": ["active", "archived", "inactive"]}, "title": {"maxLength": 5000, "type": "string", "description": "Title of the alert."}, "usage_threshold": {"description": "Encapsulates configuration of the alert to monitor usage on a specific [Billing Meter](https://docs.stripe.com/api/billing/meter).", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/thresholds_resource_usage_threshold_config"}]}}, "description": "A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.", "x-expandableFields": ["usage_threshold"], "x-resourceId": "billing.alert"}
```
