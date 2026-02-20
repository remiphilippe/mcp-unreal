# AGameModeBase

**Parent**: AInfo
**Module**: Engine

AGameModeBase defines the rules and flow of the game. It only exists on the server (or standalone). It controls which pawn class to spawn for players, how players enter the game, and game-wide state. For games with match states (waiting, playing, game over), use AGameMode instead.

## Key Properties

- `DefaultPawnClass` — The APawn subclass spawned for new players (default: ADefaultPawn)
- `PlayerControllerClass` — The APlayerController subclass created for players
- `GameStateClass` — The AGameStateBase subclass for shared game state
- `HUDClass` — The AHUD subclass for the player's HUD
- `PlayerStateClass` — The APlayerState subclass for per-player state
- `SpectatorClass` — The ASpectatorPawn class for spectating players

## Key Functions

- `InitGame(const FString& MapName, const FString& Options, FString& ErrorMessage)` — Called before actors initialize. Override for custom game setup.
- `PreLogin(const FString& Options, const FString& Address, const FUniqueNetIdRepl& UniqueId, FString& ErrorMessage)` — Validates a player before login. Return non-empty ErrorMessage to reject.
- `PostLogin(APlayerController* NewPlayer)` — Called after a player successfully joins. Good place for initial spawn.
- `HandleStartingNewPlayer(APlayerController* NewPlayer)` — Spawns the default pawn for a new player.
- `SpawnDefaultPawnFor(AController* NewPlayer, AActor* StartSpot)` — Creates the pawn instance. Override to customize spawn.
- `FindPlayerStart(AController* Player)` — Finds a APlayerStart actor for spawning.
- `RestartPlayer(AController* NewPlayer)` — Respawns a player after death.
- `GetDefaultPawnClassForController(AController* InController)` — Returns the pawn class to spawn. Override for per-player pawn types.
- `ShouldSpawnAtStartSpot(AController* Player)` — Whether to use saved start position.
