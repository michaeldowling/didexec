# Propose Collaboration Interaction Design

## Overview

```mermaid

---
title: Propose Collaboration
---

sequenceDiagram
    actor Originator
    participant Device@{ "type" : "boundary" }
    participant Framework@{ "type" : "control" }

    participant AgentMatrix@{ "type": "boundary" }
    participant CollaboratorMatrix@{ "type": "boundary" }

    participant Collaborator@{"type": "boundary"}

    Originator ->> Device: User Intent Requiring Collaboration
    activate Device
    
    Device ->> Framework: Create Collaboration Proposal (VC)
    activate Framework
    Framework -->> Device: Raw VC
    deactivate Framework

    Device ->> Framework: Sign Proposal
    activate Framework
    Framework -->> Device: Signed VC for Presentation
    deactivate Framework
    
    Device ->> Framework: Send Proposal
    activate Framework

    loop For each collaborator
        Framework --) CollaboratorMatrix: GetMappingMessage(for: collaborator)
        activate CollaboratorMatrix
        CollaboratorMatrix -->> Framework: .well-known for collab did
        deactivate CollaboratorMatrix
        Framework ->> Framework: Encrypt Signed VC with Collab pubkey
        Framework --) AgentMatrix: Send Message
        AgentMatrix --) CollaboratorMatrix: Deliver Message
        CollaboratorMatrix ->> Collaborator: Deliver Message
                
    end
    
    deactivate Framework

    deactivate Device

```

## Actors

### Originator

The originator, or first, creator and sender of a Collaboration message.  As the originator, it is their role to collect all co-signed ProposedCollaboration messages and, when completed, send the final Collaboration VC to all participants.

### Device

Device is the hardware used by the Agent where a User Intent is performed that requires a Collaboration.  For example, a web browser on a laptop, a mobile application, AI Agent, etc.  Devices use the Framework to perform operations on the Collaboration.

### Framework

### AgentMatrix

### CollaboratorMatrix

### Collaborator


