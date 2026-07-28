---
title: Update a Service
page_id: operation-put-tailnet-tailnet-services-servicename-6dfb444f
path: operations/services
description: |-
    Update or create the specified Service. If the Service does not exist, it will
    create a Service with the provided details. When creating a new Service, the name
    in the request body must match the serviceName path parameter. When updating an existing
    Service, the path parameter is the current name of the Service, and the name in the request
    body can be used to rename the Service.

    OAuth Scope: `services`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PUT
api_endpoints:
    - /tailnet/{tailnet}/services/{serviceName}
operation_ids:
    - updateService
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update a Service

`PUT /tailnet/{tailnet}/services/{serviceName}`

Operation ID: `updateService`

Update or create the specified Service. If the Service does not exist, it will
create a Service with the provided details. When creating a new Service, the name
in the request body must match the serviceName path parameter. When updating an existing
Service, the path parameter is the current name of the Service, and the name in the request
body can be used to rename the Service.

OAuth Scope: `services`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/serviceName'
```

## Definition

```yaml
summary: Update a Service
description: |
    Update or create the specified Service. If the Service does not exist, it will
    create a Service with the provided details. When creating a new Service, the name
    in the request body must match the serviceName path parameter. When updating an existing
    Service, the path parameter is the current name of the Service, and the name in the request
    body can be used to rename the Service.

    OAuth Scope: `services`.
operationId: updateService
tags:
    - Services
requestBody:
    description: The Service details to create or update.
    required: true
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/VIPServiceInfoPut'
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/VIPServiceInfo'
    '400':
        $ref: '#/components/responses/400'
    '403':
        description: Access to modify the Service is forbidden.
        $ref: '#/components/responses/403'
    '404':
        description: Service not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
