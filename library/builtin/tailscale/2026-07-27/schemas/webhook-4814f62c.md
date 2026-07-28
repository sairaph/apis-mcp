---
title: Webhook
page_id: schema-webhook-4814f62c
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Webhook

```yaml
type: object
properties:
    endpointId:
        type: string
        example: '123456'
        description: |
            ID of the webhook endpoint.
    endpointUrl:
        type: string
        example: https://example.com/endpoint
        description: |
            The endpoint that events are sent to from Tailscale via POST requests.
    providerType:
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
    creatorLoginName:
        type: string
        example: user@example.com
        description: |
            The login name for the creator of the webhook endpoint.
            In some cases, such as webhooks created with an OAuth client, this can be blank.
    created:
        type: string
        format: date-time
        example: '2022-12-01T05:23:30Z'
        description: |
            The time the webhook endpoint was created.
    lastModified:
        type: string
        format: date-time
        example: '2022-12-01T05:23:30Z'
        description: |
            The time the webhook endpoint was last modified.
    subscriptions:
        type: array
        items:
            type: string
            enum:
                - nodeCreated
                - nodeNeedsApproval
                - nodeApproved
                - nodeKeyExpiringInOneDay
                - nodeKeyExpired
                - nodeDeleted
                - nodeSigned
                - nodeNeedsSignature
                - policyUpdate
                - userCreated
                - userNeedsApproval
                - userSuspended
                - userRestored
                - userDeleted
                - userApproved
                - userRoleUpdated
                - subnetIPForwardingNotEnabled
                - exitNodeIPForwardingNotEnabled
        example:
            - nodeCreated
            - userDeleted
        description: |
            The list of subscribed events that trigger POST requests to the configured endpoint URL.
            Learn more about [webhook events](/kb/1213/webhooks#events).
    secret:
        type: string
        format: password
        example: xxxxx
        description: |
            The webhook secret associated with the endpoint.
            Only populated on creation or when the secret is rotated.

            This secret is used for generating the `Tailscale-Webhook-Signature` header in requests sent to the endpoint URL.
            Learn more about [verifying webhook event signatures](/kb/1213/webhooks#verifying-an-event-signature).
```
