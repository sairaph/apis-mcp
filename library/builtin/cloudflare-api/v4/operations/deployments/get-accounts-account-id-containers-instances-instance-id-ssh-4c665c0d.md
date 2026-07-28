---
title: Get credentials to SSH into a Container
page_id: operation-get-accounts-account-id-containers-instances-instance-id-ssh-39cb916d
path: operations/deployments
description: Get a JWT to hit the SSH port on a given container.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/containers/instances/{instance_id}/ssh
operation_ids:
    - containerWranglerSsh
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get credentials to SSH into a Container

`GET /accounts/{account_id}/containers/instances/{instance_id}/ssh`

Operation ID: `containerWranglerSsh`

Get a JWT to hit the SSH port on a given container.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "containerWranglerSsh", "summary": "Get credentials to SSH into a Container", "description": "Get a JWT to hit the SSH port on a given container.", "parameters": [{"name": "instance_id", "in": "path", "description": "The ID of the container instance. This is the same ID listed in Wrangler and the Cloudflare Dash", "required": true, "schema": {"$ref": "#/components/schemas/cc_InstanceID"}}], "responses": {"200": {"description": "Credentials to SSH into a Container", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_WranglerSSHResponse"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Unknown account", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Deployment not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "503": {"description": "Instance exists but is not yet ready to use", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Deployments"], "x-api-token-group": ["Workers Containers Write", "Workers Containers Read"]}
```
