---
title: api-shield_api_discovery_state
page_id: schema-api-shield-api-discovery-state-f1e81b6e
path: schemas
description: |-
    State of operation in API Discovery
      * `review` - Operation is not saved into API Shield Endpoint Management
      * `saved` - Operation is saved into API Shield Endpoint Management
      * `ignored` - Operation is marked as ignored
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_api_discovery_state

State of operation in API Discovery
  * `review` - Operation is not saved into API Shield Endpoint Management
  * `saved` - Operation is saved into API Shield Endpoint Management
  * `ignored` - Operation is marked as ignored

```yaml
{"description": "State of operation in API Discovery\n  * `review` - Operation is not saved into API Shield Endpoint Management\n  * `saved` - Operation is saved into API Shield Endpoint Management\n  * `ignored` - Operation is marked as ignored\n", "type": "string", "enum": ["review", "saved", "ignored"], "x-auditable": true}
```
