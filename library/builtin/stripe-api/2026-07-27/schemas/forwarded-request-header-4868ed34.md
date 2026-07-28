---
title: forwarded_request_header
page_id: schema-forwarded-request-header-4868ed34
path: schemas
description: Header data.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# forwarded_request_header

Header data.

```yaml
{"title": "ForwardedRequestHeader", "required": ["name", "value"], "type": "object", "properties": {"name": {"maxLength": 5000, "type": "string", "description": "The header name."}, "value": {"maxLength": 5000, "type": "string", "description": "The header value."}}, "description": "Header data.", "x-expandableFields": []}
```
