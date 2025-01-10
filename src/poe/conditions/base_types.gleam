import gleam/list
import gleam/string

import util.{quoted}

pub type BaseType {
  BaseType(base_type: String)
  BaseTypes(base_types: List(String))
}

pub fn to_string(base_type: BaseType) -> String {
  case base_type {
    BaseType(base_type) -> ["BaseType", quoted(base_type)] |> string.join(" ")

    BaseTypes(base_types) -> {
      let qs = base_types |> list.map(quoted)
      ["BaseType", ..qs] |> string.join(" ")
    }
  }
}
