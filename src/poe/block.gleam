import gleam/list
import gleam/string_tree.{type StringTree}

import poe/actions.{type Action}
import poe/conditions.{type Condition}

pub type Block {
  Show(conditions: List(Condition), actions: List(Action))
  // Hide(conditions: Conditions)
}

const tab = "  "

pub fn to_string(block: Block) -> String {
  case block {
    Show(conditions, actions) -> {
      let conditions = conditions_to_string_tree(conditions)
      let actions = actions_to_string_tree(actions)

      [string_tree.from_string("Show"), conditions, actions]
      |> string_tree.join("\n")
      |> string_tree.to_string
    }
    // Hide(cond) -> render_("Hide", cond, decorations.hidden, global_style)
  }
}

fn conditions_to_string_tree(conditions: List(Condition)) -> StringTree {
  conditions
  |> list.map(fn(c) {
    c
    |> conditions.to_string_tree
    |> string_tree.prepend(tab)
  })
  |> string_tree.concat
}

fn actions_to_string_tree(actions: List(Action)) -> StringTree {
  actions
  |> list.map(fn(a) {
    a
    |> actions.to_string_tree
    |> string_tree.prepend(tab)
  })
  |> string_tree.concat
}
