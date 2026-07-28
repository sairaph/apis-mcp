---
title: Update tailnet settings
page_id: operation-patch-tailnet-tailnet-settings-d9bd99ec
path: operations/tailnetsettings
description: |-
    Update the settings for a specific tailnet.

    OAuth Scope: `feature_settings` - required to update all settings except those governed by the below scopes.

    OAuth Scope: `logs:network` - required to update the `networkFlowLoggingOn` setting.

    OAuth Scope: `networking_settings` - required to update the `httpsCertificates` setting.

    OAuth Scope: `policy_file` - required to update the `aclsExternallyManagedOn` & `aclsExternalLink` settings.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PATCH
api_endpoints:
    - /tailnet/{tailnet}/settings
operation_ids:
    - updateTailnetSettings
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update tailnet settings

`PATCH /tailnet/{tailnet}/settings`

Operation ID: `updateTailnetSettings`

Update the settings for a specific tailnet.

OAuth Scope: `feature_settings` - required to update all settings except those governed by the below scopes.

OAuth Scope: `logs:network` - required to update the `networkFlowLoggingOn` setting.

OAuth Scope: `networking_settings` - required to update the `httpsCertificates` setting.

OAuth Scope: `policy_file` - required to update the `aclsExternallyManagedOn` & `aclsExternalLink` settings.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Update tailnet settings
description: |
    Update the settings for a specific tailnet.

    OAuth Scope: `feature_settings` - required to update all settings except those governed by the below scopes.

    OAuth Scope: `logs:network` - required to update the `networkFlowLoggingOn` setting.

    OAuth Scope: `networking_settings` - required to update the `httpsCertificates` setting.

    OAuth Scope: `policy_file` - required to update the `aclsExternallyManagedOn` & `aclsExternalLink` settings.
operationId: updateTailnetSettings
tags:
    - TailnetSettings
requestBody:
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/TailnetSettings'
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
