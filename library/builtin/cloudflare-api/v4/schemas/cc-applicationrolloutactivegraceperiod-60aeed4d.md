---
title: cc_ApplicationRolloutActiveGracePeriod
page_id: schema-cc-applicationrolloutactivegraceperiod-60aeed4d
path: schemas
description: |-
    Grace period for active instances to stay alive before becoming eligible for shutdown signal due to a rollout, in seconds.
    Defaults to 0.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ApplicationRolloutActiveGracePeriod

Grace period for active instances to stay alive before becoming eligible for shutdown signal due to a rollout, in seconds.
Defaults to 0.

```yaml
{"description": "Grace period for active instances to stay alive before becoming eligible for shutdown signal due to a rollout, in seconds.\nDefaults to 0.\n", "type": "integer", "default": 0, "maximum": 604800, "minimum": 0}
```
