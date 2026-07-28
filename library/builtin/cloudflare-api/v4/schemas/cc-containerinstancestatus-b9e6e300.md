---
title: cc_ContainerInstanceStatus
page_id: schema-cc-containerinstancestatus-b9e6e300
path: schemas
description: The current status of a container instance.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ContainerInstanceStatus

The current status of a container instance.

```yaml
{"description": "The current status of a container instance.", "type": "object", "properties": {"exit_code": {"description": "The exit code of the container, if status is stopped_with_code.", "type": "integer"}, "last_change": {"description": "Epoch timestamp (milliseconds) of the last status change.", "type": "number"}, "status": {"description": "The current lifecycle status of the container.", "type": "string", "enum": ["starting", "running", "healthy", "stopping", "stopped", "stopped_with_code"]}}, "required": ["status", "last_change"]}
```
