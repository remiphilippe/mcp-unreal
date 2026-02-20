# UCharacterMovementComponent

**Parent**: UPawnMovementComponent
**Module**: Engine

UCharacterMovementComponent handles all movement for ACharacter including walking, jumping, falling, swimming, flying, and custom movement modes. It manages ground detection, step-up, slope handling, and network prediction/correction. This is the most complex movement component in UE and the default for all character-based gameplay.

## Key Properties

- `MaxWalkSpeed` — Maximum walking speed in cm/s (default: 600)
- `MaxWalkSpeedCrouched` — Speed while crouching (default: 300)
- `JumpZVelocity` — Initial Z velocity when jumping (default: 420)
- `GravityScale` — Multiplier for gravity (1.0 = normal, 0.0 = no gravity)
- `AirControl` — Amount of lateral control while airborne (0-1, default: 0.05)
- `GroundFriction` — Friction when walking on ground (default: 8.0)
- `BrakingDecelerationWalking` — Deceleration when not applying input (default: 2048)
- `MaxAcceleration` — Maximum acceleration rate (default: 2048)
- `MaxStepHeight` — Maximum height the character can step up (default: 45)
- `bCanWalkOffLedges` — Whether the character can walk off edges
- `MovementMode` — Current movement mode (Walking, Falling, Swimming, Flying, Custom)
- `bOrientRotationToMovement` — Auto-rotate character to face movement direction

## Key Functions

- `SetMovementMode(EMovementMode NewMode)` — Changes movement mode (Walking, Falling, Swimming, Flying, Custom).
- `AddImpulse(FVector Impulse)` — Applies physics impulse to the character.
- `AddForce(FVector Force)` — Applies continuous force during the next movement update.
- `Launch(FVector LaunchVelocity)` — Launches character with given velocity.
- `StopMovementImmediately()` — Zeros all velocity.
- `IsFalling()` — Returns true if character is in the air.
- `IsMovingOnGround()` — Returns true if character is on walkable surface.
- `IsSwimming()` — Returns true if character is swimming.
- `IsFlying()` — Returns true if character is in flying mode.
- `GetMaxSpeed()` — Returns max speed for current movement mode.
- `FindFloor(FVector CapsuleLocation, FFindFloorResult& OutFloor)` — Tests for walkable ground below the character.
