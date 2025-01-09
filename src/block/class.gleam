import gleam/list
import gleam/string

import util.{quoted}

pub type Class {
  Class(class: String)
  Classes(classes: List(String))
}

pub fn to_string(class: Class) -> String {
  case class {
    Class(class) -> ["Class", quoted(class)] |> string.join(" ")

    Classes(classes) -> {
      let qs = classes |> list.map(quoted)
      ["Class", ..qs] |> string.join(" ")
    }
  }
}
