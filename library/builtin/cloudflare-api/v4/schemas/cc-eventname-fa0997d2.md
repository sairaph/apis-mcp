---
title: cc_EventName
page_id: schema-cc-eventname-fa0997d2
path: schemas
description: |-
    Name of the event that describes the kind event that happened.
      - SchedulerPlaced: It's the first event that creates a Cloudchamber placement. It happens when the runtime was able to retrieve deployment resources and start verify everything is correct.
      - NetworkingIPAssigned: It's sent when the Cloudchamber runtime has mapped the IP to the container.
      - VMStarted: It's sent when the Cloudchamber runtime has started the VM. However, it does not mean that the container is healthy.
      - ImagePulled: It's sent when the Cloudchamber runtime has pulled the image successfully.
      - ImagePullError: It's sent when the Cloudchamber runtime is having issues pulling image. The message and details have more information on what happened for debugging.
      - VMFailedToStart: It's sent when the Cloudchamber runtime was unable to boot the VM.
      - VMStopping: It's sent when the scheduler is stopping the VM.
      - VMStopped: It's sent when the VM has finally exited.
      - VMFailed: It's sent when the scheduling of the VM failed in the current location.
      - RuntimeStartFailed: It's sent when the runtime had an internal error.
      - SSHStarted: It's sent when the container has gained network connectivity and has opened the SSH port. This event is only sent when SSH keys are configured.
      - CheckUpdate: Sent when the status of a health or readiness check changes. This may also affect the health status of the placement.
      - DurableObjectConnected: Sent when a durable object instance connects and gains control of the deployment.
        This event is only sent for durable object deployments. It is sent after VMStarted.
      - ContainerStarted: It's sent when the container has started running.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_EventName

Name of the event that describes the kind event that happened.
  - SchedulerPlaced: It's the first event that creates a Cloudchamber placement. It happens when the runtime was able to retrieve deployment resources and start verify everything is correct.
  - NetworkingIPAssigned: It's sent when the Cloudchamber runtime has mapped the IP to the container.
  - VMStarted: It's sent when the Cloudchamber runtime has started the VM. However, it does not mean that the container is healthy.
  - ImagePulled: It's sent when the Cloudchamber runtime has pulled the image successfully.
  - ImagePullError: It's sent when the Cloudchamber runtime is having issues pulling image. The message and details have more information on what happened for debugging.
  - VMFailedToStart: It's sent when the Cloudchamber runtime was unable to boot the VM.
  - VMStopping: It's sent when the scheduler is stopping the VM.
  - VMStopped: It's sent when the VM has finally exited.
  - VMFailed: It's sent when the scheduling of the VM failed in the current location.
  - RuntimeStartFailed: It's sent when the runtime had an internal error.
  - SSHStarted: It's sent when the container has gained network connectivity and has opened the SSH port. This event is only sent when SSH keys are configured.
  - CheckUpdate: Sent when the status of a health or readiness check changes. This may also affect the health status of the placement.
  - DurableObjectConnected: Sent when a durable object instance connects and gains control of the deployment.
    This event is only sent for durable object deployments. It is sent after VMStarted.
  - ContainerStarted: It's sent when the container has started running.

```yaml
{"description": "Name of the event that describes the kind event that happened.\n  - SchedulerPlaced: It's the first event that creates a Cloudchamber placement. It happens when the runtime was able to retrieve deployment resources and start verify everything is correct.\n  - NetworkingIPAssigned: It's sent when the Cloudchamber runtime has mapped the IP to the container.\n  - VMStarted: It's sent when the Cloudchamber runtime has started the VM. However, it does not mean that the container is healthy.\n  - ImagePulled: It's sent when the Cloudchamber runtime has pulled the image successfully.\n  - ImagePullError: It's sent when the Cloudchamber runtime is having issues pulling image. The message and details have more information on what happened for debugging.\n  - VMFailedToStart: It's sent when the Cloudchamber runtime was unable to boot the VM.\n  - VMStopping: It's sent when the scheduler is stopping the VM.\n  - VMStopped: It's sent when the VM has finally exited.\n  - VMFailed: It's sent when the scheduling of the VM failed in the current location.\n  - RuntimeStartFailed: It's sent when the runtime had an internal error.\n  - SSHStarted: It's sent when the container has gained network connectivity and has opened the SSH port. This event is only sent when SSH keys are configured.\n  - CheckUpdate: Sent when the status of a health or readiness check changes. This may also affect the health status of the placement.\n  - DurableObjectConnected: Sent when a durable object instance connects and gains control of the deployment.\n    This event is only sent for durable object deployments. It is sent after VMStarted.\n  - ContainerStarted: It's sent when the container has started running.\n", "type": "string", "enum": ["SchedulerPlaced", "NetworkingIPAssigned", "VMStarted", "ImagePulled", "ImagePullError", "VMFailedToStart", "NetworkingIPAssignmentFailed", "VMRunning", "VMStopping", "VMStopped", "VMFailed", "RuntimeStartFailed", "SSHStarted", "ServiceHealthUpdates", "CheckUpdate", "DurableObjectConnected", "ContainerStarted"]}
```
