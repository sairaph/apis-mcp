---
title: Update approval status of Service on a device
page_id: operation-post-tailnet-tailnet-services-servicename-device-deviceid-approved-7b997c95
path: operations/services
description: |-
    Update the approval status of the specified Service on a specific device.

    OAuth Scope: `services`, `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/services/{serviceName}/device/{deviceId}/approved
operation_ids:
    - updateServiceDeviceApproval
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update approval status of Service on a device

`POST /tailnet/{tailnet}/services/{serviceName}/device/{deviceId}/approved`

Operation ID: `updateServiceDeviceApproval`

Update the approval status of the specified Service on a specific device.

OAuth Scope: `services`, `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/serviceName'
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Update approval status of Service on a device
description: |
    Update the approval status of the specified Service on a specific device.

    OAuth Scope: `services`, `devices:core`.
operationId: updateServiceDeviceApproval
tags:
    - Services
requestBody:
    description: The approval status to set for the Service on the device.
    required: true
    content:
        application/json:
            schema:
                type: object
                properties:
                    approved:
                        type: boolean
                        description: |
                            Indicates whether to approve or revoke approval for the Service on the device.
                        example: true
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/VIPServiceApproval'
    '400':
        description: Invalid parameters or no permission to Services.
        $ref: '#/components/responses/400'
    '403':
        description: Access to the Service or device is forbidden.
        $ref: '#/components/responses/403'
    '404':
        description: Service or device not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
