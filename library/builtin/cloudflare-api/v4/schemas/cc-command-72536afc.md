---
title: cc_Command
page_id: schema-cc-command-72536afc
path: schemas
description: |-
    The command to be executed when the container starts, passed to the entrypoint.
    This can be overridden at run-time. If only the command is overridden at run-time,
    it gets passed to the default entrypoint specified in the image.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_Command

The command to be executed when the container starts, passed to the entrypoint.
This can be overridden at run-time. If only the command is overridden at run-time,
it gets passed to the default entrypoint specified in the image.

```yaml
{"description": "The command to be executed when the container starts, passed to the entrypoint.\nThis can be overridden at run-time. If only the command is overridden at run-time,\nit gets passed to the default entrypoint specified in the image.\n", "type": "array", "items": {"$ref": "#/components/schemas/cc_ExecFormParam"}, "example": ["myapp", "--default-option"]}
```
