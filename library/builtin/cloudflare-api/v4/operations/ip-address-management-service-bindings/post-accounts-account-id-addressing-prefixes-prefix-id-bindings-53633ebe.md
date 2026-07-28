---
title: Create Service Binding
page_id: operation-post-accounts-account-id-addressing-prefixes-prefix-id-bindings-be0a78bf
path: operations/ip-address-management-service-bindings
description: |-
    Creates a new Service Binding, routing traffic to IPs within the given CIDR to a service running on Cloudflare's network.
    **NOTE:** The first Service Binding created for an IP Prefix must exactly match the IP Prefix's CIDR. Subsequent Service Bindings may be created with a more-specific CIDR. Refer to the  [Service Bindings Documentation](https://developers.cloudflare.com/byoip/service-bindings/) for compatibility details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings
operation_ids:
    - ip-address-management-service-bindings-create-service-binding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Service Binding

`POST /accounts/{account_id}/addressing/prefixes/{prefix_id}/bindings`

Operation ID: `ip-address-management-service-bindings-create-service-binding`

Creates a new Service Binding, routing traffic to IPs within the given CIDR to a service running on Cloudflare's network.
**NOTE:** The first Service Binding created for an IP Prefix must exactly match the IP Prefix's CIDR. Subsequent Service Bindings may be created with a more-specific CIDR. Refer to the  [Service Bindings Documentation](https://developers.cloudflare.com/byoip/service-bindings/) for compatibility details.

## Definition

```yaml
{"operationId": "ip-address-management-service-bindings-create-service-binding", "summary": "Create Service Binding", "description": "Creates a new Service Binding, routing traffic to IPs within the given CIDR to a service running on Cloudflare's network.\n**NOTE:** The first Service Binding created for an IP Prefix must exactly match the IP Prefix's CIDR. Subsequent Service Bindings may be created with a more-specific CIDR. Refer to the  [Service Bindings Documentation](https://developers.cloudflare.com/byoip/service-bindings/) for compatibility details.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_create_binding_request"}}}}, "responses": {"201": {"description": "The created Service Binding", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/addressing_service_binding"}}}]}}}}, "4XX": {"description": "Create Service Binding response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Service Bindings"], "x-api-token-group": ["IP Prefixes: Write"]}
```
