---
title: quotes_resource_upfront
page_id: schema-quotes-resource-upfront-d010a451
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_upfront

```yaml
{"title": "QuotesResourceUpfront", "required": ["amount_subtotal", "amount_total", "total_details"], "type": "object", "properties": {"amount_subtotal": {"type": "integer", "description": "Total before any discounts or taxes are applied."}, "amount_total": {"type": "integer", "description": "Total after discounts and taxes are applied."}, "line_items": {"title": "QuotesResourceListLineItems", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "Details about each object.", "items": {"$ref": "#/components/schemas/item"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "The line items that will appear on the next invoice after this quote is accepted. This does not include pending invoice items that exist on the customer but may still be included in the next invoice.", "x-expandableFields": ["data"]}, "total_details": {"$ref": "#/components/schemas/quotes_resource_total_details"}}, "description": "", "x-expandableFields": ["line_items", "total_details"]}
```
