---
title: access_policy_resp
page_id: schema-access-policy-resp-0a1ba421
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_policy_resp

```yaml
{"type": "object", "properties": {"approval_groups": {"$ref": "#/components/schemas/access_approval_groups"}, "approval_required": {"$ref": "#/components/schemas/access_approval_required"}, "connection_rules": {"$ref": "#/components/schemas/access_connection_rules"}, "isolation_required": {"$ref": "#/components/schemas/access_isolation_required"}, "mfa_config": {"$ref": "#/components/schemas/access_mfa_config"}, "purpose_justification_prompt": {"$ref": "#/components/schemas/access_purpose_justification_prompt"}, "purpose_justification_required": {"$ref": "#/components/schemas/access_purpose_justification_required"}, "session_duration": {"$ref": "#/components/schemas/access_session_duration-3"}}, "allOf": [{"$ref": "#/components/schemas/access_base_policy_resp"}]}
```
