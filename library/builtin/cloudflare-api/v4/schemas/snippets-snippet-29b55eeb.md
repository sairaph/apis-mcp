---
title: snippets_Snippet
page_id: schema-snippets-snippet-29b55eeb
path: schemas
description: Define a snippet.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# snippets_Snippet

Define a snippet.

```yaml
{"description": "Define a snippet.", "type": "object", "properties": {"created_on": {"description": "Indicates when the snippet was created.", "type": "string", "format": "date-time", "example": "2000-01-01T00:00:00.000000Z", "title": "Created On", "x-auditable": true}, "modified_on": {"description": "Indicates when the snippet was last modified.", "type": "string", "format": "date-time", "example": "2000-01-01T00:00:00.000000Z", "title": "Modified On", "x-auditable": true}, "snippet_name": {"$ref": "#/components/schemas/snippets_SnippetName"}}, "required": ["created_on", "modified_one", "snippet_name"], "title": "Snippet"}
```
