---
title: snippets_SnippetRules
page_id: schema-snippets-snippetrules-becad33a
path: schemas
description: Lists snippet rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# snippets_SnippetRules

Lists snippet rules.

```yaml
{"description": "Lists snippet rules.", "type": "array", "items": {"description": "Define a snippet rule.", "properties": {"description": {"description": "Provide an informative description of the rule.", "type": "string", "example": "Execute my_snippet when IP address is 1.1.1.1.", "default": "", "title": "Description", "x-auditable": true}, "enabled": {"description": "Indicate whether to execute the rule.", "type": "boolean", "example": true, "default": false, "title": "Enabled", "x-auditable": true}, "expression": {"description": "Define the expression that determines which traffic matches the rule.", "type": "string", "example": "ip.src eq 1.1.1.1", "minLength": 1, "title": "Expression", "x-auditable": true}, "id": {"description": "Specify the unique ID of the rule.", "type": "string", "example": "3a03d665bac047339bb530ecb439a90d", "pattern": "^[0-9a-f]{32}$", "readOnly": true, "title": "ID", "x-auditable": true}, "last_updated": {"description": "Specify the timestamp of when the rule was last modified.", "type": "string", "format": "date-time", "example": "2000-01-01T00:00:00.000000Z", "readOnly": true, "title": "Last Updated", "x-auditable": true}, "snippet_name": {"$ref": "#/components/schemas/snippets_SnippetName"}}, "required": ["id", "expression", "snippet_name", "last_updated"], "title": "Snippet Rule", "type": "object"}, "title": "Snippet Rules"}
```
