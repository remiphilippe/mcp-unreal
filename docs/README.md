# Documentation Index Source Files

This directory contains markdown documentation files that get indexed into the Bleve search index used by the `lookup_docs` and `lookup_class` MCP tools.

## Attribution

Documentation in `ue5.7/` is derived from the [Unreal Engine documentation](https://dev.epicgames.com/documentation/en-us/unreal-engine/) by Epic Games, Inc. Unreal Engine and its documentation are copyright Epic Games, Inc. All rights reserved. This project uses summarized and reformatted excerpts for AI-assisted development tooling under fair use. Unreal Engine is a trademark of Epic Games, Inc.

Documentation in `realtimemesh/` is derived from the [RealtimeMesh](https://github.com/TriAxis-Games/RealtimeMeshComponent) plugin by TriAxis Games.

## Directory Structure

```
docs/
├── ue5.7/            # UE 5.7 API class references and guides
├── realtimemesh/     # RealtimeMesh plugin documentation
└── README.md         # This file
```

## Building the Index

```bash
mcp-unreal --build-index
```

This reads all `.md` files from the `docs/` directory tree and builds the Bleve index at the path specified by `MCP_UNREAL_DOCS_INDEX` (default: `./docs/index.bleve`).

## Adding Documentation

### Class References

Use this format for UE class reference docs:

```markdown
# ClassName

**Parent**: ParentClass
**Module**: ModuleName

Description of the class and its purpose.

## Key Properties

- `PropertyName` — Description of the property

## Key Functions

- `FunctionName(params)` — Description of the function
```

### Guides

Guides are free-form markdown. The indexer will:
1. Extract the title from the first `# Heading`
2. Infer the category from the file path and content keywords
3. Extract UE class name references (AActor, UObject, FStruct patterns)

### Categories

Documents are auto-categorized based on keywords:
- `actor` — actor, spawn, pawn, character, controller, gamemode
- `blueprint` — blueprint, graph, node, pin, compile
- `material` — material, shader, texture, rendering
- `animation` — animation, anim, skeleton, montage, state machine
- `input` — input, enhanced input, action mapping
- `realtimemesh` — realtimemesh, proceduralmesh, mesh generation, lod
- `gameplay` — gameplay, game mode, game state, player state, ability
- `rendering` — rendering, viewport, camera, light, post process
- `networking` — networking, replication, rpc, net

### Project Documentation

Your project's `CLAUDE.md` is automatically indexed at startup with source `"project"`. No need to add it here.
