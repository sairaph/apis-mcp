---
title: v2.error.no_meter
page_id: schema-v2-error-no-meter-1157cb41
path: schemas
description: A meter must exist to submit events.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.no_meter

A meter must exist to submit events.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["no_meter"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "A meter must exist to submit events."}
```
