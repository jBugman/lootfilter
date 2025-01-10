import gleam/list
import gleam/string

import poe/actions.{type Action}
import poe/conditions.{type Condition}

pub type Block {
  Show(conditions: List(Condition), actions: List(Action))
  // Hide(conditions: Conditions)
}

const tab = "  "

pub fn to_string(block: Block) -> String {
  case block {
    Show(cond, act) ->
      [
        ["Show"],
        cond |> indent_with(conditions.to_string),
        act |> indent_with(actions.to_string),
      ]
      |> list.flatten
      |> string.join("\n")
    // Hide(cond) -> render_("Hide", cond, decorations.hidden, global_style)
  }
}

fn indent_with(lines: List(s), f: fn(s) -> String) -> List(String) {
  list.map(lines, fn(line) { tab <> f(line) })
}
