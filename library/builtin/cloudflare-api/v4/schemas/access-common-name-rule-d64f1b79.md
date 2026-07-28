---
title: access_common_name_rule
page_id: schema-access-common-name-rule-d64f1b79
path: schemas
description: Matches a specific common name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_common_name_rule

Matches a specific common name.

```yaml
{"description": "Matches a specific common name.", "type": "object", "properties": {"common_name": {"type": "object", "properties": {"common_name": {"description": "The common name to match.", "type": "string", "example": "james@example.com"}}, "required": ["common_name"]}}, "required": ["common_name"], "title": "Common Name"}
```
