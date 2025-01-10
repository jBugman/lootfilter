import gleeunit/should

import poe/conditions
import poe/conditions/base_types
import poe/conditions/class
import poe/conditions/rarity
import poe/conditions/rarity_tier
import poe/op

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

pub fn sockets_gt_test() {
  conditions.Sockets(op.GT, 0)
  |> conditions.to_string
  |> should.equal("Sockets > 0")
}
