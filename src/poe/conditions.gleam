import gleam/string_tree.{type StringTree}

import poe/conditions/base_types
import poe/conditions/class
import poe/conditions/defences
import poe/conditions/quality
import poe/conditions/rarity
import poe/conditions/sockets

pub type Condition {
  BaseType(base_type: base_types.BaseType)
  Class(class: class.Class)
  Defence(defence: defences.Defence)
  Quality(quality: quality.Quality)
  Sockets(sockets: sockets.Sockets)
  Rarity(rarity: rarity.Rarity)
}

pub fn to_string(cond: Condition) -> String {
  case cond {
    BaseType(x) -> base_types.to_string(x)
    Class(x) -> class.to_string(x)
    Defence(x) -> defences.to_string(x)
    Quality(x) -> quality.to_string(x)
    Sockets(x) -> sockets.to_string(x)
    Rarity(x) -> rarity.to_string(x)
  }
}

pub fn to_string_tree(cond: Condition) -> StringTree {
  cond |> to_string |> string_tree.from_string
}
