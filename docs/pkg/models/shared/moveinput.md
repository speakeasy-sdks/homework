# MoveInput


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Accuracy`                                                                             | **int*                                                                                 | :heavy_minus_sign:                                                                     | The percent value of how likely this move is to be successful                          |
| `ContestCombos`                                                                        | [*shared.ContestComboSets](../../../pkg/models/shared/contestcombosets.md)             | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `ContestEffect`                                                                        | [*shared.MoveContestEffect](../../../pkg/models/shared/movecontesteffect.md)           | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `ContestType`                                                                          | [*shared.MoveContestType](../../../pkg/models/shared/movecontesttype.md)               | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `DamageClass`                                                                          | [*shared.DamageClass](../../../pkg/models/shared/damageclass.md)                       | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `EffectChance`                                                                         | **int*                                                                                 | :heavy_minus_sign:                                                                     | The percent value of the additional effects this move has occuring                     |
| `EffectChanges`                                                                        | [][shared.AbilityEffectChange](../../../pkg/models/shared/abilityeffectchange.md)      | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `EffectEntries`                                                                        | [][shared.VerboseEffect](../../../pkg/models/shared/verboseeffect.md)                  | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Generation`                                                                           | [shared.MoveGeneration](../../../pkg/models/shared/movegeneration.md)                  | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `ID`                                                                                   | *int*                                                                                  | :heavy_check_mark:                                                                     | The identifier for this move resource                                                  |
| `Meta`                                                                                 | [*shared.MoveMetaData](../../../pkg/models/shared/movemetadata.md)                     | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Name`                                                                                 | *string*                                                                               | :heavy_check_mark:                                                                     | The name for this move resource                                                        |
| `Names`                                                                                | [][shared.Name](../../../pkg/models/shared/name.md)                                    | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `PastValues`                                                                           | [][shared.PastMoveStatValues](../../../pkg/models/shared/pastmovestatvalues.md)        | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Power`                                                                                | **int*                                                                                 | :heavy_minus_sign:                                                                     | The base power of this move with a value of 0 if it does not have a base power         |
| `Pp`                                                                                   | *int*                                                                                  | :heavy_check_mark:                                                                     | Power points. The number of times this move can be used                                |
| `Priority`                                                                             | *int*                                                                                  | :heavy_check_mark:                                                                     | A value of 0 means this move goes last in the turn, and 1 means it goes first          |
| `StatChanges`                                                                          | [][shared.MoveStatChange](../../../pkg/models/shared/movestatchange.md)                | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `SuperContestEffect`                                                                   | [*shared.MoveSuperContestEffect](../../../pkg/models/shared/movesupercontesteffect.md) | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Target`                                                                               | [shared.Target](../../../pkg/models/shared/target.md)                                  | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `Type`                                                                                 | [shared.MoveType](../../../pkg/models/shared/movetype.md)                              | :heavy_check_mark:                                                                     | N/A                                                                                    |