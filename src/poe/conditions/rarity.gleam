import gleam/string

import poe/conditions/rarity_tier.{type RarityTier}
import poe/op.{type Op}

pub type Rarity {
  Rarity(op: Op, tier: RarityTier)
}

pub fn to_string(rarity: Rarity) -> String {
  let Rarity(op, tier) = rarity

  ["Rarity", op.to_string(op), rarity_tier.to_string(tier)]
  |> string.join(" ")
}
