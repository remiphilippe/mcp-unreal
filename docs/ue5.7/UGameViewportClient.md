# UGameViewportClient

**Module**: Engine
**Header**: `Engine/GameViewportClient.h`
**Inherits**: `UScriptViewportClient`

## Overview

`UGameViewportClient` is the viewport client for the game world, active during Play In Editor (PIE) sessions and standalone game execution. It manages the game's rendering viewport, input routing, and HUD display.

## Key Methods

| Method | Return | Description |
|--------|--------|-------------|
| `GetGameViewport()` | `FViewport*` | Returns the game viewport for rendering and pixel capture |
| `GetViewportSize(FVector2D&)` | `void` | Gets the current viewport dimensions |
| `SetViewportSize(FVector2D)` | `void` | Sets the viewport dimensions |

## Accessing the Game Viewport

```cpp
// Global accessor — valid only during PIE or standalone game
if (GEngine && GEngine->GameViewport) {
    FViewport* GameViewport = GEngine->GameViewport->GetGameViewport();
    // Use for screenshot capture, viewport queries, etc.
}
```

## Editor vs Game Viewport

| Property | Editor Viewport | Game Viewport |
|----------|----------------|---------------|
| Accessor | `GEditor->GetActiveViewport()` | `GEngine->GameViewport->GetGameViewport()` |
| Available | Always in editor | Only during PIE/game |
| Shows | Editor scene view | Game camera view |
| Type | `FEditorViewportClient` | `UGameViewportClient` |
| HUD | No game HUD | Full game HUD |

## Viewport Capture Pattern

```cpp
// Capture PIE viewport if active, else editor viewport
FViewport* Viewport = nullptr;
if (GEditor->IsPlayingSessionInEditor() && GEngine && GEngine->GameViewport) {
    Viewport = GEngine->GameViewport->GetGameViewport();
} else if (GEditor) {
    Viewport = GEditor->GetActiveViewport();
}

if (Viewport) {
    TArray<FColor> Bitmap;
    Viewport->ReadPixels(Bitmap);
    // Process captured pixels
}
```

## Related Classes

- `FViewport` — Low-level viewport for rendering and input
- `UEditorEngine` — `GetActiveViewport()` for editor viewport
- `FWorldContext` — Contains `GameViewport` field for each world context
- `ULocalPlayer` — Player's viewport client reference
