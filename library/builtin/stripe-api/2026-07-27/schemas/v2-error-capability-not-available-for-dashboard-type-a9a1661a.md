---
title: v2.error.capability_not_available_for_dashboard_type
page_id: schema-v2-error-capability-not-available-for-dashboard-type-a9a1661a
path: schemas
description: Feature cannot be requested for the dashboard type.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.capability_not_available_for_dashboard_type

Feature cannot be requested for the dashboard type.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["capability_not_available_for_dashboard_type"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Feature cannot be requested for the dashboard type."}
```
