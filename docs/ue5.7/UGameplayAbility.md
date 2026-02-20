# UGameplayAbility

**Parent**: UObject
**Module**: GameplayAbilities

Individual gameplay ability with lifecycle, costs, cooldowns, and tag requirements. Subclass `UGameplayAbility` to implement discrete actions (attacks, spells, interactions) that integrate with the Gameplay Ability System. The ability defines what tags must be present to activate, what tags it applies while active, what it costs, and how long it must cool down before re-use.

## Key Properties

- `AbilityTags` — Gameplay tags identifying this ability; used by other abilities and effects to reference or cancel it by category
- `CancelAbilitiesWithTag` — Tag container; any currently active ability whose `AbilityTags` match these tags is cancelled when this ability activates
- `BlockAbilitiesWithTag` — Tag container; while this ability is active, other abilities whose `AbilityTags` match these tags cannot be activated
- `ActivationOwnedTags` — Tags granted to the ability system component for the duration of this ability's activation; automatically removed on `EndAbility`
- `ActivationRequiredTags` — The ability can only activate if the owner's ability system component currently has all of these tags
- `ActivationBlockedTags` — The ability cannot activate if the owner's ability system component has any of these tags
- `CostGameplayEffectClass` — Optional `UGameplayEffect` subclass applied on commit to deduct resources (mana, stamina, ammo); if the cost cannot be paid the commit fails
- `CooldownGameplayEffectClass` — Optional `UGameplayEffect` subclass with a duration applied on commit to block re-activation until the duration expires
- `bReplicateInputDirectly` — When true input press/release events are replicated to the server immediately; use for abilities requiring server-side input timing
- `NetExecutionPolicy` — Controls where the ability runs: `LocalPredicted` (client predicts, server confirms), `LocalOnly` (client only, cosmetic), `ServerInitiated` (server decides, client reacts), `ServerOnly` (server runs, never on client)
- `InstancingPolicy` — Memory strategy: `NonInstanced` (no instance created, all state must be on ASC), `InstancedPerActor` (one instance shared across all activations), `InstancedPerExecution` (new instance per activation, safest for async abilities)

## Key Functions

- `CanActivateAbility(const FGameplayAbilitySpecHandle Handle, const FGameplayAbilityActorInfo* ActorInfo, const FGameplayTagContainer* SourceTags, const FGameplayTagContainer* TargetTags, FGameplayTagContainer* OptionalRelevantTags)` — Full pre-activation check including tag requirements, cooldown, cost, and any custom conditions; called by the ASC before invoking `ActivateAbility`
- `ActivateAbility(FGameplayAbilitySpecHandle Handle, const FGameplayAbilityActorInfo* ActorInfo, FGameplayAbilityActivationInfo ActivationInfo, const FGameplayEventData* TriggerEventData)` — Main entry point; override to implement ability logic; must call `CommitAbility` or `EndAbility` before returning for instanced abilities
- `EndAbility(FGameplayAbilitySpecHandle Handle, const FGameplayAbilityActorInfo* ActorInfo, FGameplayAbilityActivationInfo ActivationInfo, bool bReplicateEndAbility, bool bWasCancelled)` — Cleans up the ability, removes `ActivationOwnedTags`, and notifies the ASC; always call this to terminate an ability
- `CommitAbility(FGameplayAbilitySpecHandle Handle, const FGameplayAbilityActorInfo* ActorInfo, FGameplayAbilityActivationInfo ActivationInfo)` — Applies cost and cooldown effects; returns false if either cannot be applied; call early in `ActivateAbility` to gate execution on resource availability
- `CommitAbilityCost(FGameplayAbilitySpecHandle Handle, const FGameplayAbilityActorInfo* ActorInfo, FGameplayAbilityActivationInfo ActivationInfo)` — Applies only the cost effect without triggering the cooldown; used when cost and cooldown should be committed separately
- `CommitAbilityCooldown(FGameplayAbilitySpecHandle Handle, const FGameplayAbilityActorInfo* ActorInfo, FGameplayAbilityActivationInfo ActivationInfo, const bool ForceCooldown)` — Applies only the cooldown effect; `ForceCooldown` bypasses the cooldown check and applies regardless
- `ApplyGameplayEffectToOwner(FGameplayAbilitySpecHandle Handle, const FGameplayAbilityActorInfo* ActorInfo, FGameplayAbilityActivationInfo ActivationInfo, const UGameplayEffect* GameplayEffect, float GameplayEffectLevel, int32 Stacks)` — Applies a gameplay effect to the ability's owner actor
- `ApplyGameplayEffectToTarget(FGameplayAbilitySpecHandle Handle, const FGameplayAbilityActorInfo* ActorInfo, FGameplayAbilityActivationInfo ActivationInfo, FGameplayAbilityTargetDataHandle Target, TSubclassOf<UGameplayEffect> GameplayEffectClass, float GameplayEffectLevel, int32 Stacks)` — Applies a gameplay effect to one or more target actors described by the target data handle
- `SendGameplayEvent(FGameplayTag EventTag, FGameplayEventData Payload)` — Fires a gameplay event on the owner's ASC that can trigger other abilities listening for this tag via the `GameplayEvent` trigger type
- `GetAbilitySystemComponentFromActorInfo()` — Returns the `UAbilitySystemComponent` of the owning actor; safe to call during activation
- `GetAvatarActorFromActorInfo()` — Returns the avatar `AActor` (the physical pawn) associated with this ability's actor info; may differ from the owner when abilities are granted to a player state
- `GetOwningActorFromActorInfo()` — Returns the owner `AActor` (the actor holding the ASC, often the player state or pawn)
- `K2_EndAbility()` — Blueprint-callable wrapper for `EndAbility` with the replicate and cancelled flags preset; use this from Blueprint implementations instead of calling `EndAbility` directly
- `MakeOutgoingGameplayEffectSpec(TSubclassOf<UGameplayEffect> GameplayEffectClass, float Level, int32 Stacks)` — Creates an `FGameplayEffectSpecHandle` pre-configured with this ability's context, source tags, and level; pass the result to apply functions
