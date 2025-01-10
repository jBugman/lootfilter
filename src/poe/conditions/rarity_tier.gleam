pub type RarityTier {
  Normal
  Magic
  Rare
  Unique
}

pub fn to_string(tier: RarityTier) -> String {
  case tier {
    Normal -> "Normal"
    Magic -> "Magic"
    Rare -> "Rare"
    Unique -> "Unique"
  }
}
