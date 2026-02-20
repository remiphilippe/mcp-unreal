# UAnimMontage

**Parent**: UAnimCompositeBase
**Module**: Engine

Animation montage asset compositing multiple animation sequences into a single playable unit with named sections, slot-based blending, and animation notifies. Montages are played through `UAnimInstance::Montage_Play` and are the primary mechanism for gameplay-triggered animations such as attacks, interactions, and reactions. They support branching between sections and blending in/out independently of the base pose.

## Key Properties

- `BlendIn` — `FAlphaBlend` defining the curve and duration for blending the montage in when playback begins
- `BlendOut` — `FAlphaBlend` defining the curve and duration for blending the montage out when it ends or is stopped
- `BlendOutTriggerTime` — `float` time in seconds before the montage end at which the blend-out begins; set to `-1` to use the blend-out duration as the trigger offset
- `bEnableAutoBlendOut` — `bool` controlling whether the montage automatically initiates blend-out when it reaches the end; disable for looping or manually-terminated montages
- `CompositeSections` — `TArray<FCompositeSection>` listing all named sections within the montage, each with a name, start time, and optional next-section link for branching
- `SlotAnimTracks` — `TArray<FSlotAnimationTrack>` mapping animation slot names to the sequences blended into those slots; a montage can animate multiple skeleton regions simultaneously using different slots
- `AnimNotifyTracks` — `TArray<FAnimNotifyTrack>` holding the notify and notify state events placed along the montage timeline
- `bLoop` — `bool` enabling continuous looping when the montage reaches its end
- `RateScale` — `float` global playback rate multiplier applied on top of the per-play-call rate; default `1.0`
- `TimeStretchCurveName` — `FName` of a curve asset used to non-uniformly stretch the montage timeline at runtime without altering the underlying sequences

## Key Functions

- `GetSectionIndex(FName SectionName)` — Returns `int32` index of the named section, or `INDEX_NONE` if not found
- `GetSectionName(int32 SectionIndex)` — Returns `FName` for the section at the given index
- `GetNumSections()` — Returns `int32` total number of sections defined in this montage
- `IsValidSectionName(FName SectionName)` — Returns `bool` indicating whether a section with the given name exists in this montage
- `GetDefaultBlendInTime()` — Returns `float` blend-in duration in seconds as configured on the asset
- `GetDefaultBlendOutTime()` — Returns `float` blend-out duration in seconds as configured on the asset
- `GetSectionLength(int32 SectionIndex)` — Returns `float` duration in seconds of the specified section
- `GetPlayLength()` — Returns `float` total playback duration of the full montage in seconds
- `GetAnimationData(FName SlotName)` — Returns `FAnimTrack*` for the animation track associated with the given slot name, or `nullptr` if the slot is not used by this montage
- `HasRootMotion()` — Returns `bool` indicating whether any sequence within this montage contains root motion data
