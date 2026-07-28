---
title: Set policy file
page_id: operation-post-tailnet-tailnet-acl-10e5a01a
path: operations/policyfile
description: |-
    Sets the ACL for the given tailnet. HuJSON and JSON are both accepted inputs.
    An `If-Match` header can be set to avoid missed updates.

    On success, returns the updated ACL in JSON or HuJSON according to the `Accept` header.
    Otherwise, errors are returned for incorrectly defined ACLs, ACLs with failing tests on attempted updates, and mismatched `If-Match` header and `ETag`.

    Learn more about [policy file ACL syntax](https://tailscale.com/kb/1337/acl-syntax).

    OAuth Scope: `policy_file`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/acl
operation_ids:
    - setPolicyFile
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set policy file

`POST /tailnet/{tailnet}/acl`

Operation ID: `setPolicyFile`

Sets the ACL for the given tailnet. HuJSON and JSON are both accepted inputs.
An `If-Match` header can be set to avoid missed updates.

On success, returns the updated ACL in JSON or HuJSON according to the `Accept` header.
Otherwise, errors are returned for incorrectly defined ACLs, ACLs with failing tests on attempted updates, and mismatched `If-Match` header and `ETag`.

Learn more about [policy file ACL syntax](https://tailscale.com/kb/1337/acl-syntax).

OAuth Scope: `policy_file`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/AcceptHeaderParam'
```

## Definition

```yaml
summary: Set policy file
description: |
    Sets the ACL for the given tailnet. HuJSON and JSON are both accepted inputs.
    An `If-Match` header can be set to avoid missed updates.

    On success, returns the updated ACL in JSON or HuJSON according to the `Accept` header.
    Otherwise, errors are returned for incorrectly defined ACLs, ACLs with failing tests on attempted updates, and mismatched `If-Match` header and `ETag`.

    Learn more about [policy file ACL syntax](https://tailscale.com/kb/1337/acl-syntax).

    OAuth Scope: `policy_file`.
operationId: setPolicyFile
tags:
    - PolicyFile
parameters:
    - name: If-Match
      in: header
      required: false
      schema:
        type: string
      description: |
        This is a safety mechanism to avoid overwriting other users' updates to the tailnet policy file.

        - Set the `If-Match` value to that of the `ETag` header returned in a `GET` request to `/api/v2/tailnet/{tailnet}/acl`.
          Tailscale compares the `ETag` value in your request to that of the current tailnet file and only replaces the file if there's a match.
          (A mismatch indicates that another update has been made to the file.) For example: `-H "If-Match: \"e0b2816b418\""`.
        - Alternately, set the `If-Match` value to `ts-default` to ensure that the policy file is replaced *only if the current policy file is still the untouched default created automatically* for each tailnet.
          For example: `-H "If-Match: \"ts-default\""`.
requestBody:
    content:
        application/json:
            schema:
                type: object
            examples:
                exampleJSONPolicyFile:
                    $ref: '#/components/examples/ExampleJSONPolicyFile'
        application/hujson:
            schema:
                type: string
            examples:
                exampleHuJSONPolicyFile:
                    $ref: '#/components/examples/ExampleHuJSONPolicyFile'
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                examples:
                    exampleJSONPolicyFile:
                        $ref: '#/components/examples/ExampleJSONPolicyFile'
            application/hujson:
                schema:
                    type: string
                examples:
                    exampleHuJSONPolicyFile:
                        $ref: '#/components/examples/ExampleHuJSONPolicyFile'
    '400':
        description: ACL validation or test error.
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '412':
        description: If-Match hash mismatch.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/Error'
                    example:
                        message: precondition failed, invalid old hash
    '500':
        $ref: '#/components/responses/500'
x-codeSamples:
    - label: Shell cURL
      lang: shell
      source: |-
        curl --request POST \
          --url https://api.tailscale.com/api/v2/tailnet/example.com/acl \
          --header 'Authorization: Bearer YOUR_SECRET_TOKEN' \
          --data-binary '{
          "acls": [
            {
              "action": "accept",
              "ports": [
                "*:*"
              ],
              "users": [
                "*"
              ]
            }
          ],
          "groups": {
            "group:example": [
              "user1@example.com",
              "user2@example.com"
            ]
          },
          "hosts": {
            "example-host-1": "100.100.100.100"
          }
        }'
```
