---
title: rulesets_Response
page_id: schema-rulesets-response-ab29efda
path: schemas
description: A response object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_Response

A response object.

```yaml
{"description": "A response object.", "type": "object", "properties": {"errors": {"$ref": "#/components/schemas/rulesets_Errors"}, "messages": {"$ref": "#/components/schemas/rulesets_Messages"}, "result": {"description": "A result.", "title": "Result"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "title": "Success"}}, "required": ["result", "success", "errors", "messages"], "title": "Response"}
```
