---
title: snippets_Message
page_id: schema-snippets-message-ec98fc39
path: schemas
description: Describes an API message.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# snippets_Message

Describes an API message.

```yaml
{"description": "Describes an API message.", "type": "object", "properties": {"code": {"description": "Identify the message code.", "type": "integer", "example": 10000, "title": "Code", "x-auditable": true}, "message": {"description": "Describes the message text.", "type": "string", "example": "something bad happened", "minLength": 1, "title": "Description", "x-auditable": true}}, "required": ["message"], "title": "Message"}
```
