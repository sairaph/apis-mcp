---
title: v2.error.stripe_loss_liable_cannot_be_deleted
page_id: schema-v2-error-stripe-loss-liable-cannot-be-deleted-b362c33f
path: schemas
description: Account with Stripe-owned loss liability and dashboard cannot be deleted.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.stripe_loss_liable_cannot_be_deleted

Account with Stripe-owned loss liability and dashboard cannot be deleted.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message", "user_message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["stripe_loss_liable_cannot_be_deleted"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}, "user_message": {"type": "string", "description": "A user-friendly message that can be shown to end-users"}}, "description": "Information about the error that occurred"}}, "description": "Account with Stripe-owned loss liability and dashboard cannot be deleted."}
```
