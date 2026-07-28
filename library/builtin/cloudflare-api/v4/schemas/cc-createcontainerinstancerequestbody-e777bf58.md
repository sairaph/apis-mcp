---
title: cc_CreateContainerInstanceRequestBody
page_id: schema-cc-createcontainerinstancerequestbody-e777bf58
path: schemas
description: Request body for creating a new container instance within an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_CreateContainerInstanceRequestBody

Request body for creating a new container instance within an application.

```yaml
{"description": "Request body for creating a new container instance within an application.", "type": "object", "properties": {"enable_internet": {"description": "Whether to enable outbound internet access for the container.", "type": "boolean", "default": true}, "entrypoint": {"description": "Command to run in the container. Overrides the image's default entrypoint.", "type": "array", "items": {"type": "string"}}, "environment_variables": {"description": "Environment variables to pass to the container.", "type": "object", "additionalProperties": {"type": "string"}}, "name": {"description": "A human-readable name for the instance. Used as the Durable Object name via getByName().", "type": "string", "maxLength": 253, "minLength": 1, "pattern": "^[a-z0-9]([a-z0-9.-]*[a-z0-9])?$"}}, "required": ["name"]}
```
