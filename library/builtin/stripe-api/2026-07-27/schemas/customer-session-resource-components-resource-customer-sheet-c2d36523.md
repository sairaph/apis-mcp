---
title: customer_session_resource_components_resource_customer_sheet
page_id: schema-customer-session-resource-components-resource-customer-sheet-c2d36523
path: schemas
description: This hash contains whether the customer sheet is enabled and the features it supports.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_session_resource_components_resource_customer_sheet

This hash contains whether the customer sheet is enabled and the features it supports.

```yaml
{"title": "CustomerSessionResourceComponentsResourceCustomerSheet", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Whether the customer sheet is enabled."}, "features": {"description": "This hash defines whether the customer sheet supports certain features.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/customer_session_resource_components_resource_customer_sheet_resource_features"}]}}, "description": "This hash contains whether the customer sheet is enabled and the features it supports.", "x-expandableFields": ["features"]}
```
