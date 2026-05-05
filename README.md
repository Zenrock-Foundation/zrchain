# zrChain

![Banner!](/docs/img/banner.png)

[![License: Source Available License](https://img.shields.io/github/license/Zenrock-Foundation/zrchain.svg?style=flat-square)](https://github.com/Zenrock-Foundation/zrchain/blob/main/LICENSE)
[![Version](https://img.shields.io/github/tag/Zenrock-Foundation/zenrock.svg?style=flat-square)](https://github.com/Zenrock-Foundation/zrchain/releases/latest)
![Go](https://img.shields.io/badge/go-1.26-blue.svg)
[![GitHub Super-Linter](https://img.shields.io/github/actions/workflow/status/Zenrock-Foundation/zrchain/lint.yml?style=flat-square&label=Lint)](https://github.com/marketplace/actions/super-linter)
[![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.com/invite/zenrockfoundation)
[![Twitter](https://badgen.net/badge/icon/twitter?icon=twitter&label)](https://x.com/zenrock)

## Overview

**Zenrock** is a Cosmos SDK-based blockchain providing bedrock security infrastructure for omnichain operations. It specializes in managing wrapped assets (Bitcoin via zenBTC, ZCash via zenZEC, and more) with oracle consensus, MPC-based signing, and cross-chain settlement.

**Core Features:**
- 🔐 **Wrapped Assets** - Bitcoin deposits → zenBTC, ZCash deposits → zenZEC
- 🤝 **Oracle Consensus** - Vote extensions for multi-chain data validation
- 🔑 **MPC Signing** - Secure multi-party computation for transaction signing
- 🔗 **Cross-Chain** - Solana, Ethereum, and more via Sidecar oracle
- 🪐 **IBC Ready** - Full Inter-Blockchain Communication support

## New to the Team?

👋 **Start with the [Developer Onboarding Guide](./docs/ONBOARDING.md)** - It covers:
- Getting up and running in 30 minutes
- Understanding the architecture
- First week checklist
- Common development tasks

Then read [`CLAUDE.md`](./CLAUDE.md) for detailed developer reference and [`AGENTS.md`](./AGENTS.md) for architecture deep-dives.

# Table of Contents
- [Overview](#overview)
- [New to the Team?](#new-to-the-team)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [Contributing](#contributing)
- [Commit & Release Workflow](#commit--release-workflow)
- [Resources](#resources)
- [Support](#support)
- [License](#license)

## Architecture

### High-Level Design

```
┌─────────────────────────────────────────────────────────────┐
│                    zrChain (Cosmos Blockchain)              │
│  ┌────────┬─────────┬──────────┬─────────────────────────┐ │
│  │identity│ treasury│ validation│ zenbtc / dct / zentp   │ │
│  │ policy │  mint   │   wasm    │ Other Custom Modules   │ │
│  └────────┴─────────┴──────────┴─────────────────────────┘ │
│                         ↕ (gRPC)                            │
└──────────────────────────────────────────────────────────────┘
                              ↕
┌──────────────────────────────────────────────────────────────┐
│            Sidecar Oracle (Off-Chain Service)                │
│  • Watches Bitcoin, ZCash, Solana, Ethereum networks        │
│  • Runs EigenLayer operator loops with local keystores                              │
│  • Provides price feeds and block headers                   │
│  • Reports consensus signals to validators                          │
└──────────────────────────────────────────────────────────────┘
```

### Key Modules

| Module | Purpose |
|--------|---------|
| `x/identity` | Keyrings, workspaces, and multi-sig management |
| `x/treasury` | Asset custody and signature request workflows |
| `x/zenbtc` | Bitcoin wrapped asset (production, v0) |
| `x/dct` | Digital Currency Tokens framework (v1+, extensible) |
| `x/validation` | Validator management and oracle consensus |
| `x/policy` | Approval policies and governance |
| `x/mint` | Token minting |
| `x/zentp` | Zenrock token protocol |

See [`CLAUDE.md`](./CLAUDE.md) for full module documentation.

## Quick Start

### Requirements

```bash
./scripts/dev-setup.sh
```

Run the bootstrap script once after cloning to install everything in one shot. It:
- Installs Go 1.25+, [`just`](https://github.com/casey/just), [`uv`](https://docs.astral.sh/uv/), and the full buf/protoc toolchain
- Detects macOS/Linux package managers automatically
- Guides Docker Desktop/Engine installation where automation isn’t possible

Prefer to manage things manually? Make sure you still have Go 1.25+, `make`, Docker, and the proto toolchain exposed on your `PATH`.

### Installation & Local Network

```bash
# Install
git clone git@github.com:Zenrock-Foundation/zrchain.git
cd zrchain
make install

# Start local network (open 2 terminal tabs)
# Terminal 1:
./init.sh --localnet 1

# Terminal 2:
./init.sh --localnet 2

# Verify it's running
zenrockd status
```

See [Developer Onboarding Guide](./docs/ONBOARDING.md) for detailed setup instructions and common commands.

### Project Structure

```
zrchain/
├── app/              # Blockchain setup and module wiring
├── cmd/              # CLI binary entry points
├── x/                # Custom modules (identity, treasury, zenbtc, dct, etc.)
├── sidecar/          # Off-chain oracle service
├── proto/            # Protocol Buffer definitions
├── tests/            # End-to-end tests
├── docs/             # Documentation
└── Makefile          # Build commands
```

**Quick reference:**
- 📖 **Development Guide** → [`docs/ONBOARDING.md`](./docs/ONBOARDING.md)
- 🔧 **Developer Reference** → [`CLAUDE.md`](./CLAUDE.md)
- 🏗️ **Architecture Details** → [`AGENTS.md`](./AGENTS.md)

### Agentic Worktrees

Parallel development is driven by the `Justfile` wrappers around the scripts in `scripts/git/`:
- `just worktree-new` (`scripts/git/worktree-new.sh`) – create a branch+worktree; when Graphite (`gt`) is installed, the script can call `gt create -m …` on the parent branch so the stack is tracked automatically.
- `just worktree-switch` – jump into an existing branch/worktree; with Graphite it uses `gt get` / `gt checkout` to stay in sync and can auto-track non-Graphite branches.
- `just worktree-cleanup` – interactive cleanup of stale worktrees.

Each script still works without Graphite, but when the CLI is available you get prompted to opt-in so the branch metadata is captured for stacked PRs. This is the foundation for the graphite + worktrees + agent workflow used by our AI agents.

## Contributing

We appreciate all contributions to Zenrock and review them closely. See [CONTRIBUTING.md](./CONTRIBUTING.md) for:
- Development guidelines
- Git conventions (commits, branches, PRs)
- Security considerations
- Code review process

## Commit & Release Workflow

Zenrock uses [`semantic-release`](https://semantic-release.gitbook.io/) alongside the Conventional Commits specification to publish new versions automatically. Commit messages must follow the `type(optional-scope)!: description` pattern:
- `feat`: triggers a **minor** version bump.
- `fix` and `perf`: trigger a **patch** bump.
- Commits that include `BREAKING CHANGE` notes (or use `type!: ...`) trigger a **major** bump.
- Other conventional types (`docs`, `chore`, `refactor`, `test`, etc.) do not advance the version but must still follow the format.

See the [Git conventions section](./CONTRIBUTING.md#git-conventions) for detailed guidance and examples.

## Resources

### Documentation

| Document | For New Devs? | Topic |
|----------|:---:|---------|
| [docs/ONBOARDING.md](./docs/ONBOARDING.md) | ✅ **START HERE** | Getting started, first steps, development setup |
| [CLAUDE.md](./CLAUDE.md) | ✅ Recommended | Developer reference, commands, patterns, architecture |
| [AGENTS.md](./AGENTS.md) | For deep dives | Wrapped asset flows, technical architecture |
| [CONTRIBUTING.md](./CONTRIBUTING.md) | ✅ Recommended | Contribution guidelines, git conventions, security |

### External Links

- **[Zenrock Website](https://www.zenrocklabs.io/)** - Project overview and roadmap
- **[Discord](https://discord.com/invite/zenrockfoundation)** - Validator + contributor chat
- **[Telegram](https://t.me/officialZenrock)** - Community updates
- **[Twitter](https://x.com/zenrock)** - Updates and news

## Support

- 💬 **Questions?** - Ask in [Discord](https://discord.com/invite/zenrockfoundation) or GitHub Issues
- 🐛 **Found a bug?** - [Report it here](https://github.com/Zenrock-Foundation/zrchain/issues)
- 💡 **Feature request?** - [Create an issue](https://github.com/Zenrock-Foundation/zrchain/issues)

## License

Licensed under the Source Available License, Zenrock Foundation DAO. See [LICENSE](./LICENSE) file for details.
