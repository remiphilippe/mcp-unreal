# APawn

**Parent**: AActor
**Module**: Engine

APawn is the base class for all actors that can be possessed and controlled by a player or AI. It defines the interface for controller attachment, movement input consumption, and input component binding. ACharacter extends APawn by adding a UCharacterMovementComponent and UCapsuleComponent. Use APawn directly for non-humanoid entities such as vehicles or turrets.

## Key Properties

- `Controller` — The AController currently possessing this pawn (nullptr if unpossessed).
- `PlayerState` — The APlayerState associated with this pawn's controller; holds score and network identity.
- `bUseControllerRotationYaw` — If true, the pawn's yaw is driven by the controller's yaw rotation.
- `bUseControllerRotationPitch` — If true, the pawn's pitch is driven by the controller's pitch rotation.
- `bUseControllerRotationRoll` — If true, the pawn's roll is driven by the controller's roll rotation.
- `AutoPossessPlayer` — Automatically possess this pawn with the specified local player on BeginPlay (e.g., EAutoReceiveInput::Player0).
- `AutoPossessAI` — Controls when an AI controller auto-possesses this pawn (PlacedInWorld, Spawned, etc.).
- `AIControllerClass` — The AController subclass used when an AI auto-possesses this pawn.
- `BaseEyeHeight` — Vertical offset from the pawn root for first-person view and GetPawnViewLocation().

## Key Functions

- `PossessedBy(AController* NewController)` — Called when a controller takes possession. Override to react to possession.
- `UnPossessed()` — Called when the current controller releases possession.
- `GetController()` — Returns the current AController; cast to APlayerController or AAIController as needed.
- `GetMovementComponent()` — Returns the UPawnMovementComponent (e.g., UCharacterMovementComponent) attached to this pawn.
- `AddMovementInput(FVector WorldDirection, float ScaleValue, bool bForce)` — Accumulates directional input for the movement component to consume.
- `GetPendingMovementInputVector()` — Returns the accumulated movement input vector before it is consumed.
- `ConsumeMovementInputVector()` — Returns and clears the pending input vector; called by the movement component each tick.
- `SetupPlayerInputComponent(UInputComponent* PlayerInputComponent)` — Override to bind axis/action input events when possessed by a player.
- `GetViewRotation()` — Returns the rotation used for the pawn's point of view (may defer to controller).
- `GetNavAgentPropertiesRef()` — Returns the FNavAgentProperties used by the navigation system for this pawn.
- `IsControlled()` — Returns true if a controller is currently possessing this pawn.
- `IsPlayerControlled()` — Returns true if the possessing controller is an APlayerController.
- `IsLocallyControlled()` — Returns true if this pawn is controlled by a local (non-network-proxy) player controller.
