---
title: rulesets_Message
page_id: schema-rulesets-message-9bb3c8da
path: schemas
description: A message.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_Message

A message.

```yaml
{"description": "A message.", "type": "object", "properties": {"code": {"description": "A unique code for this message.", "type": "integer", "example": 10000, "title": "Code"}, "message": {"description": "A text description of this message.", "type": "string", "example": "something bad happened", "minLength": 1, "title": "Description"}, "source": {"description": "The source of this message.", "type": "object", "properties": {"pointer": {"description": "A JSON pointer to the field that is the source of the message.", "type": "string", "example": "/rules/0/action", "minLength": 1, "title": "Pointer"}}, "required": ["pointer"], "title": "Source"}}, "required": ["message"], "title": "Message"}
```
