---
title: firewall_action
page_id: schema-firewall-action-5a37f329
path: schemas
description: The action to perform when the threshold of matched traffic within the configured period is exceeded.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_action

The action to perform when the threshold of matched traffic within the configured period is exceeded.

```yaml
{"description": "The action to perform when the threshold of matched traffic within the configured period is exceeded.", "type": "object", "anyOf": [{"properties": {"mode": {"$ref": "#/components/schemas/firewall_mode"}, "response": {"$ref": "#/components/schemas/firewall_custom_response"}, "timeout": {"$ref": "#/components/schemas/firewall_timeout"}}, "type": "object"}]}
```
