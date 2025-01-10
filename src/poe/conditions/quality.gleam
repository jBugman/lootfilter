import gleam/int
import gleam/string

import poe/op.{type Op}

pub type Quality {
  Quality(op: Op, value: Int)
}

pub fn to_string(quality: Quality) -> String {
  let Quality(op, value) = quality

  ["Quality", op.to_string(op), int.to_string(value)]
  |> string.join(" ")
}
