import gleam/int
import gleam/string

import poe/conditions/base_types
import poe/conditions/class
import poe/conditions/quality
import poe/conditions/rarity
import poe/op.{type Op}

pub type Condition {
  BaseType(base_type: base_types.BaseType)
  Class(class: class.Class)

  Rarity(op: Op, tier: rarity.Rarity)

  Armour(op: Op, value: Int)
  Evasion(op: Op, value: Int)
  EnergyShield(op: Op, value: Int)

  Quality(quality: quality.Quality)
  Sockets(op: Op, amount: Int)
}

pub fn to_string(cond: Condition) -> String {
  case cond {
    BaseType(x) -> base_types.to_string(x)
    Class(x) -> class.to_string(x)

    Rarity(op, tier) ->
      ["Rarity", op.to_string(op), rarity.to_string(tier)]
      |> string.join(" ")

    Armour(op, value) -> comparable("BaseArmour", op, value)
    Evasion(op, value) -> comparable("BaseEvasion", op, value)
    EnergyShield(op, value) -> comparable("BaseEnergyShield", op, value)

    Quality(x) -> quality.to_string(x)
    Sockets(op, amount) -> comparable("Sockets", op, amount)
  }
}

fn comparable(stat: String, op: Op, value: Int) -> String {
  [stat, op.to_string(op), int.to_string(value)]
  |> string.join(" ")
}
