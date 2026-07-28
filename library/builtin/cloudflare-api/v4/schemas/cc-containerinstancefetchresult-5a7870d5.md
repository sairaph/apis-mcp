---
title: cc_ContainerInstanceFetchResult
page_id: schema-cc-containerinstancefetchresult-5a7870d5
path: schemas
description: Response returned by the container instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ContainerInstanceFetchResult

Response returned by the container instance.

```yaml
{"description": "Response returned by the container instance.", "type": "object", "properties": {"body": {"description": "Text response body returned by the container.", "type": "string"}, "headers": {"description": "HTTP response headers returned by the container.", "type": "object", "additionalProperties": {"type": "string"}}, "status": {"description": "HTTP status code returned by the container.", "type": "integer", "maximum": 599, "minimum": 100}}, "required": ["status", "headers", "body"]}
```
