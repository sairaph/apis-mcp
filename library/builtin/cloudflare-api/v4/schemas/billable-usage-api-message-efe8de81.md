---
title: billable-usage-api_message
page_id: schema-billable-usage-api-message-efe8de81
path: schemas
description: Represents an API notice or error detail.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# billable-usage-api_message

Represents an API notice or error detail.

```yaml
{"description": "Represents an API notice or error detail.", "type": "object", "properties": {"code": {"description": "Identifies the error or notice type.", "type": "integer"}, "message": {"description": "Describes the error or notice.", "type": "string"}}, "required": ["message"]}
```
