---
title: quotes_resource_from_quote
page_id: schema-quotes-resource-from-quote-af74ad6f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_from_quote

```yaml
{"title": "QuotesResourceFromQuote", "required": ["is_revision", "quote"], "type": "object", "properties": {"is_revision": {"type": "boolean", "description": "Whether this quote is a revision of a different quote."}, "quote": {"description": "The quote that was cloned.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/quote"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/quote"}]}}}, "description": "", "x-expandableFields": ["quote"]}
```
