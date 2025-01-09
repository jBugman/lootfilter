import conditions.{type Conditions}
import decorations.{type Decorations}
import gleam/list
import gleam/string
import gleam/string_tree
import styles.{type Styles}

pub type Block {
  Show(conditions: Conditions, decorations: Decorations)
  Hide(conditions: Conditions)
}

pub fn render(block: Block, global_style: Styles) -> String {
  case block {
    Show(cond, decor) -> render_("Show", cond, decor, global_style)
    Hide(cond) -> render_("Hide", cond, decorations.hidden, global_style)
  }
}

fn render_(
  prefix: String,
  cond: Conditions,
  decor: Decorations,
  global_style: Styles,
) -> String {
  let props =
    conditions.render(cond)
    |> list.append(decorations.render(decor, global_style))
    |> list.map(string.append("\t", _))
    |> string.join("\n")

  string_tree.from_strings([prefix, "\n", props, "\n\n"])
  |> string_tree.to_string
}
