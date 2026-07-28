---
title: aig-billing_GetInvoicePreviewResult
page_id: schema-aig-billing-getinvoicepreviewresult-f88c2908
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_GetInvoicePreviewResult

```yaml
{"type": "object", "properties": {"amount_due": {"type": "number"}, "amount_paid": {"type": "number"}, "amount_remaining": {"type": "number"}, "currency": {"type": "string"}, "id": {"type": "string"}, "invoice_lines": {"type": "array", "items": {"properties": {"amount": {"type": "number"}, "currency": {"type": "string"}, "description": {"type": "string", "nullable": true}, "period": {"type": "object", "properties": {"end": {"type": "number"}, "start": {"type": "number"}}, "required": ["start", "end"]}, "pretax_credit_amounts": {"type": "array", "items": {"properties": {"amount": {"type": "number"}, "credit_balance_transaction": {"type": "string", "nullable": true}, "discount": {"type": "string", "nullable": true}, "type": {"type": "string"}}, "required": ["amount", "type"], "type": "object"}}, "pricing": {"type": "object", "properties": {"unit_amount_decimal": {"type": "string", "nullable": true}}, "required": ["unit_amount_decimal"]}, "quantity": {"type": "number"}}, "required": ["description", "amount", "period", "currency", "pricing", "quantity"], "type": "object"}}, "period_end": {"type": "number"}, "period_start": {"type": "number"}, "status": {"type": "string", "enum": ["draft", "open", "paid", "uncollectible", "void"]}}, "required": ["id", "amount_due", "amount_paid", "amount_remaining", "currency", "period_start", "period_end", "status", "invoice_lines"]}
```
