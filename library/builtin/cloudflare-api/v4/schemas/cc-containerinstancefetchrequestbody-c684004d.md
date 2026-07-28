---
title: cc_ContainerInstanceFetchRequestBody
page_id: schema-cc-containerinstancefetchrequestbody-c684004d
path: schemas
description: Request body for proxying a request to a container instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ContainerInstanceFetchRequestBody

Request body for proxying a request to a container instance.

```yaml
{"description": "Request body for proxying a request to a container instance.", "type": "object", "properties": {"body": {"description": "Request body to forward to the container.", "type": "string"}, "headers": {"description": "HTTP headers to forward to the container.", "type": "object", "additionalProperties": {"type": "string"}}, "method": {"description": "HTTP method to use. Defaults to GET.", "type": "string", "default": "GET", "enum": ["GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"]}, "url": {"description": "The path to forward to the container (e.g. \"/api/data\").", "type": "string"}}, "required": ["url"]}
```
