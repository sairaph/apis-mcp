---
title: credit_note_refund
page_id: schema-credit-note-refund-03e39107
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# credit_note_refund

```yaml
{"title": "CreditNoteRefund", "required": ["amount_refunded", "refund"], "type": "object", "properties": {"amount_refunded": {"type": "integer", "description": "Amount of the refund that applies to this credit note, in cents (or local equivalent)."}, "payment_record_refund": {"description": "The PaymentRecord refund details associated with this credit note refund.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/credit_notes_payment_record_refund"}]}, "refund": {"description": "ID of the refund.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/refund"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/refund"}]}}, "type": {"type": "string", "description": "Type of the refund, one of `refund` or `payment_record_refund`.", "nullable": true, "enum": ["payment_record_refund", "refund"]}}, "description": "", "x-expandableFields": ["payment_record_refund", "refund"]}
```
