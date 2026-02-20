# FWorldContext

**Module**: Engine
**Header**: `Engine/Engine.h`
**Type**: Struct

## Overview

`FWorldContext` manages a single world instance within the engine. The editor maintains multiple world contexts simultaneously — one for the editor world, one for each PIE instance, and potentially more for asset preview worlds.

## Key Fields

| Field | Type | Description |
|-------|------|-------------|
| `WorldType` | `EWorldType::Type` | The type of world: `Editor`, `PIE`, `Game`, `Preview`, `EditorPreview` |
| `World()` | `UWorld*` | Returns the world pointer for this context |
| `PIEInstance` | `int32` | PIE instance index (-1 if not a PIE world) |
| `GameViewport` | `UGameViewportClient*` | Viewport client for game worlds (PIE/standalone) |

## Common Patterns

### Getting the Editor World Context

```cpp
FWorldContext& EditorContext = GEditor->GetEditorWorldContext();
UWorld* EditorWorld = EditorContext.World();
```

### Iterating All World Contexts

```cpp
for (const FWorldContext& Context : GEngine->GetWorldContexts()) {
    if (Context.WorldType == EWorldType::PIE) {
        UWorld* PIEWorld = Context.World();
        // Work with PIE world
    }
}
```

### Checking for PIE

```cpp
if (GEditor->IsPlayingSessionInEditor()) {
    UWorld* PIEWorld = GEditor->PlayWorld;
    // PIE is active, PlayWorld is valid
}
```

## PIE vs Editor World

When Play In Editor starts, the engine creates a new `FWorldContext` with `WorldType::PIE`. The editor world context continues to exist but its actors are not part of the game simulation. Runtime-spawned actors (from GameMode, spawners, etc.) only exist in the PIE world.

## Related Classes

- `UWorld` — The world itself
- `UGameInstance` — Manages game state across world transitions
- `UEditorEngine` — Provides `GetEditorWorldContext()`, `PlayWorld`, `IsPlayingSessionInEditor()`
- `UGameViewportClient` — Viewport client for game/PIE worlds
