import gleam/string

import block/op.{type Op}
import block/rarity_tier.{type RarityTier}

pub type Rarity {
  Rarity(op: Op, tier: RarityTier)
}

pub fn to_string(rarity: Rarity) -> String {
  let Rarity(op, tier) = rarity

  ["Rarity", op.to_string(op), rarity_tier.to_string(tier)]
  |> string.join(" ")
}
