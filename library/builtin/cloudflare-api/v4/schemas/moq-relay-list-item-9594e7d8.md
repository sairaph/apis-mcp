---
title: moq_relay_list_item
page_id: schema-moq-relay-list-item-9594e7d8
path: schemas
description: Abbreviated relay for list responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_relay_list_item

Abbreviated relay for list responses.

```yaml
{"description": "Abbreviated relay for list responses.", "type": "object", "properties": {"created": {"type": "string", "format": "date-time"}, "modified": {"type": "string", "format": "date-time"}, "name": {"type": "string"}, "uid": {"type": "string", "example": "a1b2c3d4e5f67890a1b2c3d4e5f67890"}}, "required": ["uid", "created", "modified", "name"]}
```
