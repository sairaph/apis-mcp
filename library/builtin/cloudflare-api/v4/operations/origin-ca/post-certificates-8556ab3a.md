---
title: Create Certificate
page_id: operation-post-certificates-ca6b5675
path: operations/origin-ca
description: Create an Origin CA certificate. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /certificates
operation_ids:
    - origin-ca-create-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Certificate

`POST /certificates`

Operation ID: `origin-ca-create-certificate`

Create an Origin CA certificate. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).

## Definition

```yaml
{"operationId": "origin-ca-create-certificate", "summary": "Create Certificate", "description": "Create an Origin CA certificate. You can use an Origin CA Key as your User Service Key or an API token when calling this endpoint ([see above](#requests)).", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"csr": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_csr"}, "hostnames": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostnames"}, "request_type": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_request_type"}, "requested_validity": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_requested_validity"}}, "required": ["hostnames", "request_type", "csr"]}}}}, "responses": {"200": {"description": "Create Certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-2"}}}}, "4XX": {"description": "Create Certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-2"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"user_service_key": []}], "tags": ["Origin CA"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-ca-certificates", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
