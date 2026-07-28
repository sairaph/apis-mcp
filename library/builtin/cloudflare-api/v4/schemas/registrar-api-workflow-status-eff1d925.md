---
title: registrar-api_workflow_status
page_id: schema-registrar-api-workflow-status-eff1d925
path: schemas
description: Status of an async registration workflow.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_workflow_status

Status of an async registration workflow.

```yaml
{"description": "Status of an async registration workflow.", "type": "object", "properties": {"completed": {"description": "Whether the workflow has reached a terminal state. `true` when\n`state` is `succeeded` or `failed`. `false` for `pending`,\n`in_progress`, `action_required`, and `blocked`.\n", "type": "boolean", "example": false}, "context": {"description": "Workflow-specific data for this workflow.\n\nThe workflow subject is identified by `context.domain_name` for\ndomain-centric workflows.\n", "type": "object", "additionalProperties": true}, "created_at": {"type": "string", "format": "date-time"}, "error": {"allOf": [{"$ref": "#/components/schemas/registrar-api_workflow_error"}], "nullable": true}, "links": {"$ref": "#/components/schemas/registrar-api_workflow_links"}, "state": {"description": "Workflow lifecycle state.\n- `pending`: Workflow has been created but not yet started processing.\n- `in_progress`: Actively processing. Continue polling `links.self`.\n  The workflow has an internal deadline and will not remain in this\n  state indefinitely.\n- `action_required`: Paused — requires action by the user (not the\n  system). See `context.action` for what is needed. An automated\n  polling loop must break on this state; it will not resolve on its\n  own without user intervention.\n- `blocked`: The workflow cannot make progress due to a third party\n  such as the domain extension's registry or a losing registrar.\n  No user action will help. Continue polling — the block may resolve\n  when the third party responds.\n- `succeeded`: Terminal. The operation completed successfully.\n  `completed` will be `true`. For registrations, `context.registration`\n  contains the resulting registration resource.\n- `failed`: Terminal. The operation failed. `completed` will be `true`.\n  See `error.code` and `error.message` for the reason. Do not\n  auto-retry without user review.\n", "type": "string", "example": "in_progress", "enum": ["pending", "in_progress", "action_required", "blocked", "succeeded", "failed"]}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["state", "completed", "created_at", "updated_at", "links"]}
```
