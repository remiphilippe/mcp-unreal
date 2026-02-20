# PIE (Play In Editor) World Access Patterns

## Overview

When you press Play in the Unreal Editor, a separate game world is created alongside the existing editor world. This guide covers how to correctly access actors, viewports, and subsystems in both worlds.

## Two Worlds

| Aspect | Editor World | PIE World |
|--------|-------------|-----------|
| Access | `GEditor->GetEditorWorldContext().World()` | `GEditor->PlayWorld` |
| Actors | Placed in editor, persistent | Runtime-spawned by game logic |
| Lifetime | Exists while editor is open | Created on Play, destroyed on Stop |
| Viewport | `GEditor->GetActiveViewport()` | `GEngine->GameViewport->GetGameViewport()` |
| Subsystems | Editor subsystems | World + GameInstance subsystems |

## Checking PIE State

```cpp
// Is PIE currently running?
bool bPIEActive = GEditor && GEditor->IsPlayingSessionInEditor();

// Get the PIE world (nullptr if not playing)
UWorld* PIEWorld = GEditor ? GEditor->PlayWorld : nullptr;
```

## Auto-Selection Pattern

The recommended pattern is to prefer the PIE world when active, falling back to the editor world:

```cpp
UWorld* GetBestWorld() {
    if (GEditor && GEditor->IsPlayingSessionInEditor() && GEditor->PlayWorld) {
        return GEditor->PlayWorld;
    }
    if (GEditor) {
        return GEditor->GetEditorWorldContext().World();
    }
    return nullptr;
}
```

## Viewport Selection

PIE uses `UGameViewportClient`, which shows the game camera view with HUD. The editor uses `FEditorViewportClient`, which shows the scene from the editor camera.

```cpp
FViewport* GetBestViewport() {
    if (GEditor && GEditor->IsPlayingSessionInEditor()
        && GEngine && GEngine->GameViewport) {
        return GEngine->GameViewport->GetGameViewport();
    }
    return GEditor ? GEditor->GetActiveViewport() : nullptr;
}
```

## Common Gotchas

1. **Don't cache PIE pointers.** The PIE world and all its actors are destroyed when the user stops playing. Any cached `UWorld*` or `AActor*` from PIE becomes a dangling pointer.

2. **Editor actors are not in PIE.** If you place an actor in the editor and then press Play, iterating actors in `PlayWorld` will NOT find that editor actor. PIE creates duplicates of persistent level actors.

3. **PIE instance index.** Multiple PIE clients can run simultaneously (for multiplayer testing). Each has its own world. `GEditor->PlayWorld` returns the first instance.

4. **Subsystem differences.** `UWorldSubsystem` instances are per-world, so PIE has its own set. `UEditorSubsystem` instances only exist in the editor context.

5. **GameViewport is null outside PIE.** `GEngine->GameViewport` is only valid during PIE or standalone game. Always null-check before use.

## MCP Plugin Implementation

The MCPUnreal plugin supports a `world` JSON parameter in request bodies:

| Value | Behavior |
|-------|----------|
| `"auto"` (default) | PIE world if active, else editor world |
| `"pie"` | PIE world only (error if not playing) |
| `"editor"` | Editor world only (ignores PIE) |

This allows tools like `get_level_actors` to inspect either world on demand.
