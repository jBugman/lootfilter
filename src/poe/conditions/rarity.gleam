pub type Rarity {
  Normal
  Magic
  Rare
  Unique
}

pub fn to_string(tier: Rarity) -> String {
  case tier {
    Normal -> "Normal"
    Magic -> "Magic"
    Rare -> "Rare"
    Unique -> "Unique"
  }
}
