---
title: OAuthApp
page_id: schema-oauthapp-554691ae
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# OAuthApp

```yaml
type: object
properties:
    id:
        type: string
        example: a123456CNTRL
        description: |
            The unique identifier for the OAuth app.
    name:
        type: string
        minLength: 3
        maxLength: 50
        pattern: ^[a-zA-Z0-9._-]+$
        example: my-oauth-app
        description: |
            The name of the OAuth app.
            Must be between 3 and 50 characters and contain only alphanumeric characters, dashes, periods, and underscores.
    description:
        type: string
        maxLength: 300
        example: An OAuth app used to provision devices.
        description: |
            A human-readable description of the OAuth app.
            Must be at most 300 characters.
    redirectURIs:
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
    scopes:
        type: array
        items:
            type: string
        example:
            - auth_keys:create
        description: |
            The list of OAuth scopes granted to the app.
            Must be non-empty.
            Learn more about [OAuth clients and scopes](/kb/1215/oauth-clients).
    allowedNodeAttributes:
        type: array
        items:
            type: string
        example:
            - custom:myattribute
        description: |
            The list of custom node attributes that the OAuth app is allowed to set.
    clientSecret:
        type: string
        format: password
        example: xxxxx
        description: |
            The client secret for the OAuth app.
            Only populated when the app is created and cannot be retrieved afterwards.
    created:
        type: string
        format: date-time
        example: '2022-12-01T05:23:30Z'
        description: |
            The time the OAuth app was created.
    updated:
        type: string
        format: date-time
        example: '2022-12-01T05:23:30Z'
        description: |
            The time the OAuth app was last updated.
```
