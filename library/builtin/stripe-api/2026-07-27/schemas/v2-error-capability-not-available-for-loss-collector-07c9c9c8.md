---
title: v2.error.capability_not_available_for_loss_collector
page_id: schema-v2-error-capability-not-available-for-loss-collector-07c9c9c8
path: schemas
description: A v2 Account cannot have both the specified capability and Stripe-owned loss liability.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.capability_not_available_for_loss_collector

A v2 Account cannot have both the specified capability and Stripe-owned loss liability.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["capability_not_available_for_loss_collector"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "A v2 Account cannot have both the specified capability and Stripe-owned loss liability."}
```
