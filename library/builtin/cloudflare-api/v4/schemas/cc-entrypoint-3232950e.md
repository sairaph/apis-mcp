---
title: cc_Entrypoint
page_id: schema-cc-entrypoint-3232950e
path: schemas
description: |-
    The entry point for the container, specifying the executable to run when the container starts.
    This can be overridden at run-time. If overridden, the default command from the image is ignored.
    Both entrypoint and command can be specified at run-time to completely replace the image defaults.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_Entrypoint

The entry point for the container, specifying the executable to run when the container starts.
This can be overridden at run-time. If overridden, the default command from the image is ignored.
Both entrypoint and command can be specified at run-time to completely replace the image defaults.

```yaml
{"description": "The entry point for the container, specifying the executable to run when the container starts.\nThis can be overridden at run-time. If overridden, the default command from the image is ignored.\nBoth entrypoint and command can be specified at run-time to completely replace the image defaults.\n", "type": "array", "items": {"$ref": "#/components/schemas/cc_ExecFormParam"}, "example": ["/bin/bash"]}
```
