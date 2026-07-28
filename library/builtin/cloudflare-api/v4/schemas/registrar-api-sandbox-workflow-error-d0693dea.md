---
title: registrar-api-sandbox_workflow_error
page_id: schema-registrar-api-sandbox-workflow-error-d0693dea
path: schemas
description: |-
    Error details when a workflow reaches the `failed` state. The specific
    error codes and messages depend on the workflow type (registration,
    update, etc.) and the underlying registry response. These workflow
    error codes are separate from immediate HTTP error `errors[].code`
    values returned by non-2xx responses. Surface
    `error.message` to the user for context.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_workflow_error

Error details when a workflow reaches the `failed` state. The specific
error codes and messages depend on the workflow type (registration,
update, etc.) and the underlying registry response. These workflow
error codes are separate from immediate HTTP error `errors[].code`
values returned by non-2xx responses. Surface
`error.message` to the user for context.

```yaml
{"description": "Error details when a workflow reaches the `failed` state. The specific\nerror codes and messages depend on the workflow type (registration,\nupdate, etc.) and the underlying registry response. These workflow\nerror codes are separate from immediate HTTP error `errors[].code`\nvalues returned by non-2xx responses. Surface\n`error.message` to the user for context.\n", "type": "object", "properties": {"code": {"description": "Machine-readable error code identifying the failure reason.", "type": "string", "example": "registry_rejected"}, "message": {"description": "Human-readable explanation of the failure. May include registry-specific details.", "type": "string", "example": "Registry rejected the request."}}, "required": ["code", "message"]}
```
