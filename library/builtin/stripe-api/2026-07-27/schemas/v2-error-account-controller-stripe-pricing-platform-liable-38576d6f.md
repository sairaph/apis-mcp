---
title: v2.error.account_controller_stripe_pricing_platform_liable
page_id: schema-v2-error-account-controller-stripe-pricing-platform-liable-38576d6f
path: schemas
description: If `losses_collector` is `application`, `fees_collector` must also be `application`.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.account_controller_stripe_pricing_platform_liable

If `losses_collector` is `application`, `fees_collector` must also be `application`.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["account_controller_stripe_pricing_platform_liable"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "If `losses_collector` is `application`, `fees_collector` must also be `application`."}
```
