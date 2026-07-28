---
title: Preview rule matches
page_id: operation-post-tailnet-tailnet-acl-preview-8cf6d6a6
path: operations/policyfile
description: |-
    When given a user or IP port to match against,
    returns the tailnet policy rules that apply to that resource,
    without saving the policy file to the server.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/acl/preview
operation_ids:
    - previewRuleMatches
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Preview rule matches

`POST /tailnet/{tailnet}/acl/preview`

Operation ID: `previewRuleMatches`

When given a user or IP port to match against,
returns the tailnet policy rules that apply to that resource,
without saving the policy file to the server.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- name: type
  in: query
  description: |
    Specify for which type of resource (user or IP port) matching rules are to be fetched.
    Read about [previewing changes in the admin console](https://tailscale.com/kb/1018/#previewing-changes).

    OAuth Scope: `policy_file:read`.
  required: true
  schema:
    type: string
    enum:
        - user
        - ipport
    x-enumDescriptions:
        user: |
            Specify `user` if the `previewFor` value is a user's email.
            Note that `user` remains in the API for compatibility purposes,
            but has been replaced by `src` in policy files.
        ipport: |
            Specify `ipport` if the `previewFor` value is an IP address and port.
            Note that `ipport` remains in the API for compatibility purposes,
            but has been replaced by `dst` in policy files.
    example: user
- name: previewFor
  in: query
  description: |
    - If `type` is `user`, provide the email of a valid user with registered machines.
    - If `type` is `ipport`, provide an IP address + port: `10.0.0.1:80`.

    The supplied policy file is queried with this parameter to determine which rules match.
  required: true
  schema:
    type: string
    example: 10.0.0.1:80
```

## Definition

```yaml
summary: Preview rule matches
description: |
    When given a user or IP port to match against,
    returns the tailnet policy rules that apply to that resource,
    without saving the policy file to the server.
operationId: previewRuleMatches
tags:
    - PolicyFile
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
        description: The list of rules that apply to the resource.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        matches:
                            type: array
                            items:
                                type: object
                                properties:
                                    users:
                                        type: array
                                        items:
                                            type: string
                                        description: |
                                            Source entities affected by the rule.
                                        example:
                                            - '*'
                                    ports:
                                        type: array
                                        items:
                                            type: string
                                        description: |
                                            Destinations that can be accessed.
                                        example:
                                            - '*.*'
                                    lineNumber:
                                        type: integer
                                        description: |
                                            The rule's location in the policy file.
                                        example: 19
                                required:
                                    - users
                                    - ports
                                    - lineNumber
                        type:
                            type: string
                            example: user
                            description: |
                                Echoes the `type` provided in the request.
                        previewFor:
                            type: string
                            example: user1@example.com
                            description: |
                                Echoes the `previewFor` provided in the request.
                    required:
                        - matches
                        - type
                        - previewFor
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
