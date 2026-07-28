---
title: mcn_azure_setup
page_id: schema-mcn-azure-setup-8dd3dce4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_azure_setup

```yaml
{"type": "object", "properties": {"azure_consent_url": {"type": "string"}, "integration_identity_tag": {"type": "string", "x-auditable": true}, "item_type": {"type": "string", "x-auditable": true}, "tag_cli_command": {"type": "string"}}, "required": ["item_type", "azure_consent_url", "integration_identity_tag", "tag_cli_command"]}
```
