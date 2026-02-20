# UAbilitySystemComponent

**Parent**: UGameplayTasksComponent
**Module**: GameplayAbilities

Core GAS component managing abilities, effects, and attributes. The Gameplay Ability System (GAS) component is the central hub of the GAS framework. Attach it to any actor that needs abilities, gameplay effects, or attribute sets. It handles replication of ability state, effect application, tag management, and attribute modification.

## Key Properties

- `ActivatableAbilities` — Array of `FGameplayAbilitySpec` representing every ability that has been granted to this component; each spec tracks the ability class, level, source object, and activation state
- `ActiveGameplayEffects` — `FActiveGameplayEffectsContainer` managing all currently applied gameplay effects including their duration, period, and modifier stacks
- `SpawnedAttributes` — Array of `UAttributeSet*` instances owned by this component; attribute sets define the numerical stats (health, mana, damage) that effects modify
- `DefaultStartingData` — Array of `FAttributeDefaults` used to initialize attribute set values from a data table at component start; set these in the actor defaults
- `AffectedAnimInstanceTag` — Gameplay tag used to route GAS-driven animation notify events to a specific anim instance on the owning actor's mesh
- `bSuppressGameplayCues` — When true all gameplay cues (cosmetic effects, sounds, particles) triggered by this component are silently dropped; used for server-only actors
- `bSuppressGrantAbility` — When true calls to `GiveAbility` are ignored; used during initialization sequences where granting must be deferred
- `ReplicationMode` — Controls network replication fidelity: `Full` replicates everything (single-player/listen servers), `Mixed` replicates only to the owning client (recommended for players), `Minimal` replicates only tags and cues (recommended for AI)

## Key Functions

- `GiveAbility(FGameplayAbilitySpec Spec)` — Grants an ability to this component and returns an `FGameplayAbilitySpecHandle` for future activation or removal; the spec sets the ability class, level, and input binding
- `TryActivateAbility(FGameplayAbilitySpecHandle Handle, bool bAllowRemoteActivation)` — Attempts to activate a previously granted ability, running all tag requirement and cost checks; returns true if activation succeeded
- `CancelAbilityHandle(FGameplayAbilitySpecHandle Handle)` — Cancels a specific active ability instance by its handle, calling `EndAbility` with the cancelled flag set
- `ApplyGameplayEffectToSelf(const UGameplayEffect* GameplayEffect, float Level, FGameplayEffectContextHandle EffectContext)` — Applies a gameplay effect to this component as both source and target; returns an `FActiveGameplayEffectHandle` for tracking or early removal
- `RemoveActiveGameplayEffect(FActiveGameplayEffectHandle Handle, int32 StacksToRemove)` — Removes a specific active effect; pass `-1` for `StacksToRemove` to remove all stacks at once
- `GetAttributeValue(FGameplayAttribute Attribute)` — Returns the current final value of the specified attribute after all modifier aggregation
- `SetNumericAttributeBase(FGameplayAttribute Attribute, float NewBaseValue)` — Sets the base value of an attribute, bypassing effect modifiers; use for initialization only, not gameplay changes
- `HasMatchingGameplayTag(FGameplayTag TagToCheck)` — Returns true if this component currently has the exact specified tag (no parent matching); use `HasMatchingGameplayTagAny` for broader checks
- `AddLooseGameplayTag(FGameplayTag Tag, int32 Count)` — Adds a tag that is not tied to any ability or effect, bypassing the normal grant system; must be manually removed; not replicated by default
- `InitAbilityActorInfo(AActor* OwnerActor, AActor* AvatarActor)` — Sets the owner (the actor holding this component) and avatar (the physical actor in the world, often the pawn) and must be called before using any GAS functionality; call again on possession changes
