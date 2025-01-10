import gleam/int
import gleam/list
import gleam/string

import poe/conditions/base_type
import poe/conditions/class
import poe/conditions/rarity
import poe/op.{type Op}

pub type Condition {
  BaseType(List(base_type.BaseType))
  Class(List(class.Class))

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
      let xs = xs |> list.map(base_type.to_string)
      ["BaseType", ..xs] |> string.join(" ")
    }

    Class(xs) -> {
      let xs = xs |> list.map(class.to_string)
      ["Class", ..xs] |> string.join(" ")
    }

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
