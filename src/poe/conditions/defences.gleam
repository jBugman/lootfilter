import gleam/int
import gleam/string

import poe/op.{type Op}

pub type Defence {
  Armour(op: Op, value: Int)
  Evasion(op: Op, value: Int)
  EnergyShield(op: Op, value: Int)
}

pub fn to_string(defence: Defence) -> String {
  case defence {
    Armour(cmp, value) ->
      ["BaseArmour", op.to_string(cmp), int.to_string(value)]
      |> string.join(" ")

    Evasion(cmp, value) ->
      ["BaseEvasion", op.to_string(cmp), int.to_string(value)]
      |> string.join(" ")

    EnergyShield(cmp, value) ->
      ["BaseEnergyShield", op.to_string(cmp), int.to_string(value)]
      |> string.join(" ")
  }
}
