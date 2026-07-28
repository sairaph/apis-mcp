---
title: gelato_related_person
page_id: schema-gelato-related-person-b624dfd7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_related_person

```yaml
{"title": "GelatoRelatedPerson", "required": ["account", "person"], "type": "object", "properties": {"account": {"maxLength": 5000, "type": "string", "description": "Token referencing the associated Account of the related Person resource."}, "person": {"maxLength": 5000, "type": "string", "description": "Token referencing the related Person resource."}}, "description": "", "x-expandableFields": []}
```
