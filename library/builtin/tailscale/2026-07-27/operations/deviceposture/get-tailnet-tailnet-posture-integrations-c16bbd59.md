---
title: List all posture integrations
page_id: operation-get-tailnet-tailnet-posture-integrations-372ba469
path: operations/deviceposture
description: |-
    List all of the posture integrations for a tailnet.

    OAuth Scope: `feature_settings:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/posture/integrations
operation_ids:
    - getPostureIntegrations
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List all posture integrations

`GET /tailnet/{tailnet}/posture/integrations`

Operation ID: `getPostureIntegrations`

List all of the posture integrations for a tailnet.

OAuth Scope: `feature_settings:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: List all posture integrations
description: |
    List all of the posture integrations for a tailnet.

    OAuth Scope: `feature_settings:read`.
operationId: getPostureIntegrations
tags:
    - DevicePosture
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        integrations:
                            type: array
                            items:
                                $ref: '#/components/schemas/PostureIntegration'
                            description: List of PostureIntegrations.
                example:
                    integrations:
                        - clientId: myclientid
                          cloudId: us-1
                          configUpdated: '2024-06-18T13:43:43.239839Z'
                          id: pcBEPQVMpki7DEVEL
                          provider: falcon
                          status:
                            error: |-
                                oauth2: cannot fetch token: 401 Unauthorized
                                Response: {
                                 "meta": {
                                  "query_time": 2.63e-7,
                                  "powered_by": "crowdstrike-api-gateway",
                                  "trace_id": "f36431bf-cda6-40c9-9afe-6d3db7c55026"
                                 },
                                 "errors": [
                                  {
                                   "code": 401,
                                   "message": "access denied, authorization failed"
                                  }
                                 ]
                                }
                            lastSync: '2024-06-18T08:43:43.777283-05:00'
                            matchedCount: 0
                            possibleMatchedCount: 0
                            providerHostCount: 0
                        - clientId: 93013672-b00c-4344-80ca-7ecf74f9dce1
                          cloudId: global
                          configUpdated: '2024-06-18T13:44:28.250168Z'
                          id: p56wQiqrn7mfDEVEL
                          provider: intune
                          status:
                            error: |-
                                Invalid Tenant ID.
                                Microsoft error: AADSTS90002: Tenant 'd1ae389b-5207-43a2-afca-2de6b03ac7e3' not found. Check to make sure you have the correct tenant ID and are signing into the correct cloud. Check with your subscription administrator, this may happen if there are no active subscriptions for the tenant. Trace ID: f6237360-98a2-4889-913b-e3d80aba7d00 Correlation ID: a2024a6e-7757-4406-8a8d-1b6ac2e03ad5 Timestamp: 2024-06-18 13:44:33Z
                            lastSync: '2024-06-18T08:44:33.872282-05:00'
                            matchedCount: 0
                            possibleMatchedCount: 0
                            providerHostCount: 0
                          tenantId: d1ae389b-5207-43a2-afca-2de6b03ac7e3
                        - clientId: 6cabf059-21c9-44d6-bbde-02898f7430dd
                          cloudId: mysubdomain.jamfcloud.com
                          configUpdated: '2024-06-18T13:44:56.333103Z'
                          id: pFKyGf5YerWbDEVEL
                          provider: jamfpro
                        - clientId: ''
                          cloudId: mydomain.kandji.io
                          configUpdated: '2024-06-18T13:45:13.631878Z'
                          id: ph4mxtePUR2LDEVEL
                          provider: kandji
                        - clientId: ''
                          cloudId: ''
                          configUpdated: '2024-06-18T13:45:20.919656Z'
                          id: pvvTj3FGQAh8DEVEL
                          provider: kolide
                          status:
                            error: 'Kolide returned unexpected status code: 401 Unauthorized'
                            lastSync: '2024-06-18T08:45:23.74563-05:00'
                            matchedCount: 0
                            possibleMatchedCount: 0
                            providerHostCount: 0
                        - clientId: ''
                          cloudId: mydomain.sentinelone.net
                          configUpdated: '2024-06-18T13:45:36.128461Z'
                          id: pg9Wqc8sZW2TDEVEL
                          provider: sentinelone
    '403':
        description: User does not have sufficient access to list posture integrations.
        $ref: '#/components/responses/403'
```
