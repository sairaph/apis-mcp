---
title: Device
page_id: schema-device-6ba0bdec
path: schemas
description: |-
    A Tailscale device (sometimes referred to as *node* or *machine*), is any computer or mobile device that joins a tailnet.

    Each device has a unique ID (`nodeId` in the schema below) that is used to identify the device in API calls.
    This ID can be found by going to the [Machines](https://login.tailscale.com/admin/machines) page in the admin console,
    selecting the relevant device, then finding the ID in the Machine Details section.
    You can also [list all devices](#tag/devices/get/tailnet/{tailnet}/devices) in the tailnet to get their `nodeId` values.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Device

A Tailscale device (sometimes referred to as *node* or *machine*), is any computer or mobile device that joins a tailnet.

Each device has a unique ID (`nodeId` in the schema below) that is used to identify the device in API calls.
This ID can be found by going to the [Machines](https://login.tailscale.com/admin/machines) page in the admin console,
selecting the relevant device, then finding the ID in the Machine Details section.
You can also [list all devices](#tag/devices/get/tailnet/{tailnet}/devices) in the tailnet to get their `nodeId` values.

```yaml
type: object
description: |
    A Tailscale device (sometimes referred to as *node* or *machine*), is any computer or mobile device that joins a tailnet.

    Each device has a unique ID (`nodeId` in the schema below) that is used to identify the device in API calls.
    This ID can be found by going to the [Machines](https://login.tailscale.com/admin/machines) page in the admin console,
    selecting the relevant device, then finding the ID in the Machine Details section.
    You can also [list all devices](#tag/devices/get/tailnet/{tailnet}/devices) in the tailnet to get their `nodeId` values.
properties:
    addresses:
        type: array
        items:
            type: string
        example:
            - 100.87.74.78
            - fd7a:115c:a1e0:ac82:4843:ca90:697d:c36e
        description: |
            A list of Tailscale IP addresses for the device,
            including both IPv4 (formatted as 100.x.y.z)
            and IPv6 (formatted as fd7a:115c:a1e0:a:b:c:d:e) addresses.
    id:
        type: string
        example: '92960230385'
        description: |
            The legacy identifier for a device; you
            can supply this value wherever {deviceId} is indicated in the
            endpoint. Note that although "id" is still accepted, "nodeId" is
            preferred.
    nodeId:
        type: string
        example: n292kg92CNTRL
        description: |
            The preferred identifier for a device;
            supply this value wherever {deviceId} is indicated in the endpoint.
    user:
        type: string
        example: amelie@example.com
        description: |
            The user who registered the node. For untagged nodes,
             this user is the device owner.
    name:
        type: string
        example: pangolin.tailfe8c.ts.net
        description: |
            The MagicDNS name of the device.
            Learn more about MagicDNS at https://tailscale.com/kb/1081/.
    hostname:
        type: string
        example: pangolin
        description: |
            The machine name in the admin console.
            Learn more about machine names at https://tailscale.com/kb/1098/.
    clientVersion:
        type: string
        example: v1.36.0
        description: |
            The version of the Tailscale client
            software; this is empty for external devices.
    updateAvailable:
        type: boolean
        example: false
        description: |
            'true' if a Tailscale client version
            upgrade is available. This value is empty for external devices.
    os:
        type: string
        example: linux
        description: |
            The operating system that the device is running.
    created:
        type: string
        format: date-time
        example: '2022-12-01T05:23:30Z'
        description: |
            The date on which the device was added
            to the tailnet; this is empty for external devices.
    connectedToControl:
        type: boolean
        example: true
        description: |
            Indicates if the device recently maintained a TCP connection to the Tailscale control server.
    lastSeen:
        type: string
        format: date-time
        example: '2022-12-01T05:23:30Z'
        description: |
            When the device was last connected to the Tailscale control server. Omitted if the device has never been online or `connectedToControl` is true.
    keyExpiryDisabled:
        type: boolean
        example: false
        description: |
            'true' if the keys for the device will not expire.
            Learn more at https://tailscale.com/kb/1028/.
    expires:
        type: string
        format: date-time
        example: '2023-05-30T04:44:05Z'
        description: |
            The expiration date of the device's auth key.
            Learn more about key expiry at https://tailscale.com/kb/1028/.
    authorized:
        type: boolean
        example: false
        description: |
            'true' if the device has been authorized to join the tailnet; otherwise, 'false'.
            Learn more about device authorization at https://tailscale.com/kb/1099/.
    isExternal:
        type: boolean
        example: false
        description: |
            'true', indicates that a device is not a member of the tailnet, but is shared in to the tailnet;
            if 'false', the device is a member of the tailnet.
            Learn more about node sharing at https://tailscale.com/kb/1084/.
    multipleConnections:
        type: boolean
        example: true
        description: |
            'true', indicates that multiple devices are currently connected using the same node key, which is usually a sign of node state being copied between machines.
            If only one device is connected using this node's key, the field is omitted.
            If the number of live connections goes back to 0 or 1, this field is also omitted, meaning it's not sticky. In case an attacker steals node state from a legitimate node, they can mask their activities by not connecting concurrently with the legitimate node.
    machineKey:
        type: string
        example: ''
        description: |
            For internal use and is not required for any API operations.
            This value is empty for external devices.
    nodeKey:
        type: string
        example: nodekey:01234567890abcdef
        description: |
            Mostly for internal use, required for select operations, such as adding a node to a locked tailnet.
            Learn about tailnet locks at https://tailscale.com/kb/1226/.
    blocksIncomingConnections:
        type: boolean
        example: false
        description: |
            'true' if the device is not allowed to accept any connections over Tailscale, including pings.
            Learn more in the "Allow incoming connections" section of https://tailscale.com/kb/1072/.
    enabledRoutes:
        type: array
        items:
            type: string
        example:
            - 10.0.0.0/16
            - 192.168.1.0/24
        description: |
            The subnet routes for this device that have been approved by a tailnet admin.
            Learn more about subnet routes at https://tailscale.com/kb/1019/.
    advertisedRoutes:
        type: array
        items:
            type: string
        example:
            - 10.0.0.0/16
            - 192.168.1.0/24
        description: |
            The subnets this device requests to expose.
            Learn more about subnet routes at https://tailscale.com/kb/1019/.
    clientConnectivity:
        type: object
        description: |
            clientConnectivity provides a report on the device's current physical network conditions.
        properties:
            endpoints:
                type: array
                items:
                    type: string
                example:
                    - 199.9.14.201:59128
                    - 192.68.0.21:59128
                description: |
                    Client's magicsock UDP IP:port endpoints (IPv4 or IPv6).
            mappingVariesByDestIP:
                type: boolean
                description: |
                    'true' if the host's NAT mappings vary based on the destination IP.
                example: false
            latency:
                type: object
                additionalProperties:
                    type: object
                    x-additionalPropertiesName: DERP server location
                    properties:
                        preferred:
                            type: boolean
                            description: |
                                'true' for the node's preferred DERP server for incoming traffic.
                        latencyMs:
                            type: number
                            format: float
                            description: |
                                Current latency to DERP server.
                description: Map of DERP server locations and their current latency.
                example:
                    Dallas:
                        latencyMs: 60.463043
                    New York City:
                        preferred: true
                        latencyMs: 31.323811
            clientSupports:
                type: object
                description: |
                    Identifies features supported by the client.
                properties:
                    hairPinning:
                        type:
                            - boolean
                            - 'null'
                        example: null
                        description: |
                            This information is no longer tracked and will always be null.
                    ipv6:
                        type:
                            - boolean
                            - 'null'
                        example: false
                        description: |
                            'true' if the device OS supports IPv6,
                            regardless of whether IPv6 internet connectivity is available.
                    pcp:
                        type:
                            - boolean
                            - 'null'
                        example: false
                        description: |
                            'true' if PCP port-mapping service exists on your router.
                    pmp:
                        type:
                            - boolean
                            - 'null'
                        example: false
                        description: |
                            'true' if NAT-PMP port-mapping service exists on your router.
                    udp:
                        type:
                            - boolean
                            - 'null'
                        example: false
                        description: |
                            'true' if UDP traffic is enabled on the current network;
                            if 'false', Tailscale may be unable to make direct connections, and will rely on our DERP servers.
                    upnp:
                        type:
                            - boolean
                            - 'null'
                        example: false
                        description: |
                            'true' if UPnP port-mapping service exists on your router.
        example:
            endpoints:
                - 199.9.14.201:59128
                - 192.68.0.21:59128
            latency:
                Dallas:
                    latencyMs: 60.463043
                New York City:
                    preferred: true
                    latencyMs: 31.323811
            mappingVariesByDestIP: false
            clientSupports:
                hairPinning: null
                ipv6: false
                pcp: false
                pmp: false
                udp: false
                upnp: false
    tags:
        type: array
        items:
            type: string
        example:
            - tag:golink
        description: |
            Lets you assign an identity to a device that is separate from human users, and use it as part of an ACL to restrict access.
            Once a device is tagged, the tag is the owner of that device.
            A single node can have multiple tags assigned.
            This value is empty for external devices.
            Learn more about tags at https://tailscale.com/kb/1068/.
    tailnetLockError:
        type: string
        example: ''
        description: |
            Indicates an issue with the tailnet lock node-key signature on this device.
            This field is only populated when tailnet lock is enabled.
    tailnetLockKey:
        type: string
        example: ''
        description: |
            The node's tailnet lock key.
            Every node generates a tailnet lock key (so the value will be present) even if tailnet lock is not enabled.
            Learn more about tailnet lock at https://tailscale.com/kb/1226/.
    sshEnabled:
        type: boolean
        example: false
        description: |
            'true' if Tailscale SSH is enabled on this device.
            Learn more about Tailscale SSH at https://tailscale.com/kb/1193/.
    postureIdentity:
        type: object
        description: |
            Contains extra identifiers from the device when the tailnet it is connected to has device posture identification collection enabled.
            If the device has not opted-in to posture identification collection, this will contain {"disabled": true}.
            Learn more about posture identity at https://tailscale.com/kb/1326/device-identity.
        properties:
            serialNumbers:
                type: array
                items:
                    type: string
                example:
                    - CP74LFQJXM
            disabled:
                type: boolean
        example:
            serialNumbers:
                - CP74LFQJXM
    isEphemeral:
        type: boolean
        example: false
        description: |
            'true' if the device is ephemeral.
            Learn more about ephemeral nodes at https://tailscale.com/kb/1111/ephemeral-nodes.
    distro:
        type: object
        description: |
            distro provides details of the operating system distribution, if available.
        properties:
            name:
                type: string
                example: ubuntu
                description: |
                    The operating system distribution that the device is running.
            version:
                type: string
                example: '25.04'
                description: |
                    The operating system distribution version that the device is running.
            codeName:
                type: string
                example: Plucky Puffin
                description: |
                    The operating system distribution code name that the device is running.
        example:
            name: ubuntu
            version: '25.04'
            codeName: Plucky Puffin
```
