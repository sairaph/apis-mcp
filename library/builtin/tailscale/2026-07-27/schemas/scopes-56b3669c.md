---
title: scopes
page_id: schema-scopes-56b3669c
path: schemas
description: |-
    The list of OAuth scopes granted to the app.
    Must be non-empty.
    Learn more about [OAuth clients and scopes](/kb/1215/oauth-clients).
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# scopes

The list of OAuth scopes granted to the app.
Must be non-empty.
Learn more about [OAuth clients and scopes](/kb/1215/oauth-clients).

```yaml
type: array
items:
    type: string
example:
    - auth_keys:create
description: |
    The list of OAuth scopes granted to the app.
    Must be non-empty.
    Learn more about [OAuth clients and scopes](/kb/1215/oauth-clients).
```
