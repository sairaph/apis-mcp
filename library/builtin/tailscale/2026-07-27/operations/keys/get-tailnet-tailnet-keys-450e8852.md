---
title: List tailnet keys
page_id: operation-get-tailnet-tailnet-keys-e1f2a65d
path: operations/keys
description: |-
    Returns a list of active auth keys, API access tokens and trust credentials.

    If the parameter {all} was not specified, the set of keys returned depends on the access token used to make the request:
    - If the API call is made with a user-owned API access token, this returns only the keys owned by that user.
    - If the API call is made with an access token derived from an OAuth client, this returns all OAuth clients for the tailnet.
    - If the API call is made with an access token derived from a federated identity, this returns all federated identities for the tailnet.

    OAuth Scope: `api_access_tokens:read` grants access to personal API access tokens.

    OAuth Scope: `auth_keys:read` grants access to machine auth keys.

    OAuth Scope: `oauth_keys:read` grants access to OAuth clients and OAuth tokens.

    OAuth Scope: `federated_keys:read` grants access to federated identities.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/keys
operation_ids:
    - listTailnetKeys
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List tailnet keys

`GET /tailnet/{tailnet}/keys`

Operation ID: `listTailnetKeys`

Returns a list of active auth keys, API access tokens and trust credentials.

If the parameter {all} was not specified, the set of keys returned depends on the access token used to make the request:
- If the API call is made with a user-owned API access token, this returns only the keys owned by that user.
- If the API call is made with an access token derived from an OAuth client, this returns all OAuth clients for the tailnet.
- If the API call is made with an access token derived from a federated identity, this returns all federated identities for the tailnet.

OAuth Scope: `api_access_tokens:read` grants access to personal API access tokens.

OAuth Scope: `auth_keys:read` grants access to machine auth keys.

OAuth Scope: `oauth_keys:read` grants access to OAuth clients and OAuth tokens.

OAuth Scope: `federated_keys:read` grants access to federated identities.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: List tailnet keys
description: |
    Returns a list of active auth keys, API access tokens and trust credentials.

    If the parameter {all} was not specified, the set of keys returned depends on the access token used to make the request:
    - If the API call is made with a user-owned API access token, this returns only the keys owned by that user.
    - If the API call is made with an access token derived from an OAuth client, this returns all OAuth clients for the tailnet.
    - If the API call is made with an access token derived from a federated identity, this returns all federated identities for the tailnet.

    OAuth Scope: `api_access_tokens:read` grants access to personal API access tokens.

    OAuth Scope: `auth_keys:read` grants access to machine auth keys.

    OAuth Scope: `oauth_keys:read` grants access to OAuth clients and OAuth tokens.

    OAuth Scope: `federated_keys:read` grants access to federated identities.
operationId: listTailnetKeys
tags:
    - Keys
parameters:
    - $ref: '#/components/parameters/all'
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        keys:
                            type: array
                            description: A list of the active keys.
                            items:
                                $ref: '#/components/schemas/Key'
                            example:
                                - id: XXXX14CNTRL
                                  keyType: client
                                  created: '2021-12-09T23:22:39Z'
                                  scopes:
                                    - all
                                  description: test key
                                  userId: uscwcTtzzo11DEVEL
                                - id: XXXXZ3CNTRL
                                  keyType: api
                                  expirySeconds: 7776000
                                  created: '2021-12-09T23:22:39Z'
                                  expires: '2022-03-09T23:22:39Z'
                                  scopes:
                                    - all
                                  description: production key
                                  userId: uscwcTtzzo11DEVEL
                                - id: XXXX43CNTRL
                                  keyType: auth
                                  expirySeconds: 7776000
                                  created: '2021-12-09T23:22:39Z'
                                  expires: '2022-03-09T23:22:39Z'
                                  capabilities:
                                    devices:
                                        create:
                                            reusable: true
                                            ephemeral: false
                                            preauthorized: true
                                            tags:
                                                - tag:example
                                  description: dev access
                                  userId: uscwcTtzzo11DEVEL
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
