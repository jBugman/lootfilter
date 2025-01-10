import gleam/int
import gleam/string

import poe/conditions/base_types
import poe/conditions/class
import poe/conditions/defences
import poe/conditions/quality
import poe/conditions/rarity
import poe/op.{type Op}

pub type Condition {
  BaseType(base_type: base_types.BaseType)
  Class(class: class.Class)
  Defence(defence: defences.Defence)
  Quality(quality: quality.Quality)
  Sockets(op: Op, amount: Int)
  Rarity(rarity: rarity.Rarity)
}

pub fn to_string(cond: Condition) -> String {
  case cond {
    BaseType(x) -> base_types.to_string(x)

    Class(x) -> class.to_string(x)

    Defence(x) -> defences.to_string(x)

    Quality(x) -> quality.to_string(x)

    Sockets(op, amount) ->
      ["Sockets", op.to_string(op), int.to_string(amount)] |> string.join(" ")

    Rarity(x) -> rarity.to_string(x)
  }
}
