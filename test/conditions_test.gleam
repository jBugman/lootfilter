import gleeunit/should

import poe/conditions as c
import poe/conditions/base_types
import poe/conditions/class
import poe/conditions/rarity
import poe/op

pub fn conditions_base_type_test() {
  base_types.BaseTypes(["Short Sword", "Short Bow"])
  |> c.BaseType
  |> c.to_string
  |> should.equal("BaseType \"Short Sword\" \"Short Bow\"")
}

pub fn conditions_class_test() {
  class.Class("Bows")
  |> c.Class
  |> c.to_string
  |> should.equal("Class \"Bows\"")
}

pub fn conditions_rarity_test() {
  c.Rarity(op.LT, rarity.Rare)
  |> c.to_string
  |> should.equal("Rarity < Rare")
}

pub fn sockets_gt_test() {
  c.Sockets(op.GT, 0)
  |> c.to_string
  |> should.equal("Sockets > 0")
}

pub fn armour_eq_test() {
  c.Armour(op.Eq, 100)
  |> c.to_string
  |> should.equal("BaseArmour = 100")
}

pub fn evasion_lt_test() {
  c.Evasion(op.LT, 82)
  |> c.to_string
  |> should.equal("BaseEvasion < 82")
}

pub fn es_gte_test() {
  c.EnergyShield(op.GTE, 120)
  |> c.to_string
  |> should.equal("BaseEnergyShield >= 120")
}

pub fn quality_not_eq_test() {
  c.Quality(op.NotEq, 8)
  |> c.to_string
  |> should.equal("Quality != 8")
}
