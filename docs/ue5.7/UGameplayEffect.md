# UGameplayEffect

**Parent**: UObject
**Module**: GameplayAbilities

Defines a gameplay effect — a data-driven package of attribute modifiers, tags, cues, and execution logic applied through the Gameplay Ability System. Effects can be instant (one-shot attribute changes), duration-based (temporary buffs/debuffs), or infinite (persistent until explicitly removed).

## Key Properties

- `DurationPolicy` — Controls effect lifetime: `Instant` (applies and ends immediately), `HasDuration` (lasts a set time), `Infinite` (persists until explicitly removed)
- `DurationMagnitude` — Defines how long a `HasDuration` effect lasts; supports scalable floats and curve table lookups
- `Period` — Interval in seconds at which a periodic effect re-applies its modifiers (0 = not periodic)
- `Modifiers` — Array of `FGameplayModifierInfo` entries, each specifying an `Attribute` to modify, a `ModifierOp` (`Add`, `Multiply`, `Override`), and a `ModifierMagnitude`
- `Executions` — Array of `FGameplayEffectExecutionDefinition` structs for custom `UGameplayEffectExecutionCalculation` classes that compute complex modifier values
- `ConditionalGameplayEffects` — Array of additional effects to apply conditionally when this effect executes
- `StackingType` — Controls how multiple applications stack: `None` (each application is independent), `AggregateBySource` (stacks per source), `AggregateByTarget` (stacks per target)
- `StackLimitCount` — Maximum number of stacks allowed when `StackingType` is not `None`
- `StackDurationRefreshPolicy` — Whether applying a new stack refreshes (`RefreshOnSuccessfulApplication`) or preserves (`NeverRefresh`) the existing stack duration
- `StackPeriodResetPolicy` — Whether a new stack application resets (`ResetOnSuccessfulApplication`) or preserves (`NeverReset`) the periodic timer
- `GameplayCues` — Array of `FGameplayEffectCue` entries that trigger cosmetic `UGameplayCueNotify` events on application and removal
- `InheritableOwnedTagsContainer` — Tags granted to the owning actor for the effect's duration
- `ApplicationTagRequirements` — Tags the target must (and must not) have for the effect to apply
- `RemovalTagRequirements` — Tags whose presence or absence on the target triggers early removal of the effect

## Key Functions

- `GetDurationPolicy()` — Returns the `EGameplayEffectDurationType` enum value for this effect
- `GetPeriod()` — Returns the period interval as a `float`
- `GetStackLimitCount()` — Returns the maximum stack count as an `int32`
- `GetStackingType()` — Returns the `EGameplayEffectStackingType` enum value
- `FindComponent(TSubclassOf<UGameplayEffectComponent> ComponentClass)` — Finds an attached `UGameplayEffectComponent` by class
- `GetGrantedAbilities()` — Returns the array of `FGameplayAbilitySpecDef` entries granted by this effect
- `GetBlockedAbilityTags()` — Returns the `FInheritedTagContainer` of ability tags blocked while this effect is active
- `GetDynamicAssetTags()` — Returns the `FGameplayTagContainer` of runtime-assigned asset tags for this effect
