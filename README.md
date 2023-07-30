# Cosmos SDK  for Storing and Retrieving Ethereum State

A prototype blockchain that can read state from Ethereum, verify its validity, and store it on the chain.

NB: A **"state"** here is a value stored in a specific storage slot of a specific Ethereum smart contract.

## Proposed Architecture

The solution consists of following components:

1. **Ethereum Light Client.**
    - _Objective:_ proving the value of a specific storage slot at a given block in a trustless manner.
    - _Responsibilities._
        - Fetching block headers (and the embedded Merkle roots) from Ethereum consensus client.
        - Receiving state update proposals (and the embedded Merkle proof) from the **Ethereum Relay**.
        - Verifying Merkle proofs.
        - Sending verified data to **Ethereum State Storage** for storage.
    - _Prototype simplification:_ an extremely robust external Ethereum Light
      Client ([helios](https://github.com/a16z/helios)) is used as a part of the bundle. It is not integrated with the
      Cosmos SDK Chain, and runs as a separate process on the same machine communicating with the chain via unix socket.
2. **Ethereum Relay.**
    - _Objective:_ submitting Ethereum state update proposals to the Cosmos SDK Chain.
    - _Responsibilities._
        - Fetching the required state from Ethereum execution client and a Merkle proof (`eth_getProof`).
        - Verify the Merkle proof against the Ethereum Light Client Module.
        - Submitting state update proposal to the Ethereum State Storage Module containing:
            - address,
            - storage slot,
            - block number, and
            - the Merkle proof.
3. **Ethereum State Storage.**
    - _Objective:_ storing and retrieving Ethereum state data.
    - _Responsibilities._
        - Receiving state update proposal from the Ethereum Relay Module.
        - Stores state.
        - Retrieves state on user request (Query Service Interface update).

```mermaid
graph TD;
Eth[ Ethereum ] -->| Fetch State and Merkle Proof | R[ Ethereum Relay ]
Eth -->| Fetch Headers with Merkle Root | L
R <.->| Veriy Merkle proof | L
R -->| Store Verified State | E[ Cosmos SDK Modules ]
subgraph C[ Cosmos SDK Chain ]
L[ Ethereum Light Client ]
R[ Ethereum Relay ]
E[ Ethereum State Storage ]
end
G[ Users ] <-->| Query State | E
```

```mermaid
sequenceDiagram
    participant E as Ethereum Network
    participant R as Ethereum Relayer
    participant L as Ethereum Light Client
    participant S as Ethereum State Storage
    participant Q as Query Server
    participant U as Users

    E ->> R: Emit Block Headers & Merkle Roots
    R ->> C: Transmit IBC Packets
    I ->> L: Submit State Update Proposal
    L ->> M: Validate and Forward State Data
    M ->> C: Store Verified State
    U ->> C: Query Stored State

```

## Usage

```
docker-compose up
```

## Key Design Decisions

- **The Light Client** is the choice to perform trustless validation of the fetched Ethereum state without storing the
  entire Ethereum blockchain. Alternatives could be using a trusted oracle to provide the Ethereum state, but that
  reduces the trustlessness of the system and significantly reduces security opening the system to central point of
  failure.
- **Tendermint Consensus** is built-in secure BFT consensus algorithm used here to provide agreement on the Ethereum
  state being stored. All validators in the Cosmos chain will receive the state update proposal. They will each
  independently verify the proof against their own copy of the light client, and vote on whether the state update is
  valid. If 2/3 of validators agree on the update, it is added to the chain.
- **Ethereum query** The `eth_getProof` method is preferred over `eth_getStorageAt` because it not only returns the same
  value but also provides a Merkle proof. This proof is used to validate the authenticity of the returned value against
  the light client.

## Future Improvements

- The current system is extremely simplified, and implemented as a single Cosmos App. The modular design is anticipated
  in the initial component breakdown, and each component is expected to be implemented as a separate Cosmos SDK Module.
  Moving the logic into modules will allow for easier maintenance and upgrades.
- Ethereum Light Client as a Cosmos SDK Module. Separate reusable solution going forward that is in great demand among
  Cosmos SDK developers.
- Ethereum State Storage as a Cosmos SDK Module.
- IBC Relayer is a great option for a cross-chain communication method for the purpose of submitting state update
  proposals to the Cosmos SDK Chain instead of Ethereum Relay Module.
- Synchronize queries to the Ethereum to avoid invalid inconsistent state updates due to the timely fetching process.
- Security checks, at first long-range, replay, or man-in-the-middle attacks on the IBC Relayer, and then on the
  Ethereum Light Client Module.
- Failover and redundancy: implement automated recovery from failures.
- Test the system with a large number of nodes, a terraform or similar script to be developed to deploy the system, as
  well as a load testing tool.