---
title: Validate and test policy file
page_id: operation-post-tailnet-tailnet-acl-validate-7d3c0c33
path: operations/policyfile
description: |-
    This endpoint works in one of two modes, neither of which modifies your current tailnet policy file:

    - Run ACL tests: When the request body contains ACL tests as a JSON array,
      Tailscale runs ACL tests against the tailnet's current policy file.
      Learn more about [ACL tests](https://tailscale.com/kb/1337/acl-syntax#tests).
    - Validate a new policy file: When the request body is a JSON object,
      Tailscale interprets the body as a hypothetical new tailnet policy file with new ACLs,
      including any new rules and tests.
      It validates that the policy file is parsable and runs tests to validate the existing rules.

    In either case, this method does not modify the tailnet policy file in any way.

    OAuth Scope: `policy_file:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/acl/validate
operation_ids:
    - validateAndTestPolicyFile
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Validate and test policy file

`POST /tailnet/{tailnet}/acl/validate`

Operation ID: `validateAndTestPolicyFile`

This endpoint works in one of two modes, neither of which modifies your current tailnet policy file:

- Run ACL tests: When the request body contains ACL tests as a JSON array,
  Tailscale runs ACL tests against the tailnet's current policy file.
  Learn more about [ACL tests](https://tailscale.com/kb/1337/acl-syntax#tests).
- Validate a new policy file: When the request body is a JSON object,
  Tailscale interprets the body as a hypothetical new tailnet policy file with new ACLs,
  including any new rules and tests.
  It validates that the policy file is parsable and runs tests to validate the existing rules.

In either case, this method does not modify the tailnet policy file in any way.

OAuth Scope: `policy_file:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Validate and test policy file
description: |
    This endpoint works in one of two modes, neither of which modifies your current tailnet policy file:

    - Run ACL tests: When the request body contains ACL tests as a JSON array,
      Tailscale runs ACL tests against the tailnet's current policy file.
      Learn more about [ACL tests](https://tailscale.com/kb/1337/acl-syntax#tests).
    - Validate a new policy file: When the request body is a JSON object,
      Tailscale interprets the body as a hypothetical new tailnet policy file with new ACLs,
      including any new rules and tests.
      It validates that the policy file is parsable and runs tests to validate the existing rules.

    In either case, this method does not modify the tailnet policy file in any way.

    OAuth Scope: `policy_file:read`.
operationId: validateAndTestPolicyFile
tags:
    - PolicyFile
requestBody:
    content:
        application/json:
            schema:
                oneOf:
                    - type: array
                      items:
                        type: object
                        properties:
                            src:
                                type: string
                                example: dave@example.com
                                description: "Specifies the user identity to test, which can be \na [user's email address](https://tailscale.com/kb/1337/acl-syntax#reference-users),\na [group](https://tailscale.com/kb/1337/acl-syntax#groups),\na [tag](https://tailscale.com/kb/1068/acl-tags),\nor a [host](https://tailscale.com/kb/1337/acl-syntax#hosts) that maps to an IP address.\nThe test case runs from the perspective of a device authenticated with the provided identity.\n"
                            srcPostureAttrs:
                                type: object
                                description: "Specifies the [device posture attributes](https://tailscale.com/kb/1337/acl-syntax#proto-1)\nas key-value pairs to use when evaluating posture conditions in access rules.\nYou only need to use this field if the access rules contain \n[device posture conditions](https://tailscale.com/kb/1288/device-posture#device-posture-conditions).\n"
                                additionalProperties:
                                    x-additionalPropertiesName: Posture attributes
                                    anyOf:
                                        - type: string
                                        - type: number
                                        - type: boolean
                                example:
                                    node:os: windows
                            proto:
                                type: string
                                example: tcp
                                description: |
                                    Specifies the IP protocol for `accept` and `deny` rules,
                                    similar to the `proto` field in [ACL rules](https://tailscale.com/kb/1337/acl-syntax#acls).
                                    When omitted, the test checks for either TCP or UDP access.
                            accept:
                                type: array
                                items:
                                    type: string
                                description: |
                                    Specifies destinations to accept. Each destination in the list is of the form `host:port`
                                    where `port` is a single numeric port and `host` is in the format described in the
                                    [acl syntax](https://tailscale.com/kb/1337/acl-syntax#accept-and-deny-destinations) documentation.

                                    Sources in `src` and `destinations` must refer to specific entities and do not support `*` wildcards.
                                    For example, an `accept` destination cannot be `tags:*`.
                            deny:
                                type: array
                                items:
                                    type: string
                                description: |
                                    Specifies destinations to deny. Each destination in the list is of the form `host:port`
                                    where `port` is a single numeric port and `host` is in the format described in the
                                    [acl syntax](https://tailscale.com/kb/1337/acl-syntax#accept-and-deny-destinations) documentation.

                                    Sources in `src` and `destinations` must refer to specific entities and do not support `*` wildcards.
                                    For example, a `deny` destination cannot be `tags:*`.
                        required:
                            - src
                      description: |
                        Array of ACL tests.
                    - type: string
                      description: |
                        The JSON representation of the policy file.
            examples:
                aclTest:
                    summary: Perform ACL tests
                    value:
                        - src: user1@example.com
                          accept:
                            - example-host-1:22
                          deny:
                            - example-host-2:100
                aclValidateJSON:
                    summary: Validate JSON policy file
                    value: |
                        {
                          "acls": [
                           { "action": "accept", "src": ["100.105.106.107"], "dst": ["1.2.3.4:*"] },
                          ],
                          "tests", [
                            {"src": "100.105.106.107", "allow": ["1.2.3.4:80"]}
                          ],
                        }
        application/hujson:
            schema:
                type: string
                description: |
                    The HuJSON representation of the policy file.
            examples:
                aclValidateHuJSON:
                    summary: Validate HuJSON policy file
                    value: |
                        // Example/default ACLs for unrestricted connections.
                        {
                          // Declare static groups of users beyond those in the identity service.
                          "groups": {
                            "group:example": ["user1@example.com", "user2@example.com"]
                          },
                          // Declare convenient hostname aliases to use in place of IP addresses.
                          "hosts": {
                            "example-host-1": "100.100.100.100"
                          },
                          // Access control lists.
                          "acls": [
                            // Match absolutely everything. Comment out this section if you want
                            // to define specific ACL restrictions.
                            { "action": "accept", "users": ["*"], "ports": ["*:*"] }
                          ]
                        }
responses:
    '200':
        description: Validation or tests have run. An empty response body implies passing validation or tests.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        message:
                            type: string
                            example: test(s) failed
                        data:
                            type: array
                            items:
                                type: object
                            example:
                                - user: user1@example.com
                                  errors:
                                    - 'address "2.2.2.2:22": want: Drop, got: Accept'
                examples:
                    testsOrValidationFailed:
                        summary: Tests or validation failed
                        value:
                            message: test(s) failed
                            data:
                                - user: user1@example.com
                                  errors:
                                    - 'address "2.2.2.2:22": want: Drop, got: Accept'
                    validateSCIMGroupsNotSynced:
                        summary: Groups not syncing from SCIM
                        value:
                            message: warning(s) found
                            data:
                                - user: group:unknown@example.com
                                  warnings:
                                    - group is not syncing from SCIM and will be ignored by rules in the policy file
                    testsOrValidationSuccess:
                        summary: Tests or validation succeeded
                        value: {}
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
