---
title: cc_ContainerInstanceExecRequestBody
page_id: schema-cc-containerinstanceexecrequestbody-ea111021
path: schemas
description: Request body for executing a command in a running container instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ContainerInstanceExecRequestBody

Request body for executing a command in a running container instance.

```yaml
{"description": "Request body for executing a command in a running container instance.", "type": "object", "properties": {"command": {"description": "Command and arguments to execute without shell interpretation.", "type": "array", "items": {"type": "string"}, "minItems": 1}, "cwd": {"description": "Working directory for the command.", "type": "string"}, "environment_variables": {"description": "Environment variables to set for the command.", "type": "object", "additionalProperties": {"type": "string"}}, "stdin": {"description": "Base64-encoded bytes to provide to the command on stdin.", "type": "string", "format": "byte", "pattern": "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$"}, "user": {"description": "Container user to execute the command as.", "type": "string"}}, "required": ["command"]}
```
