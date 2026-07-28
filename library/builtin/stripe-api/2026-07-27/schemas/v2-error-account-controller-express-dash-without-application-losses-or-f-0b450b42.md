---
title: v2.error.account_controller_express_dash_without_application_losses_or_fees
page_id: schema-v2-error-account-controller-express-dash-without-application-losses-or-f-0b450b42
path: schemas
description: If `dashboard` is `express`, `fees_collector` must be `application` and `losses_collector` must be `application`.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.account_controller_express_dash_without_application_losses_or_fees

If `dashboard` is `express`, `fees_collector` must be `application` and `losses_collector` must be `application`.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["account_controller_express_dash_without_application_losses_or_fees"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "If `dashboard` is `express`, `fees_collector` must be `application` and `losses_collector` must be `application`."}
```
