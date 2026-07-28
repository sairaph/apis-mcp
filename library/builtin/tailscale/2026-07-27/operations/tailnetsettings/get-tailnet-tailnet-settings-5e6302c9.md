---
title: Get tailnet settings
page_id: operation-get-tailnet-tailnet-settings-aa77a62a
path: operations/tailnetsettings
description: |-
    Retrieve the settings for a specific tailnet.

    OAuth Scope: `feature_settings:read` - required to view all settings except those governed by the below scopes.

    OAuth Scope: `logs:network:read` - required to view the `networkFlowLoggingOn` setting.

    OAuth Scope: `networking_settings:read` - required to view the `httpsCertificates` setting.

    OAuth Scope: `policy_file:read` - required to view the `aclsExternallyManagedOn` & `aclsExternalLink` settings.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/settings
operation_ids:
    - getTailnetSettings
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get tailnet settings

`GET /tailnet/{tailnet}/settings`

Operation ID: `getTailnetSettings`

Retrieve the settings for a specific tailnet.

OAuth Scope: `feature_settings:read` - required to view all settings except those governed by the below scopes.

OAuth Scope: `logs:network:read` - required to view the `networkFlowLoggingOn` setting.

OAuth Scope: `networking_settings:read` - required to view the `httpsCertificates` setting.

OAuth Scope: `policy_file:read` - required to view the `aclsExternallyManagedOn` & `aclsExternalLink` settings.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Get tailnet settings
description: |
    Retrieve the settings for a specific tailnet.

    OAuth Scope: `feature_settings:read` - required to view all settings except those governed by the below scopes.

    OAuth Scope: `logs:network:read` - required to view the `networkFlowLoggingOn` setting.

    OAuth Scope: `networking_settings:read` - required to view the `httpsCertificates` setting.

    OAuth Scope: `policy_file:read` - required to view the `aclsExternallyManagedOn` & `aclsExternalLink` settings.
operationId: getTailnetSettings
tags:
    - TailnetSettings
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/TailnetSettings'
    '400':
        $ref: '#/components/responses/400'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
