---
title: ConfigurationAuditLog
page_id: schema-configurationauditlog-15ae6e32
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# ConfigurationAuditLog

```yaml
type: object
properties:
    eventTime:
        type: string
        description: Timestamp of the audit log event, in RFC 3339 format.
        example: '2024-06-06T15:25:26.583893Z'
    type:
        type: string
        enum:
            - CONFIG
        description: The type of log (always "CONFIG").
    deferredAt:
        type: string
        description: Timestamp recording the time that the audit log rate limiter enqueued the record to be logged at a future time, in RFC 3339 format.
        example: '0001-01-01T00:00:00Z'
    eventGroupID:
        type: string
        description: Identifier assigned to one or more audit log events, all of which are the result of a single operation.
        example: 0378d8f57300d172ef7ae3826e097ef0
    origin:
        type: string
        enum:
            - ADMIN_CONSOLE
            - CONFIG_API
            - CONTROL
            - IDENTITY_PROVIDER
            - NODE
            - SUPPORT_REQUEST
            - STRIPE
            - SECURITY_NOTIFICATION
            - LEGAL_NOTIFICATION
        description: The initiator of the action that generated the event, typically an API or user interface, like the Tailscale admin panel.
        example: ADMIN_CONSOLE
    actor:
        type: object
        description: The person who caused the action related to this event.
        properties:
            id:
                type: string
                description: The ID (user ID or node ID) of the actor.
                example: uZKk3KSfrH11DEVEL
            type:
                type: string
                enum:
                    - USER
                    - NODE
                    - AUTOMATED_WORKER
                    - OAUTH_CLIENT
                    - SCIM
                    - MULLVAD
                    - LOGSTREAM
                    - SECRET_SCANNER
                description: The entity type of the actor.
                example: USER
            loginName:
                type: string
                description: The login name of the actor at time of the action.
                example: lion.dahlia.armadillo@example.com
            displayName:
                type: string
                description: The display name of the actor at time of the action.
            tags:
                type: array
                items:
                    type: string
                description: Indicates the tags owning a node. Its value is only set if `type` is `NODE`.
                example:
                    - server
                    - datacenter1
    target:
        type: object
        description: The object of this event's action.
        properties:
            id:
                type: string
                description: The unique ID (user id, tailnet SID, or node id) of the target.
                example: nBLYviWLGB21DEVEL
            name:
                type: string
                description: Name of the entity at time of the action.
                example: silver-robin-horse-albatross-armadillo.taile18a.ts.net
            type:
                type: string
                enum:
                    - TAILNET
                    - USER
                    - GROUP
                    - NODE
                    - API_KEY
                    - INVITE
                    - SHARE
                    - BILLING
                    - ADMIN_CONSOLE
                    - WEB_INTERFACE
                    - WEBHOOK_ENDPOINT
                    - FAILED_REQUEST
                description: The entity type of Target.
                example: NODE
            isEphemeral:
                type: boolean
                description: Indicates whether the target is ephemeral. Its value should only be set if `type` is `NODE``.
                example: true
            property:
                type: string
                enum:
                    - ACL
                    - ACL_TAGS
                    - ACCOUNT_EMAIL
                    - ADDRESS
                    - ALLOWED_IPS
                    - AUTO_APPROVED_ROUTES
                    - ATTRIBUTES
                    - BILLING_OWNER
                    - COLLECT_SERVICES
                    - COLLECT_POSTURE_IDENTITY
                    - MULLVAD_VPN
                    - DNS_CONFIG
                    - EMAIL
                    - EXIT_NODE
                    - FEATURE
                    - FILE_SHARING
                    - HTTPS
                    - KEY_EXPIRY_TIME
                    - KEY_EXPIRY
                    - LOG_EXIT_FLOWS
                    - LOGSTREAM_ENDPOINT
                    - MAGIC_DNS
                    - MACHINE_AUTH_NEEDED
                    - MACHINE_APPROVAL_NEEDED
                    - USER_APPROVAL_REQUIRED
                    - MACHINE_NAME
                    - MAX_KEY_DURATION
                    - NETWORK_FLOW_LOGGING
                    - GEOSTEERING
                    - NODE_SHARE
                    - TAILNET_INVITE
                    - PAYMENT_INFO
                    - POSTURE_IDENTITY
                    - POSTURE_INTEGRATION
                    - USER_ROLE
                    - SCIM
                    - SECURITY_EMAIL
                    - STRIPE_CUSTOMER_ID
                    - SUBSCRIPTION
                    - SUBSCRIBED_EVENTS
                    - SUPPORT_EMAIL
                    - SECRET
                    - TCD
                    - TKA
                    - AUTH_PROVIDER
                description: The property name on this target which was updated by the event. When empty, the event didn't update any fields on this target.
                example: ALLOWED_IPS
    action:
        type: string
        enum:
            - LOGIN
            - LOGOUT
            - CREATE
            - UPDATE
            - DELETE
            - CANCEL
            - REVOKE
            - APPROVE
            - SUSPEND
            - RESTORE
            - ENABLE
            - DISABLE
            - ACCEPT
            - EXPIRED
            - PUSH_USER
            - PUSH_GROUP
            - VERIFY
            - JOIN_WAITLIST
            - INVITE
            - JOIN
            - LEAVE
            - RESEND
            - MIGRATE_AUTH_PROVIDER
        description: The type of change attempted against the `target`.
        example: CREATE
    old:
        description: The value of `target.property`` prior to the event.
        anyOf:
            - type: string
            - type: number
            - type: integer
            - type: boolean
            - type: array
              items: {}
            - type: object
    new:
        description: The value of `target.property` after the event.
        anyOf:
            - type: string
            - type: number
            - type: integer
            - type: boolean
            - type: array
              items: {}
            - type: object
    actionDetails:
        type: string
        description: Additional information about the event, such as a client-provided reason, if it exists.
    error:
        type: string
        description: Provided when the configuration change failed to be completed. It is a user-presentable reason for the failure.
required:
    - eventTime
    - type
    - eventGroupID
    - origin
    - actor
    - target
    - action
example:
    action: CREATE
    actor:
        displayName: Lion Dahlia Armadillo
        id: uZKk3KSfrH11DEVEL
        loginName: lion.dahlia.armadillo@example.com
        type: USER
    deferredAt: '0001-01-01T00:00:00Z'
    eventGroupID: 0378d8f57300d172ef7ae3826e097ef0
    eventTime: '2024-06-06T15:25:26.583893Z'
    origin: ADMIN_CONSOLE
    target:
        id: nBLYviWLGB21DEVEL
        isEphemeral: true
        name: silver-robin-horse-albatross-armadillo.taile18a.ts.net
        type: NODE
    type: CONFIG
```
