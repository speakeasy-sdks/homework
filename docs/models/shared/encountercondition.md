# EncounterCondition


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ID`                                                               | *int64*                                                            | :heavy_check_mark:                                                 | The identifier for this encounter condition resource               |
| `Name`                                                             | *string*                                                           | :heavy_check_mark:                                                 | The name for this encounter condition resource                     |
| `Names`                                                            | [][Name](../../models/shared/name.md)                              | :heavy_check_mark:                                                 | The name of this encounter condition listed in different languages |
| `Values`                                                           | [][NamedAPIResource](../../models/shared/namedapiresource.md)      | :heavy_check_mark:                                                 | A list of values that are used with this encounter condition       |