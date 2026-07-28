---
title: v2.error.default_outbound_destination_invalid
page_id: schema-v2-error-default-outbound-destination-invalid-8ecc3dc7
path: schemas
description: Outbound Destination ID is invalid.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.default_outbound_destination_invalid

Outbound Destination ID is invalid.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["default_outbound_destination_invalid"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Outbound Destination ID is invalid."}
```
