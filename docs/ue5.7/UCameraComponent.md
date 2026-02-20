# UCameraComponent

**Parent**: USceneComponent
**Module**: Engine

UCameraComponent defines a viewpoint that can be activated by a player camera manager. Attach it to a pawn or actor and call `SetViewTarget` on the player controller to switch to it. It supports perspective and orthographic projection, post-process effects blending, and aspect ratio constraints. Use with USpringArmComponent for third-person camera rigs.

## Key Properties

- `FieldOfView` — Horizontal field of view angle in degrees for perspective projection (default 90).
- `AspectRatio` — Width-to-height ratio used when bConstrainAspectRatio is true (e.g., 1.777 for 16:9).
- `OrthoWidth` — World-space width of the orthographic view volume when ProjectionMode is Orthographic.
- `bUsePawnControlRotation` — If true, this component's world rotation is driven by the owning pawn's controller rotation each tick.
- `bConstrainAspectRatio` — If true, black bars are added to enforce the AspectRatio; prevents stretching on non-matching displays.
- `PostProcessSettings` — FPostProcessSettings struct controlling bloom, exposure, color grading, depth of field, and other post-process parameters.
- `PostProcessBlendWeight` — 0–1 weight applied to PostProcessSettings when blending with other post-process volumes (1 = full override).
- `ProjectionMode` — ECameraProjectionMode::Perspective or ECameraProjectionMode::Orthographic.

## Key Functions

- `GetCameraView(float DeltaTime, FMinimalViewInfo& DesiredView)` — Fills DesiredView with the current camera parameters; called by the camera manager each frame.
- `SetFieldOfView(float InFieldOfView)` — Updates the horizontal FOV in degrees.
- `SetAspectRatio(float InAspectRatio)` — Updates the constrained aspect ratio value.
- `SetProjectionMode(ECameraProjectionMode::Type InProjectionMode)` — Switches between perspective and orthographic projection.
- `SetPostProcessBlendWeight(float Weight)` — Sets how strongly this camera's PostProcessSettings override scene volumes.
- `SetConstraintAspectRatio(bool bInConstrainAspectRatio)` — Enables or disables letterboxing for aspect ratio enforcement.
- `GetForwardVector()` — Returns the forward direction vector of this camera in world space (inherited from USceneComponent).
- `AddOrUpdateBlendable(TScriptInterface<IBlendableInterface> InBlendableObject, float InWeight)` — Adds or updates a blendable post-process material or settings object on this camera.
