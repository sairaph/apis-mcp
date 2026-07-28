---
title: v2.error.duplicate_meter_event
page_id: schema-v2-error-duplicate-meter-event-b106e1b0
path: schemas
description: A meter event with a duplicate identifier has already been submitted.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.duplicate_meter_event

A meter event with a duplicate identifier has already been submitted.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["duplicate_meter_event"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "A meter event with a duplicate identifier has already been submitted."}
```
