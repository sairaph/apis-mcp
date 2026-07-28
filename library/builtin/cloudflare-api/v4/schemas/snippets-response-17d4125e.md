---
title: snippets_Response
page_id: schema-snippets-response-17d4125e
path: schemas
description: Return all API responses using this object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# snippets_Response

Return all API responses using this object.

```yaml
{"description": "Return all API responses using this object.", "type": "object", "properties": {"errors": {"$ref": "#/components/schemas/snippets_Errors"}, "messages": {"$ref": "#/components/schemas/snippets_Messages"}, "result": {"description": "Contain the response result.", "type": "object", "title": "Result"}, "success": {"description": "Indicate whether the API call was successful.", "type": "boolean", "title": "Success", "x-auditable": true}}, "required": ["result", "success", "errors", "messages"], "title": "Response"}
```
