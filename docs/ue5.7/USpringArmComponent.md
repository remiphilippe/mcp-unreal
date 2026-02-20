# USpringArmComponent

**Parent**: USceneComponent
**Module**: Engine

USpringArmComponent implements a camera boom that maintains a child component (typically UCameraComponent) at a fixed arm length behind a target while detecting collision with the world and pulling the child closer to avoid clipping. It also provides configurable camera lag for smooth follow behavior and rotation lag for a trailing rotation effect. Attach UCameraComponent to its socket named `SpringEndpoint`.

## Key Properties

- `TargetArmLength` — The desired distance from the socket origin to the end of the arm in centimeters. Collision may shorten this at runtime.
- `bDoCollisionTest` — If true, traces against the world to prevent the arm from passing through geometry and shortens the arm to the hit point.
- `ProbeSize` — Radius of the sphere used for collision probing along the arm; larger values give more clearance.
- `ProbeChannel` — The ECollisionChannel used by the collision probe sweep (default Camera channel).
- `bUsePawnControlRotation` — If true, the spring arm rotates to match the owning pawn's controller rotation (standard third-person setup).
- `bEnableCameraLag` — If true, the arm end lags behind the socket origin using CameraLagSpeed for smoothing positional transitions.
- `CameraLagSpeed` — Interpolation speed (lower = more lag) for positional lag when bEnableCameraLag is true.
- `bEnableCameraRotationLag` — If true, the arm rotation lags behind the target rotation using CameraRotationLagSpeed.
- `CameraRotationLagSpeed` — Interpolation speed for rotational lag when bEnableCameraRotationLag is true.
- `CameraLagMaxDistance` — Maximum distance the lagged camera position can be from the ideal position; prevents extreme lag at high speeds.
- `SocketOffset` — Local offset applied at the arm's end point (camera socket) after arm length and collision; useful for shoulder offsets.
- `TargetOffset` — Local offset applied at the arm's origin (socket) before arm direction is computed; effectively moves the pivot.
- `bInheritPitch` — Whether to inherit pitch from the parent component's rotation.
- `bInheritYaw` — Whether to inherit yaw from the parent component's rotation.
- `bInheritRoll` — Whether to inherit roll from the parent component's rotation.

## Key Functions

- `GetTargetRotation()` — Returns the rotation the spring arm is targeting this frame, accounting for inherited axes and controller rotation.
- `GetDesiredRotation()` — Returns the desired arm rotation before rotation lag smoothing is applied.
- `GetUnfixedCameraPosition()` — Returns the ideal camera position at full TargetArmLength before collision shortening.
- `IsCollisionFixApplied()` — Returns true if the arm is currently shortened due to a collision hit, useful for gameplay reactions to camera obstruction.
