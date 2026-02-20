# UAttributeSet

**Parent**: UObject
**Module**: GameplayAbilities

Container for gameplay attributes used by the Gameplay Ability System (GAS). Subclass this to define the numeric stats for your characters (health, mana, stamina, damage, etc.). Each attribute is a `FGameplayAttributeData` UPROPERTY that GAS tracks, replicates, and modifies through `UGameplayEffect` modifiers. One `UAttributeSet` subclass is registered per `UAbilitySystemComponent`.

## Key Properties

- `Health` — Example `FGameplayAttributeData` attribute representing current health; define your own via `UPROPERTY(ReplicatedUsing=OnRep_Health)`
- `MaxHealth` — Example attribute for the health ceiling; used with clamping in `PreAttributeChange`
- `Mana` — Example attribute for current mana resource
- `Damage` — Transient meta-attribute often used as an intermediary in execution calculations before applying to health
- `AttackPower` — Example attribute for outgoing damage scaling; referenced in gameplay effect modifier magnitudes

## Key Functions

- `PreAttributeChange(const FGameplayAttribute& Attribute, float& NewValue)` — Called before an attribute's current value changes; use to clamp `NewValue` (e.g. `NewValue = FMath::Clamp(NewValue, 0.f, GetMaxHealth())`)
- `PostGameplayEffectExecute(const FGameplayEffectModCallbackData& Data)` — Called after a `UGameplayEffect` modifies a base attribute; use to apply damage, trigger death, or clamp final values
- `PreAttributeBaseChange(const FGameplayAttribute& Attribute, float& NewValue)` — Called before the base value of an attribute changes; mirrors `PreAttributeChange` but for base (un-buffed) values
- `GetLifetimeReplicatedProps(TArray<FLifetimeProperty>& OutLifetimeProps)` — Override to declare which `FGameplayAttributeData` properties replicate via `DOREPLIFETIME_CONDITION_NOTIFY`
- `InitFromMetaDataTable(const UDataTable* DataTable)` — Bulk-initialises attribute base values from a `UDataTable` with rows matching attribute names
- `GetGameplayAttributeValueChangeDelegate(FGameplayAttribute Attribute)` — Returns a multicast delegate that fires whenever the attribute's value changes; use to bind UI updates
- `OnRep_Health()` — Example `RepNotify` function called on clients when `Health` replicates; must call `GAMEPLAYATTRIBUTE_REPNOTIFY(UMyAttributeSet, Health)`
- `ClampAttributeOnChange(const FGameplayAttribute& Attribute, float& NewValue)` — Helper pattern (not a built-in UE function) commonly implemented to centralise clamping logic across attributes
- `GetOwningActor()` — Returns the `AActor*` that owns the `UAbilitySystemComponent` this attribute set is registered with
- `GetOwningAbilitySystemComponent()` — Returns the `UAbilitySystemComponent*` this attribute set belongs to
- `PrintDebug()` — Logs all attribute current and base values to the output log for debugging
