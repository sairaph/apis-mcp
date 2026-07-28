---
title: providerType
page_id: schema-providertype-bd5020d6
path: schemas
description: |-
    The provider type for the webhook destination, or an empty string if none are applicable.
    Outgoing webhook events are sent in the format expected by the provider type if non-empty.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# providerType

The provider type for the webhook destination, or an empty string if none are applicable.
Outgoing webhook events are sent in the format expected by the provider type if non-empty.

```yaml
type: string
enum:
    - slack
    - mattermost
    - googlechat
    - discord
example: slack
description: |
    The provider type for the webhook destination, or an empty string if none are applicable.
    Outgoing webhook events are sent in the format expected by the provider type if non-empty.
```
