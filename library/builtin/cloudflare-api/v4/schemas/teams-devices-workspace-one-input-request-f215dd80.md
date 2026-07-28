---
title: teams-devices_workspace_one_input_request
page_id: schema-teams-devices-workspace-one-input-request-f215dd80
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_workspace_one_input_request

```yaml
{"type": "object", "properties": {"compliance_status": {"description": "Compliance Status.", "type": "string", "example": "compliant", "enum": ["compliant", "noncompliant", "unknown"], "x-auditable": true}, "connection_id": {"description": "Posture Integration ID.", "type": "string", "example": "bc7cbfbb-600a-42e4-8a23-45b5e85f804f", "x-auditable": true}}, "required": ["connection_id", "compliance_status"], "title": "Workspace One S2S Input"}
```
