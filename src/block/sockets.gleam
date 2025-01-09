import gleam/int
import gleam/string

import block/op.{type Op}

pub type Sockets {
  Sockets(op: Op, value: Int)
}

pub fn to_string(sockets: Sockets) -> String {
  let Sockets(op, value) = sockets

  ["Sockets", op.to_string(op), int.to_string(value)]
  |> string.join(" ")
}
