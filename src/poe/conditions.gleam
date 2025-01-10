import gleam/int
import gleam/list
import gleam/string

import poe/conditions/base_types
import poe/conditions/class
import poe/conditions/rarity
import poe/op.{type Op}

pub type Condition {
  BaseType(List(base_types.BaseType))
  Class(class: class.Class)

  Rarity(op: Op, tier: rarity.Rarity)

  Armour(op: Op, value: Int)
  Evasion(op: Op, value: Int)
  EnergyShield(op: Op, value: Int)

  Quality(op: Op, value: Int)
  Sockets(op: Op, amount: Int)
}

pub fn to_string(cond: Condition) -> String {
  case cond {
    BaseType(xs) -> {
      let bt = xs |> list.map(base_types.to_string)
      ["BaseType", ..bt] |> string.join(" ")
    }
    Class(x) -> class.to_string(x)

    Rarity(op, tier) ->
      ["Rarity", op.to_string(op), rarity.to_string(tier)]
      |> string.join(" ")

    Armour(op, value) -> comparable("BaseArmour", op, value)
    Evasion(op, value) -> comparable("BaseEvasion", op, value)
    EnergyShield(op, value) -> comparable("BaseEnergyShield", op, value)

    Quality(op, value) -> comparable("Quality", op, value)
    Sockets(op, amount) -> comparable("Sockets", op, amount)
  }
}

fn comparable(stat: String, op: Op, value: Int) -> String {
  [stat, op.to_string(op), int.to_string(value)]
  |> string.join(" ")
}
