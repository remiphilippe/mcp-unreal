# UCapsuleComponent

**Parent**: UPrimitiveComponent
**Module**: Engine

UCapsuleComponent is a vertical capsule-shaped collision primitive used as the root collision shape for ACharacter. It provides efficient character vs. world collision detection with smooth edge handling for walking over small obstacles. The capsule is defined by a half-height (distance from center to the top hemisphere center) and a radius. Scale from the component hierarchy is factored in via the "Scaled" accessors.

## Key Properties

- `CapsuleHalfHeight` — Half the height of the capsule cylinder portion in unscaled local units (centimeters). Total height is 2 * (CapsuleHalfHeight + CapsuleRadius).
- `CapsuleRadius` — Radius of the capsule hemispheres and cylinder in unscaled local units (centimeters).
- `bDynamicObstacle` — If true, registers this capsule as an avoidance obstacle in the RVO avoidance system so AI agents path around it.

## Key Functions

- `SetCapsuleSize(float NewRadius, float NewHalfHeight, bool bUpdateOverlaps)` — Sets both radius and half-height at once; preferred over setting them individually to trigger a single geometry rebuild.
- `SetCapsuleHalfHeight(float HalfHeight, bool bUpdateOverlaps)` — Sets only the half-height, leaving the radius unchanged.
- `SetCapsuleRadius(float Radius, bool bUpdateOverlaps)` — Sets only the radius, leaving the half-height unchanged.
- `GetScaledCapsuleHalfHeight()` — Returns the half-height multiplied by the world-space Z scale of this component.
- `GetScaledCapsuleRadius()` — Returns the radius multiplied by the world-space XY scale of this component.
- `GetUnscaledCapsuleHalfHeight()` — Returns the raw half-height before component scale is applied.
- `GetUnscaledCapsuleRadius()` — Returns the raw radius before component scale is applied.
- `InitCapsuleSize(float InRadius, float InHalfHeight)` — Initializes capsule dimensions without triggering overlap updates; use in constructors before BeginPlay.
