# ACharacter

**Parent**: APawn
**Module**: Engine

ACharacter is a specialized APawn that includes a UCharacterMovementComponent for walking, jumping, swimming, and flying. It is the standard base class for player-controlled and AI-controlled humanoid characters. Uses a UCapsuleComponent for collision and a USkeletalMeshComponent for the visual mesh.

## Key Properties

- `Mesh` — The USkeletalMeshComponent for the character's visual representation
- `CapsuleComponent` — UCapsuleComponent for collision (root component)
- `CharacterMovement` — UCharacterMovementComponent handling all movement logic
- `bIsCrouched` — Whether the character is currently crouching
- `JumpMaxCount` — Maximum number of jumps (1 = single jump, 2 = double jump)
- `JumpMaxHoldTime` — How long the player can hold jump for variable height

## Key Functions

- `Jump()` — Triggers a jump. Respects JumpMaxCount for multi-jump.
- `StopJumping()` — Stops jump hold for variable height jumps.
- `Crouch()` — Transitions to crouching state. Shrinks capsule.
- `UnCrouch()` — Returns to standing state.
- `LaunchCharacter(FVector LaunchVelocity, bool bXYOverride, bool bZOverride)` — Applies an impulse to the character (knockback, jump pads).
- `GetCharacterMovement()` — Returns the UCharacterMovementComponent.
- `GetCapsuleComponent()` — Returns the UCapsuleComponent.
- `GetMesh()` — Returns the USkeletalMeshComponent.
- `OnLanded(const FHitResult& Hit)` — Called when character lands after falling. Override for landing effects.
- `OnJumped()` — Called when jump is triggered. Override for jump effects.
- `CanJumpInternal()` — Override to add custom jump conditions.
- `OnMovementModeChanged(EMovementMode PrevMode, uint8 PrevCustomMode)` — Called when movement mode changes (walking, falling, swimming, etc.).
