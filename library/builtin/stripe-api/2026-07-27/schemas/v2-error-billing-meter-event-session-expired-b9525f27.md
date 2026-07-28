---
title: v2.error.billing_meter_event_session_expired
page_id: schema-v2-error-billing-meter-event-session-expired-b9525f27
path: schemas
description: The temporary session token has expired.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.billing_meter_event_session_expired

The temporary session token has expired.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message", "type"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["billing_meter_event_session_expired"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "type": {"type": "string", "description": "The type of error returned", "enum": ["temporary_session_expired"]}}, "description": "Information about the error that occurred"}}, "description": "The temporary session token has expired."}
```
