# Directory structure

The graph below will give you an idea how this project is structured. The structure will constantly evolving as the project grows so, be sure to check the last commit datetime compared to this document's last changed to prevent confusion.<br>
<br>
(2026/01/22 17:43 UTC+7)

```bash
.
├── cmd/            
│   ├── build/       # Binary executable
│   └── node/        # Entry point
├── config/          
│   └── disver.toml  # Peer config (self)
├── docs/            # Project documentation
├── internal/
│   ├── cli/         # Front-end (CLI)
│   ├── crypto_rust/ # Rust directory
│   ├── host/        # Peer (self) struct
│   ├── rpc/         # Message handling
│   └── utils/       # Utility functions
└── pkg/             # Exportable items
    ├── crypto/      # Cryptography-related
    ├── protocol/    # Protocol-related
    └── types/       # Structs and types

```

