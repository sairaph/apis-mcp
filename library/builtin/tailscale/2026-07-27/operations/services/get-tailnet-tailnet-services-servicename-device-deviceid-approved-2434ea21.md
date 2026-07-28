---
title: Get approval status of Service on a device
page_id: operation-get-tailnet-tailnet-services-servicename-device-deviceid-approved-17b2161e
path: operations/services
description: |-
    Retrieve the approval status of the specified Service on a specific device.

    OAuth Scope: `services`, `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/services/{serviceName}/device/{deviceId}/approved
operation_ids:
    - getServiceDeviceApproval
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get approval status of Service on a device

`GET /tailnet/{tailnet}/services/{serviceName}/device/{deviceId}/approved`

Operation ID: `getServiceDeviceApproval`

Retrieve the approval status of the specified Service on a specific device.

OAuth Scope: `services`, `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/serviceName'
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Get approval status of Service on a device
description: |
    Retrieve the approval status of the specified Service on a specific device.

    OAuth Scope: `services`, `devices:core`.
operationId: getServiceDeviceApproval
tags:
    - Services
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
