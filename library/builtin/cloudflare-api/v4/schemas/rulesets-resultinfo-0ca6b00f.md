---
title: rulesets_ResultInfo
page_id: schema-rulesets-resultinfo-0ca6b00f
path: schemas
description: Information to navigate the results.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ResultInfo

Information to navigate the results.

```yaml
{"description": "Information to navigate the results.", "type": "object", "properties": {"cursors": {"description": "The set of cursors.", "type": "object", "properties": {"after": {"allOf": [{"$ref": "#/components/schemas/rulesets_Cursor"}, {"title": "After Cursor"}]}}, "required": ["after"], "title": "Cursors"}}, "title": "Result Info"}
```
