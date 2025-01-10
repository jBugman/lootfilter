import gleeunit/should

import poe/conditions as c
import poe/conditions/class
import poe/conditions/rarity
import poe/op

pub fn base_type_1_test() {
  c.BaseType(["Short Sword", "Short Bow"])
  |> c.to_string
  |> should.equal("BaseType \"Short Sword\" \"Short Bow\"")
}

pub fn base_type_2_test() {
  c.BaseType(["Short Sword"])
  |> c.to_string
  |> should.equal("BaseType \"Short Sword\"")
}

pub fn class_test() {
  c.Class([class.Currency])
  |> c.to_string
  |> should.equal("Class \"Currency\"")
}

pub fn classes_test() {
  c.Class([class.Bows, class.Wands])
  |> c.to_string
  |> should.equal("Class \"Bows\" \"Wands\"")
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
