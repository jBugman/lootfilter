import gleeunit/should

import block/base_types
import block/class
import block/conditions
import block/op
import block/rarity
import block/rarity_tier

pub fn conditions_base_type_test() {
  base_types.BaseTypes(["Short Sword", "Short Bow"])
  |> conditions.BaseType
  |> conditions.to_string
  |> should.equal("BaseType \"Short Sword\" \"Short Bow\"")
}

pub fn conditions_class_test() {
  class.Class("Bows")
  |> conditions.Class
  |> conditions.to_string
  |> should.equal("Class \"Bows\"")
}

pub fn conditions_rarity_test() {
  rarity.Rarity(op.LT, rarity_tier.Rare)
  |> conditions.Rarity
  |> conditions.to_string
  |> should.equal("Rarity < Rare")
}
