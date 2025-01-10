import util.{quoted}

pub type Class {
  // Gems
  SkillGems
  SupportGems
  SpiritGems
  // One Handed Weapons
  Claws
  Daggers
  Wands
  OneHandSwords
  OneHandAxes
  OneHandMaces
  Sceptres
  Spears
  Flails
  // Two Handed Weapons
  Bows
  Staves
  TwoHandSwords
  TwoHandAxes
  TwoHandMaces
  Quarterstaves
  Crossbows
  Traps
  FishingRods
  // Off-hand
  Quivers
  Shields
  Foci
  // Armour
  Gloves
  Boots
  BodyArmours
  Helmets
  // Jewellery
  Amulets
  Rings
  Belts
  // Flasks
  Flasks
  LifeFlasks
  ManaFlasks
  Charms
  // Currency
  Currency
  StackableCurrency
  DistilledEmotions
  Essence
  Splinter
  Catalysts
  // Waystones
  Waystones
  MapFragments
  MiscMapItems
  // Jewels
  Jewels
}

pub fn to_string(class: Class) -> String {
  case class {
    SkillGems -> "Skill Gems"
    SupportGems -> "Support Gems"
    SpiritGems -> "Spirit Gems"

    Claws -> "Claws"
    Daggers -> "Daggers"
    Wands -> "Wands"
    OneHandSwords -> "One Hand Swords"
    OneHandAxes -> "One Hand Axes"
    OneHandMaces -> "One Hand Maces"
    Sceptres -> "Sceptres"
    Spears -> "Spears"
    Flails -> "Flails"

    Bows -> "Bows"
    Staves -> "Staves"
    TwoHandSwords -> "Two Hand Swords"
    TwoHandAxes -> "Two Hand Axes"
    TwoHandMaces -> "Two Hand Maces"
    Quarterstaves -> "Quarterstaves"
    Crossbows -> "Crossbows"
    Traps -> "Traps"
    FishingRods -> "Fishing Rods"

    Quivers -> "Quivers"
    Shields -> "Shields"
    Foci -> "Foci"

    Gloves -> "Gloves"
    Boots -> "Boots"
    BodyArmours -> "Body Armours"
    Helmets -> "Helmets"

    Amulets -> "Amulets"
    Rings -> "Rings"
    Belts -> "Belts"

    Flasks -> "Flasks"
    LifeFlasks -> "Life Flasks"
    ManaFlasks -> "Mana Flasks"
    Charms -> "Charms"

    Currency -> "Currency"
    StackableCurrency -> "Stackable Currency"
    DistilledEmotions -> "Distilled Emotions"
    Essence -> "Essence"
    Splinter -> "Splinter"
    Catalysts -> "Catalysts"

    Waystones -> "Waystones"
    MapFragments -> "Map Fragments"
    MiscMapItems -> "Misc Map Items"

    Jewels -> "Jewels"
  }
  |> quoted
}
