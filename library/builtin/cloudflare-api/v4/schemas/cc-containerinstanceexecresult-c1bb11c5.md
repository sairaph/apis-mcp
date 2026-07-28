---
title: cc_ContainerInstanceExecResult
page_id: schema-cc-containerinstanceexecresult-c1bb11c5
path: schemas
description: Result of executing a command in a container instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ContainerInstanceExecResult

Result of executing a command in a container instance.

```yaml
{"description": "Result of executing a command in a container instance.", "type": "object", "properties": {"exit_code": {"description": "Exit code returned by the command.", "type": "integer"}, "stderr": {"description": "Base64-encoded standard error from the command.", "type": "string", "format": "byte", "pattern": "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$"}, "stdout": {"description": "Base64-encoded standard output from the command.", "type": "string", "format": "byte", "pattern": "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$"}}, "required": ["stdout", "stderr", "exit_code"]}
```
