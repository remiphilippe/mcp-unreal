# APlayerController

**Parent**: AController
**Module**: Engine

APlayerController is the interface between a human player and the game world. It handles player input, camera management, HUD, and possession of pawns. Each local player has exactly one APlayerController. In multiplayer, the server has a controller for each connected client.

## Key Properties

- `PlayerCameraManager` — APlayerCameraManager that controls the player's view
- `PlayerInput` — UPlayerInput that processes raw input
- `InputComponent` — UInputComponent for binding input actions
- `bShowMouseCursor` — Whether the mouse cursor is visible
- `bEnableClickEvents` — Whether click events are generated
- `DefaultMouseCursor` — The default cursor shape

## Key Functions

- `Possess(APawn* InPawn)` — Takes control of a pawn. Called automatically by GameMode.
- `UnPossess()` — Releases the currently controlled pawn.
- `GetPawn()` — Returns the currently possessed APawn.
- `GetCharacter()` — Returns the possessed pawn cast to ACharacter (or null).
- `SetViewTarget(AActor* NewViewTarget)` — Sets what the camera looks at.
- `SetViewTargetWithBlend(AActor* NewViewTarget, float BlendTime)` — Smoothly transitions camera to new target.
- `GetHitResultUnderCursor(ECollisionChannel Channel, bool bTraceComplex, FHitResult& HitResult)` — Raycasts from cursor position into the world.
- `GetHitResultUnderCursorForObjects(TArray<TEnumAsByte<EObjectTypeQuery>>& ObjectTypes, bool bTraceComplex, FHitResult& HitResult)` — Raycasts from cursor against specific object types.
- `ClientTravel(const FString& URL, ETravelType TravelType)` — Client-initiated level change.
- `ConsoleCommand(const FString& Command)` — Executes a console command.
- `SetInputMode(FInputModeDataBase& InData)` — Sets input mode (Game Only, UI Only, Game and UI).
- `ProjectWorldLocationToScreen(FVector WorldLocation, FVector2D& ScreenLocation)` — Projects 3D point to screen coordinates.
