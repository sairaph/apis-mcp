---
title: forwarded_request_context
page_id: schema-forwarded-request-context-23b09c13
path: schemas
description: Metadata about the forwarded request.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# forwarded_request_context

Metadata about the forwarded request.

```yaml
{"title": "ForwardedRequestContext", "required": ["destination_duration", "destination_ip_address"], "type": "object", "properties": {"destination_duration": {"type": "integer", "description": "The time it took in milliseconds for the destination endpoint to respond."}, "destination_ip_address": {"maxLength": 5000, "type": "string", "description": "The IP address of the destination."}}, "description": "Metadata about the forwarded request.", "x-expandableFields": []}
```
