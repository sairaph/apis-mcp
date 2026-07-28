---
title: TailnetSettings
page_id: schema-tailnetsettings-d97e4d91
path: schemas
description: Settings for a tailnet.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# TailnetSettings

Settings for a tailnet.

```yaml
type: object
description: |
    Settings for a tailnet.
properties:
    aclsExternallyManagedOn:
        type:
            - boolean
            - 'null'
        example: false
        description: |
            Prevents users from editing policies in the admin console to avoid conflicts with external management workflows like GitOps or Terraform.
    aclsExternalLink:
        type: string
        format: uri
        example: https://github.com/example/tailnet-policy
        description: |
            Link to the external tailnet policy definition or management solution for this tailnet.
    devicesApprovalOn:
        type:
            - boolean
            - 'null'
        example: false
        description: |
            Whether [device approval](/kb/1099/device-approval) is enabled for the tailnet.
    devicesAutoUpdatesOn:
        type:
            - boolean
            - 'null'
        example: false
        description: |
            Whether [auto updates](/kb/1067/update#auto-updates) are enabled for devices that belong to this tailnet.
    devicesKeyDurationDays:
        type: integer
        minimum: 1
        maximum: 180
        example: 180
        description: |
            The [key expiry](/kb/1028/key-expiry) duration for devices on this tailnet.
    usersApprovalOn:
        type:
            - boolean
            - 'null'
        example: true
        description: |
            Whether [user approval](/kb/1239/user-approval) is enabled for this tailnet.
    usersRoleAllowedToJoinExternalTailnets:
        type: string
        enum:
            - none
            - admin
            - member
        example: admin
        description: |
            Which user roles are allowed to [join external tailnets](/kb/1271/invite-any-user).
    networkFlowLoggingOn:
        type:
            - boolean
            - 'null'
        example: false
        description: |
            Whether [network flog logs](/kb/1219/network-flow-logs) are enabled for the tailnet.
    regionalRoutingOn:
        type:
            - boolean
            - 'null'
        example: false
        description: |
            Whether [regional routing](/kb/1115/high-availability#regional-routing) is enabled for the tailnet.
    postureIdentityCollectionOn:
        type:
            - boolean
            - 'null'
        example: false
        description: |
            Whether [identity collection](/kb/1326/device-identity) is enabled for [device posture](/kb/1288/device-posture) integrations for the tailnet.
    httpsEnabled:
        type:
            - boolean
            - 'null'
        example: false
        description: |
            Whether provisioning of [HTTPS certificates](/kb/1153/enabling-https) is enabled for this tailnet.
```
