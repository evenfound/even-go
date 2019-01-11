# even-go: EVEN node service implementation on Go

The even-go repository is the node sevice implementation on Go and the embodiment of the EVEN network specification. 

This is a full-featured [[EVEN]](https://evenfound.org/) node with a convenient 
* RPC/IPC interface,
* JSON-REST HTTP interface,
* Embedding as shared code possibility.

It allows users to become part of the [[EVEN]](https://evenfound.org/) network as both a transaction relay
and network information provider through the easy-to-use [[API]](https://evenfound.org/reference).

It is specially designed for users seeking a fast, efficient and fully-compatible network setup.

Running an EVEN node also allows any wallets - self mades or EVEN distros with use ours UI API, users a node to directly connect to for their own wallet transactions.

<!-- *-* **License:** GPLv3 -->

## How to get started

The EVEN network is an independent peer-to-peer network with a first-user, friend-to-friend, network structure based on IPFS functionality:

- As a “first user” network for accessing data streams and APIs provided by other users, you do not need to worry about network identification — the identified service of the node immediately connects to EVEN Network in IPFS.

- As a friend-to-friend network, you must make sure that your service has successfully connected to the network and sees the ping of its members.
 
Everyone will be welcoming and very happy to help you get connected.
If you want to get tokens for your testcase, please just ask in one of the communication channels.

## How it work
> Simple EVEN Network architecture illustration

<p align="center">
  <img src="https://github.com/evenfound/even-network/blob/develop/doc/even-node.png">
</p>

## Folders Structure

> Structure EVEN Node project and short introduction all its subrpojects

<!-- ### A typical top-level directory layout -->

    .
    ├── build                   # Compiled files
    ├── node                    # Source code files
    │   ├── app                 # Node main files 
    │   ├── hdgen               # HD generator
    │   ├── iipfs               # IPFS wrapper
    │   └── ...
    └── README.md

> Use short lowercase names at least for the top-level files and folders except  `README.md`



