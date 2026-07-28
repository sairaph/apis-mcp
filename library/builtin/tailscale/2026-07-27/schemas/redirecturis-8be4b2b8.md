---
title: redirectURIs
page_id: schema-redirecturis-8be4b2b8
path: schemas
description: |-
    The list of permitted redirect URIs for the OAuth authorization code flow.
    At least one redirect URI is required.

    Each URI must use the `https` scheme, except for `localhost`, `127.0.0.1`, and `::1`,
    which may use any scheme. Raw IP address hosts are not permitted.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# redirectURIs

The list of permitted redirect URIs for the OAuth authorization code flow.
At least one redirect URI is required.

Each URI must use the `https` scheme, except for `localhost`, `127.0.0.1`, and `::1`,
which may use any scheme. Raw IP address hosts are not permitted.

```yaml
type: array
items:
    type: string
example:
    - https://example.com/oauth/callback
description: |
    The list of permitted redirect URIs for the OAuth authorization code flow.
    At least one redirect URI is required.

    Each URI must use the `https` scheme, except for `localhost`, `127.0.0.1`, and `::1`,
    which may use any scheme. Raw IP address hosts are not permitted.
```
