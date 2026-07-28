---
title: aig-billing_GetCreditBalanceResult
page_id: schema-aig-billing-getcreditbalanceresult-101622cd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_GetCreditBalanceResult

```yaml
{"type": "object", "properties": {"balance": {"type": "number"}, "first_topup_success": {"type": "boolean"}, "has_default_payment_method": {"type": "boolean"}, "payment_method": {"type": "object", "nullable": true, "properties": {"brand": {"type": "string"}, "last4": {"type": "string"}}}, "topup_config": {"type": "object", "properties": {"amount": {"type": "number", "nullable": true}, "disabledReason": {"type": "string", "nullable": true}, "error": {"type": "string", "nullable": true}, "lastFailedAt": {"type": "number", "nullable": true}, "threshold": {"type": "number", "nullable": true}}, "required": ["threshold", "amount", "error", "disabledReason", "lastFailedAt"]}}, "required": ["balance", "has_default_payment_method", "payment_method", "topup_config"]}
```
