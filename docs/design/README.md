# Design and Functional Specifications

## Overview

The DIDExec Protocol attempts to solve the problem of agents (Humans on Devices, AI Agents, Systems runnning on behalf of Humans) being able to securely and privately pass data and instructions to each other in a way that is verifiable, auditable, and compliant with the principles of self-sovereign identity and a decentralized, distributed web.

Using the Matrix Protocol as the base communication layer, the DIDExec Protocol defines a set of message types, data formats, and interaction patterns that enable agents to exchange information and instructions in a secure and verifiable way. The protocol is designed to be flexible and extensible, allowing for a wide range of use cases and applications.

## Concepts

### Agent

An Agent is an entity that can send and receive messages using the DIDExec Protocol. Agents can be humans, AI systems, or any other type of entity that can interact with the protocol. Each Agent has a unique identifier (DID) and can have associated metadata, such as a public key for encryption and signing.

### Collaboration

A collaboration is a piece of data that represents a shared context and/or agreement between two or more Agents. It is a structured data object that contains information about the parties involved, the purpose of the collaboration, and any relevant terms or conditions. Collaborations can be used to establish trust and facilitate interactions between agents.

### Shared Instructions

Shared Instructions are a bundle of code, functions, and data that is shared between Agents as part of a collaboration. They can include tasks, goals, or any other type of information that is relevant to the collaboration. Shared Instructions are designed to be flexible and can be used in a wide range of scenarios, from simple task delegation to complex multi-agent coordination.  Shared Instructions are typically written in Javascript and is executed in a distributed manner across the agents involved in the collaboration.

### Verifiable Credentials

Verifiable Credentials are a standard way of representing and exchanging information about an Agent's attributes, qualifications, or other relevant information. They are cryptographically signed and can be verified by other Agents to establish trust and credibility. Verifiable Credentials are used in the DIDExec Protocol to provide additional context and information about the parties involved in a collaboration as well as the rules and attributes of the collaboration itself.

Verifiable Credentials are further defined [here](https://www.w3.org/TR/vc-data-model/).



## Interactions

The DIDExec protocol defines a set of interactions that Agents use to establish a collaboration and exchange Shared Instreuction. These interactions are designed to be flexible and can be adapted to a wide range of use cases and scenarios. The main interactions include:

* Collaboration Interactions
  * Propose Collaboration
  * Accept Collaboration
  * Reject Collaboration
  * Update Collaboration
  * Terminate Collaboration

* Shared Instructions Interactions
  * Execute Shared Instruction

* Protocol to Device Data Interactions
  * Shared Instruction Call Data
  * Emit/Capture Protocol Events

### Propose Collaboration



## Message Types Reference
