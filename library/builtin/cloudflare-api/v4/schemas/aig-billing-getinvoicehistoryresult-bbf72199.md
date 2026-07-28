---
title: aig-billing_GetInvoiceHistoryResult
page_id: schema-aig-billing-getinvoicehistoryresult-bbf72199
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_GetInvoiceHistoryResult

```yaml
{"type": "object", "properties": {"invoices": {"type": "array", "items": {"properties": {"amount_due": {"type": "number"}, "amount_paid": {"type": "number"}, "amount_remaining": {"type": "number"}, "attempt_count": {"type": "number"}, "attempted": {"type": "boolean"}, "auto_advance": {"type": "boolean", "nullable": true}, "created": {"type": "number"}, "created_by": {"type": "string"}, "currency": {"type": "string"}, "description": {"type": "string", "nullable": true}, "id": {"type": "string", "nullable": true}, "invoice_origin": {"type": "string"}, "invoice_pdf": {"type": "string", "nullable": true}, "status": {"type": "string", "nullable": true}}, "required": ["amount_paid", "amount_due", "amount_remaining", "currency"], "type": "object"}}, "pagination": {"type": "object", "properties": {"has_more": {"type": "boolean"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}}, "required": ["has_more", "page", "per_page", "total_count"]}}, "required": ["invoices", "pagination"]}
```
